//go:build !test

package main

import (
	"fmt"
	"os"

	"github.com/vc-souza/gga/algo"
	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
	"github.com/vc-souza/gga/viz"
)

const (
	fileIn  = "BFS-before.dot"
	fileOut = "BFS-after.dot"
)

var soloVertex = "a#"

var (
	defaultGraphFmt = ds.FmtAttrs{
		"layout":  "dot",
		"nodesep": "0.8",
		"ranksep": "0.5",
		"pad":     "0.2",
	}

	defaultVertexFmt = ds.FmtAttrs{
		"shape":     "Mrecord",
		"style":     "filled",
		"fillcolor": "#7289da",
		"fontcolor": "#ffffff",
		"color":     "#ffffff",
		"penwidth":  "1.1",
	}

	defaultEdgeFmt = ds.FmtAttrs{
		"penwidth": "1.2",
	}
)

func buildBFSInput() (*ds.Graph[ds.Text], *ds.Text) {
	g, vars, err := ds.NewTextParser().Parse(ut.BasicUUG + soloVertex)

	if err != nil {
		panic(err)
	}

	return g, vars["s"]
}

func onBFSTreeVertex(v *ds.GraphVertex[ds.Text], n *algo.BFSNode[ds.Text]) {
	v.SetFmtAttr("label", fmt.Sprintf(`{ %s | d = %d }`, v.Label(), int(n.Distance)))
}

func onBFSUnVertex(v *ds.GraphVertex[ds.Text], n *algo.BFSNode[ds.Text]) {
	v.SetFmtAttr("label", fmt.Sprintf(`{ %s | ∞ }`, v.Label()))
	v.SetFmtAttr("fillcolor", "#ED2839")
}

func onBFSSourceVertex(v *ds.GraphVertex[ds.Text], n *algo.BFSNode[ds.Text]) {
	v.SetFmtAttr("label", fmt.Sprintf(`{ %s | source }`, v.Label()))
	v.SetFmtAttr("penwidth", "1.7")
	v.SetFmtAttr("color", "#000000")
}

func onBFSTreeEdge(e *ds.GraphEdge[ds.Text]) {
	e.SetFmtAttr("penwidth", "3.0")
}

func main() {
	// build graph, establish a source vertex
	g, src := buildBFSInput()
	ex := viz.NewExporter(g)

	ex.DefaultGraphFmt = defaultGraphFmt
	ex.DefaultVertexFmt = defaultVertexFmt
	ex.DefaultEdgeFmt = defaultEdgeFmt

	fIn, err := os.Create(fileIn)

	if err != nil {
		panic(err)
	}

	defer fIn.Close()

	// export the input graph
	ex.Export(fIn)

	// run BFS with the given source
	tree, err := algo.BFS(g, src)

	if err != nil {
		panic(err)
	}

	fOut, err := os.Create(fileOut)

	if err != nil {
		panic(err)
	}

	defer fOut.Close()

	vi := viz.NewBFSViz(g, tree, src)

	// set the desired custom formatting
	vi.DefaultGraphFmt = defaultGraphFmt
	vi.DefaultVertexFmt = defaultVertexFmt
	vi.DefaultEdgeFmt = defaultEdgeFmt
	vi.OnTreeVertex = onBFSTreeVertex
	vi.OnSourceVertex = onBFSSourceVertex
	vi.OnUnVertex = onBFSUnVertex
	vi.OnTreeEdge = onBFSTreeEdge

	// annotate the input graph with the result of the BFS,
	// then export the annotated version
	err = vi.Export(fOut)

	if err != nil {
		panic(err)
	}
}
