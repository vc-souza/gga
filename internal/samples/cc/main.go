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
	fileIn  = "CC-before.dot"
	fileOut = "CC-after.dot"
)

var algos = map[string]algo.CCAlgo[ds.Text]{
	"union-find": algo.CCUnionFind[ds.Text],
	"dfs":        algo.CCDFS[ds.Text],
}

func input() *ds.G[ds.Text] {
	g, _, err := ds.Parse(ut.UUGDisc)

	if err != nil {
		panic(err)
	}

	return g
}

func exportStart(g *ds.G[ds.Text]) {
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

	ccs, err := f(g)

	if err != nil {
		panic(err)
	}

	vi := viz.NewCCViz(g, ccs, viz.Themes.LightBreeze)

	vi.OnCCVertex = func(v *ds.GV[ds.Text], c int) {
		v.SetFmtAttr("label", fmt.Sprintf(`{ %s | cc #%d }`, v.Label(), c))
	}

	vi.OnCCEdge = func(e *ds.GE[ds.Text], c int) {
		e.SetFmtAttr("penwidth", "2.0")
	}

	exportEnd(vi)
}
