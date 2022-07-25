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
	fileIn  = "MST-before.dot"
	fileOut = "MST-after.dot"
)

var algos = map[string]algo.MSTAlgo{
	"kruskal": algo.MSTKruskal,
	"prim":    algo.MSTPrim,
}

func input() *ds.G {
	g, _, err := ds.Parse(ut.WUGSimple)

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

	viz.Snapshot(g, fIn, customTheme{})
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

type customTheme struct {
	viz.LightBreezeTheme
}

func (t customTheme) SetGraphFmt(attrs ds.FAttrs) {
	t.LightBreezeTheme.SetGraphFmt(attrs)
	attrs["nodesep"] = "0.3"
	attrs["ranksep"] = "0.2"
}

func (t customTheme) SetEdgeFmt(attrs ds.FAttrs) {
	t.LightBreezeTheme.SetEdgeFmt(attrs)
	attrs["fontsize"] = "10.0"
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

	mst, err := f(g)

	if err != nil {
		panic(err)
	}

	vi := viz.NewMSTViz(g, mst, customTheme{})

	vi.OnMSTEdge = func(v int, e int) {
		vi.Graph.V[v].E[e].SetFmtAttr("penwidth", "3.0")
	}

	exportEnd(vi)
}
