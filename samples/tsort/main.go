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

var theme viz.Theme[ds.Text] = viz.LightBreezeTheme[ds.Text]{}
var input = ut.UDGDress

func buildInput() *ds.Graph[ds.Text] {
	g, _, err := ds.NewTextParser().Parse(input)

	if err != nil {
		panic(err)
	}

	return g
}

func main() {
	g := buildInput()
	ex := viz.NewExporter(g)

	viz.SetTheme(ex, theme)

	fIn, err := os.Create(fileIn)

	if err != nil {
		panic(err)
	}

	defer fIn.Close()

	// export the input graph
	ex.Export(fIn)

	ord, err := algo.TSort(g)

	if err != nil {
		panic(err)
	}

	fOut, err := os.Create(fileOut)

	if err != nil {
		panic(err)
	}

	defer fOut.Close()

	vi := viz.NewTSortViz(g, ord)

	vi.Theme = theme

	vi.OnVertex = func(v *ds.GraphVertex[ds.Text], rank int) {
		v.SetFmtAttr("label", fmt.Sprintf(`%s | %d`, v.Label(), rank))
		// v.SetFmtAttr("pos", fmt.Sprintf("%d,10!", rank+5))
	}

	// annotate the input graph with the result of the DFS,
	// then export the annotated version
	if err := vi.Export(fOut); err != nil {
		panic(err)
	}
}
