fetch("/graph.json")
  .then(r => r.json())
  .then(g => {
    const elements = []

    g.nodes.forEach(n => elements.push({ data: { id: n.id } }))
    g.edges.forEach(e => elements.push({ data: { source: e.source, target: e.target } }))

    console.log(g.layout)

    layout = { name: "cose" }

    if (g.layout === "dagre") {
      layout = { name: "dagre", rankDir: "LR" }
    } else if (g.layout === "klay") {
      layout = {
        name: "klay",
        klay: {
          direction: "RIGHT",
          edgeRouting: "ORTHOGONAL",
          spacing: 30
        }
      }
    } else if (g.layout === "cola") {
      layout = {
        name: 'cola',
        avoidOverlap: true,
        nodeSpacing: 30
      }
    }

    cytoscape({
      container: document.getElementById("cy"),
      elements: elements,
      layout: layout,
      style: [
        {
          selector: "node", style: {
            "background-color": "#007acc",
            "label": "data(id)"
          }
        },
        {
          selector: "edge", style: {
            "width": 2,
            "curve-style": "bezier",
            "target-arrow-shape": "triangle"
          }
        },
      ],
    })
  })
