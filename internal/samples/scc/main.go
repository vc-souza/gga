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
	fileIn  = "SCC-before.dot"
	fileOut = "SCC-after.dot"
)

var algos = map[string]algo.SCCAlgo{
	"kosaraju": algo.SCCKosaraju,
	"tarjan":   algo.SCCTarjan,
}

func input() *ds.G {
	g, _, err := ds.Parse(ut.UDGDeps)

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
	if len(os.Args) < 2 {
		panic("Too few args!")
	}

	if len(os.Args) > 2 {
		panic("Too many args!")
	}

	f, ok := algos[os.Args[1]]

	if !ok {
		panic(fmt.Sprintf("invalid option '%s'", os.Args[1]))
	}

	g := input()

	exportStart(g)

	sccs, err := f(g)

	if err != nil {
		panic(err)
	}

	vi := viz.NewSCCViz(g, sccs, viz.Themes.LightBreeze)

	vi.OnSCCVertex = func(v int, c int) {
		label := fmt.Sprintf(`{ %s | cc #%d }`, vi.Graph.V[v].Label(), c)
		vi.Graph.V[v].SetFmtAttr("label", label)
	}

	vi.OnSCCEdge = func(v int, e int, c int) {
		vi.Graph.V[v].E[e].SetFmtAttr("penwidth", "2.0")
	}

	vi.OnCrossSCCEdge = func(v int, e int, cSrc, cDst int) {
		vi.Graph.V[v].E[e].SetFmtAttr("penwidth", "0.5")
		vi.Graph.V[v].E[e].SetFmtAttr("style", "dotted")
	}

	exportEnd(vi)
}
