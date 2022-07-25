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

func input() *ds.G {
	g, _, err := ds.Parse(ut.UDGSimple + "\n7#")

	if err != nil {
		panic(err)
	}

	return g
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
	g := input()

	exportStart(g)

	fst, edges, err := algo.DFS(g, true)

	if err != nil {
		panic(err)
	}

	vi := viz.NewDFSViz(g, fst, edges, viz.Themes.LightBreeze)

	vi.OnTreeVertex = func(v int, n algo.DFNode) {
		label := fmt.Sprintf(`%s | { d = %d | f = %d }`, vi.Graph.V[v].Label(), n.Discovery, n.Finish)
		vi.Graph.V[v].SetFmtAttr("label", label)
	}

	vi.OnRootVertex = func(v int, n algo.DFNode) {
		vi.Graph.V[v].SetFmtAttr("penwidth", "1.7")
		vi.Graph.V[v].SetFmtAttr("color", "#000000")
	}

	vi.OnTreeEdge = func(v int, e int) {
		vi.Graph.V[v].E[e].SetFmtAttr("penwidth", "3.0")
	}

	vi.OnForwardEdge = func(v int, e int) {
		vi.Graph.V[v].E[e].SetFmtAttr("label", "F")
	}

	vi.OnBackEdge = func(v int, e int) {
		vi.Graph.V[v].E[e].SetFmtAttr("label", "B")
	}

	vi.OnCrossEdge = func(v int, e int) {
		vi.Graph.V[v].E[e].SetFmtAttr("label", "C")
	}

	exportEnd(vi)
}
