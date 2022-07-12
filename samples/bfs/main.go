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
	bfsFileIn  = "BFS-in.dot"
	bfsFileOut = "BFS-out.dot"
)

var (
	defaultGraphFmt = ds.FmtAttrs{
		"layout": "dot",
	}

	defaultVertexFmt = ds.FmtAttrs{
		"shape": "circle",
		"style": "filled",
	}

	defaultEdgeFmt = ds.FmtAttrs{
		"pencolor": "#000000",
	}
)

func buildBFSInput() (*ds.Graph[ds.Text], *ds.Text) {
	g, vars, err := ds.NewTextParser().Parse(ut.BasicUG)

	if err != nil {
		panic(err)
	}

	return g, vars["s"]
}

func onBFSTreeVertex(v *ds.GraphVertex[ds.Text], n *algo.BFSNode[ds.Text]) {
	v.SetFmtAttr("label", fmt.Sprintf(`%s\nd=%d`, v.Label(), int(n.Distance)))
	v.SetFmtAttr("fillcolor", "#000000")
	v.SetFmtAttr("fontcolor", "#ffffff")
	v.SetFmtAttr("shape", "doublecircle")
}

func onBFSUnVertex(v *ds.GraphVertex[ds.Text], n *algo.BFSNode[ds.Text]) {
	v.SetFmtAttr("fillcolor", "#ff0000")
	v.SetFmtAttr("fontcolor", "#ffffff")
}

func onBFSSourceVertex(v *ds.GraphVertex[ds.Text], n *algo.BFSNode[ds.Text]) {
	v.SetFmtAttr("shape", "circle")
	v.SetFmtAttr("pos", "0,0!")
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

	fIn, err := os.Create(bfsFileIn)

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

	fOut, err := os.Create(bfsFileOut)

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
