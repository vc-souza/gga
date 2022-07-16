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

var algos = map[string]algo.SCCAlgorithm{
	"kosaraju": algo.SCCKosaraju[ds.Text],
	"tarjan":   algo.SCCTarjan[ds.Text],
}

func input() *ds.Graph[ds.Text] {
	g, _, err := ds.NewTextParser().Parse(ut.UDGDeps)

	if err != nil {
		panic(err)
	}

	return g
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

	vi.OnSCCVertex = func(v *ds.GraphVertex[ds.Text], c int) {
		v.SetFmtAttr("label", fmt.Sprintf(`{ %s | cc #%d }`, v.Label(), c))
	}

	vi.OnSCCEdge = func(e *ds.GraphEdge[ds.Text], c int) {
		e.SetFmtAttr("penwidth", "2.0")
	}

	vi.OnCrossSCCEdge = func(e *ds.GraphEdge[ds.Text], cSrc, cDst int) {
		e.SetFmtAttr("penwidth", "0.5")
		e.SetFmtAttr("style", "dotted")
	}

	exportEnd(vi)
}
