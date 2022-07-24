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
	fileIn  = "TSort-before.dot"
	fileOut = "TSort-after.dot"
)

func input() *ds.G {
	g, _, err := ds.Parse(ut.UDGDress)

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

type customTheme struct {
	viz.LightBreezeTheme
}

func (t customTheme) SetGraphFmt(attrs ds.FAttrs) {
	t.LightBreezeTheme.SetGraphFmt(attrs)
	attrs["nodesep"] = "0.1"
	attrs["ranksep"] = "0.2"
}

func main() {
	g := input()

	exportStart(g)

	ord, err := algo.TSort(g)

	if err != nil {
		panic(err)
	}

	vi := viz.NewTSortViz(g, ord, customTheme{})

	vi.OnVertexRank = func(v int, rank int) {
		label := fmt.Sprintf(`%s | %d`, vi.Graph.V[v].Label(), rank)
		vi.Graph.V[v].SetFmtAttr("label", label)
	}

	vi.OnOrderEdge = func(v int, e int, dst int, exists bool) {
		if exists {
			return
		}

		vi.Extra = append(
			vi.Extra,
			fmt.Sprintf(
				"%s -> %s [style=invis]",
				viz.Quoted(vi.Graph.V[v].Item),
				viz.Quoted(vi.Graph.V[dst].Item),
			),
		)
	}

	exportEnd(vi)
}
