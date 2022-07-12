//go:build !test

package main

import (
	"fmt"
	"os"

	"github.com/vc-souza/gga/algo"
	"github.com/vc-souza/gga/ds"
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
	a := ds.Text("A")
	r := ds.Text("R")
	s := ds.Text("S")
	t := ds.Text("T")
	u := ds.Text("U")
	v := ds.Text("V")
	w := ds.Text("W")
	x := ds.Text("X")
	y := ds.Text("Y")

	g := ds.NewUndirectedGraph[ds.Text]()

	g.AddVertex(&a)

	g.AddUnweightedEdge(&r, &s)
	g.AddUnweightedEdge(&r, &v)

	g.AddUnweightedEdge(&s, &r)
	g.AddUnweightedEdge(&s, &w)

	g.AddUnweightedEdge(&t, &u)
	g.AddUnweightedEdge(&t, &w)
	g.AddUnweightedEdge(&t, &x)

	g.AddUnweightedEdge(&u, &t)
	g.AddUnweightedEdge(&u, &x)
	g.AddUnweightedEdge(&u, &y)

	g.AddUnweightedEdge(&v, &r)

	g.AddUnweightedEdge(&w, &s)
	g.AddUnweightedEdge(&w, &t)
	g.AddUnweightedEdge(&w, &x)

	g.AddUnweightedEdge(&x, &t)
	g.AddUnweightedEdge(&x, &u)
	g.AddUnweightedEdge(&x, &w)
	g.AddUnweightedEdge(&x, &y)

	g.AddUnweightedEdge(&y, &u)
	g.AddUnweightedEdge(&y, &x)

	return g, &s
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
