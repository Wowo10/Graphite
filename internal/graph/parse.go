package graph

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type Node struct {
	ID string `json:"id"`
}

type Edge struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

type Graph struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}

type goListPkg struct {
	ImportPath string   `json:"ImportPath"`
	Imports    []string `json:"Imports"`
}

type goListModule struct {
	Path string `json:"Path"`
}

func BuildGraph() (*Graph, error) {
	modulePath, err := findModulePath()
	if err != nil {
		return nil, err
	}

	log.Println("modPath:", modulePath)

	pkgs, err := listPackages(modulePath)
	if err != nil {
		return nil, err
	}

	graph := &Graph{
		Nodes: make([]Node, 0, len(pkgs)),
		Edges: []Edge{},
	}

	for _, p := range pkgs {
		graph.Nodes = append(graph.Nodes, Node{
			ID: strings.TrimPrefix(p.ImportPath, modulePath),
		})
		for _, imp := range p.Imports {
			if cutImp, ok := strings.CutPrefix(imp, modulePath); ok {
				graph.Edges = append(graph.Edges, Edge{
					Target: strings.TrimPrefix(p.ImportPath, modulePath),
					Source: cutImp})
			}
		}
	}

	return graph, nil
}

func listPackages(modulePath string) ([]goListPkg, error) {
	pkgsCmd := exec.Command("go", "list", "-json", "./...")
	pkgsOut, err := pkgsCmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to run 'go list -json ./...': %w", err)
	}

	dec := json.NewDecoder(bytes.NewReader(pkgsOut))

	pkgs := make([]goListPkg, 0)

	for {
		var p goListPkg
		if err := dec.Decode(&p); err != nil {
			if err.Error() == "EOF" {
				break
			}
			if strings.Contains(err.Error(), "EOF") {
				break
			}
			return nil, fmt.Errorf("error decoding go list json: %w", err)
		}

		if strings.HasPrefix(p.ImportPath, modulePath) {
			log.Println("append:", p.ImportPath)
			pkgs = append(pkgs, p)
		}
	}
	return pkgs, nil
}

func findModulePath() (string, error) {
	modCmd := exec.Command("go", "list", "-m", "-json")
	modOut, err := modCmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to run 'go list -m -json': %w", err)
	}
	var mod goListModule
	if err := json.Unmarshal(modOut, &mod); err != nil {
		return "", fmt.Errorf("failed to parse module info: %w", err)
	}
	modulePath := strings.TrimSpace(mod.Path)
	if modulePath == "" {
		return "", fmt.Errorf("module path empty (are you in a go module?)")
	}
	return modulePath, nil
}
