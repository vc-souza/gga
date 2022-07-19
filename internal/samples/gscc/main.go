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

func input() *ds.G[ds.Text] {
	g, _, err := ds.Parse(ut.UDGDeps)

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

func exportEnd[T ds.Item](v viz.AlgoViz[T], path string) {
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

	viSCC.OnSCCVertex = func(v *ds.GV[ds.Text], c int) {
		v.SetFmtAttr("label", fmt.Sprintf(`{ %s | cc #%d }`, v.Label(), c))
	}

	viSCC.OnSCCEdge = func(e *ds.GE[ds.Text], c int) {
		e.SetFmtAttr("penwidth", "2.0")
	}

	viSCC.OnCrossSCCEdge = func(e *ds.GE[ds.Text], cSrc, cDst int) {
		e.SetFmtAttr("penwidth", "0.5")
		e.SetFmtAttr("style", "dotted")
	}

	vi := viz.NewGSCCViz(gscc, customTheme{})

	vi.OnGSCCVertex = func(v *ds.GV[ds.Group[ds.Text]]) {
		s := make([]string, 0, len(v.Ptr.Items))

		for _, item := range v.Ptr.Items {
			s = append(s, item.Label())
		}

		v.SetFmtAttr("label", fmt.Sprintf("{ %s }", strings.Join(s, " | ")))
	}

	exportEnd[ds.Text](viSCC, fileOutSCC)
	exportEnd[ds.Group[ds.Text]](vi, fileOut)
}
