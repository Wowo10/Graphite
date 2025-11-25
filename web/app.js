fetch("/graph.json")
  .then(r => r.json())
  .then(g => {
    const elements = [];

    g.nodes.forEach(n => elements.push({ data: { id: n.id } }));
    g.edges.forEach(e => elements.push({ data: { source: e.source, target: e.target } }));

    cytoscape({
      container: document.getElementById("cy"),
      elements: elements,
      layout: { name: "cose" },
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
    });
  });
