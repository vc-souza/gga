//go:build !test

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/vc-souza/gga/algo"
	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
	"github.com/vc-souza/gga/viz"
)

const (
	fileIn     = "GSCC-before.dot"
	fileOutSCC = "GSCC-after-SCC.dot"
	fileOut    = "GSCC-after.dot"
)

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

func exportEnd(v viz.AlgoViz, path string) {
	fOut, err := os.Create(path)

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
	attrs["rankdir"] = "LR"
}

func main() {
	g := input()

	exportStart(g)

	gscc, sccs, err := algo.GSCC(g)

	if err != nil {
		panic(err)
	}

	viSCC := viz.NewSCCViz(g, sccs, viz.Themes.LightBreeze)

	viSCC.OnSCCVertex = func(v int, c int) {
		label := fmt.Sprintf(`{ %s | cc #%d }`, viSCC.Graph.V[v].Label(), c)
		viSCC.Graph.V[v].SetFmtAttr("label", label)
	}

	viSCC.OnSCCEdge = func(v int, e int, c int) {
		viSCC.Graph.V[v].E[e].SetFmtAttr("penwidth", "2.0")
	}

	viSCC.OnCrossSCCEdge = func(v int, e int, cSrc, cDst int) {
		viSCC.Graph.V[v].E[e].SetFmtAttr("penwidth", "0.5")
		viSCC.Graph.V[v].E[e].SetFmtAttr("style", "dotted")
	}

	vi := viz.NewGSCCViz(gscc, customTheme{})

	vi.OnGSCCVertex = func(v int) {
		items := vi.Graph.V[v].Item.(ds.Group).Items

		s := make([]string, 0, len(items))

		for _, item := range items {
			s = append(s, item.Label())
		}

		label := fmt.Sprintf("{ %s }", strings.Join(s, " | "))

		vi.Graph.V[v].SetFmtAttr("label", label)
	}

	exportEnd(viSCC, fileOutSCC)
	exportEnd(vi, fileOut)
}
