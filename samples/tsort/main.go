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

func input() *ds.Graph[ds.Text] {
	g, _, err := ds.NewTextParser().Parse(ut.UDGDress)

	if err != nil {
		panic(err)
	}

	return g
}

func start(g *ds.Graph[ds.Text]) {
	fIn, err := os.Create(fileIn)

	if err != nil {
		panic(err)
	}

	defer fIn.Close()

	viz.Snapshot(g, fIn, viz.Themes.LightBreeze)
}

func end(v viz.AlgoViz[ds.Text]) {
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
	g := input()

	start(g)

	ord, err := algo.TSort(g)

	if err != nil {
		panic(err)
	}

	vi := viz.NewTSortViz(g, ord, viz.Themes.LightBreeze)

	vi.OnVertex = func(v *ds.GraphVertex[ds.Text], rank int) {
		v.SetFmtAttr("label", fmt.Sprintf(`%s | %d`, v.Label(), rank))
	}

	end(vi)
}
