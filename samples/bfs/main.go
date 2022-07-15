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

func input() (*ds.Graph[ds.Text], *ds.Text) {
	g, vars, err := ds.NewTextParser().Parse(ut.UUGSimple + "\na#")

	if err != nil {
		panic(err)
	}

	return g, vars["s"]
}

func exportStart(g *ds.Graph[ds.Text]) {
	fIn, err := os.Create(fileIn)

	if err != nil {
		panic(err)
	}

	defer fIn.Close()

	viz.Snapshot(g, fIn, viz.Themes.LightBreeze)
}

func exportEnd(v viz.AlgoViz[ds.Text]) {
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
	g, src := input()

	exportStart(g)

	tree, err := algo.BFS(g, src)

	if err != nil {
		panic(err)
	}

	vi := viz.NewBFSViz(g, tree, src, viz.Themes.LightBreeze)

	vi.OnTreeVertex = func(v *ds.GraphVertex[ds.Text], n *algo.BFNode[ds.Text]) {
		v.SetFmtAttr("label", fmt.Sprintf(`{ %s | d = %d }`, v.Label(), int(n.Distance)))
	}

	vi.OnSourceVertex = func(v *ds.GraphVertex[ds.Text], n *algo.BFNode[ds.Text]) {
		v.SetFmtAttr("label", fmt.Sprintf(`{ %s | source }`, v.Label()))
		v.SetFmtAttr("penwidth", "1.7")
		v.SetFmtAttr("color", "#000000")
	}

	vi.OnUnVertex = func(v *ds.GraphVertex[ds.Text], n *algo.BFNode[ds.Text]) {
		v.SetFmtAttr("label", fmt.Sprintf(`{ %s | ∞ }`, v.Label()))
		v.SetFmtAttr("fillcolor", "#ED2839")
	}

	vi.OnTreeEdge = func(e *ds.GraphEdge[ds.Text]) {
		e.SetFmtAttr("penwidth", "3.0")
	}

	exportEnd(vi)
}
