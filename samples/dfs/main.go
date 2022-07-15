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
	fileIn  = "DFS-before.dot"
	fileOut = "DFS-after.dot"
)

func input() *ds.Graph[ds.Text] {
	g, _, err := ds.NewTextParser().Parse(ut.UDGSimple + "\n7#")

	if err != nil {
		panic(err)
	}

	return g
}

func start(g *ds.Graph[ds.Text]) {
	fIn, err := os.Create(fileIn)

	if err != nil {
		panic(err)
	}

	defer fIn.Close()

	viz.Snapshot(g, fIn, viz.Themes.LightBreeze)
}

func end(v viz.AlgoViz[ds.Text]) {
	fOut, err := os.Create(fileOut)

	if err != nil {
		panic(err)
	}

	defer fOut.Close()

	// annotate the input graph with the result of the DFS,
	// then export the annotated version
	if err := viz.ExportViz(v, fOut); err != nil {
		panic(err)
	}
}

func main() {
	g := input()

	start(g)

	fst, edges, err := algo.DFS(g, true)

	if err != nil {
		panic(err)
	}

	vi := viz.NewDFSViz(g, fst, edges, viz.Themes.LightBreeze)

	vi.OnTreeVertex = func(v *ds.GraphVertex[ds.Text], n *algo.DFNode[ds.Text]) {
		v.SetFmtAttr("label", fmt.Sprintf(` %s | { d = %d | f = %d }`, v.Label(), n.Discovery, n.Finish))
	}

	vi.OnRootVertex = func(v *ds.GraphVertex[ds.Text], n *algo.DFNode[ds.Text]) {
		v.SetFmtAttr("penwidth", "1.7")
		v.SetFmtAttr("color", "#000000")
	}

	vi.OnTreeEdge = func(e *ds.GraphEdge[ds.Text]) {
		e.SetFmtAttr("penwidth", "3.0")
	}

	vi.OnForwardEdge = func(e *ds.GraphEdge[ds.Text]) {
		e.SetFmtAttr("label", "F")
	}

	vi.OnBackEdge = func(e *ds.GraphEdge[ds.Text]) {
		e.SetFmtAttr("label", "B")
	}

	vi.OnCrossEdge = func(e *ds.GraphEdge[ds.Text]) {
		e.SetFmtAttr("label", "C")
	}

	end(vi)
}
