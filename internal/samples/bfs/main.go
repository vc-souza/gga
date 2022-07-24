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

func input() (*ds.G, int) {
	g, idx, err := ds.Parse(ut.UUGSimple + "\na#")

	if err != nil {
		panic(err)
	}

	return g, idx("s")
}

func exportStart(g *ds.G) {
	fIn, err := os.Create(fileIn)

	if err != nil {
		panic(err)
	}

	defer fIn.Close()

	viz.Snapshot(g, fIn, viz.Themes.LightBreeze)
}

func exportEnd(v viz.AlgoViz) {
	fOut, err := os.Create(fileOut)

	if err != nil {
		panic(err)
	}

	defer fOut.Close()

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

	vi.OnTreeVertex = func(v int, n algo.BFNode) {
		label := fmt.Sprintf(`{ %s | d = %d }`, vi.Graph.V[v].Label(), int(n.Distance))
		vi.Graph.V[v].SetFmtAttr("label", label)
	}

	vi.OnSourceVertex = func(v int, n algo.BFNode) {
		label := fmt.Sprintf(`{ %s | source }`, vi.Graph.V[v].Label())

		vi.Graph.V[v].SetFmtAttr("label", label)
		vi.Graph.V[v].SetFmtAttr("penwidth", "1.7")
		vi.Graph.V[v].SetFmtAttr("color", "#000000")
	}

	vi.OnUnVertex = func(v int, n algo.BFNode) {
		vi.Graph.V[v].SetFmtAttr("label", fmt.Sprintf(`{ %s | ∞ }`, vi.Graph.V[v].Label()))
		vi.Graph.V[v].SetFmtAttr("fillcolor", "#ED2839")
	}

	vi.OnTreeEdge = func(v int, e int) {
		vi.Graph.V[v].E[e].SetFmtAttr("penwidth", "3.0")
	}

	exportEnd(vi)
}
