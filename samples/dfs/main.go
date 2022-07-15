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

var input = ut.UDGSimple + "\n7#"

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

	viz.SetTheme(ex, viz.Themes.LightBreeze)

	fIn, err := os.Create(fileIn)

	if err != nil {
		panic(err)
	}

	defer fIn.Close()

	// export the input graph
	ex.Export(fIn)

	fst, edges, err := algo.DFS(g, true)

	if err != nil {
		panic(err)
	}

	fOut, err := os.Create(fileOut)

	if err != nil {
		panic(err)
	}

	defer fOut.Close()

	vi := viz.NewDFSViz(g, fst, edges)

	vi.Theme = viz.Themes.LightBreeze

	vi.OnTreeVertex = func(v *ds.GraphVertex[ds.Text], n *algo.DFNode[ds.Text]) {
		v.SetFmtAttr("label", fmt.Sprintf(` %s | { d = %d | f = %d }`, v.Label(), n.Discovery, n.Finish))
	}

	vi.OnRootVertex = func(v *ds.GraphVertex[ds.Text], n *algo.DFNode[ds.Text]) {
		v.SetFmtAttr("penwidth", "1.7")
		v.SetFmtAttr("color", "#000000")
	}

	vi.OnTreeEdge = func(e *ds.GraphEdge[ds.Text]) {
		e.SetFmtAttr("penwidth", "3.0")
	}

	vi.OnForwardEdge = func(e *ds.GraphEdge[ds.Text]) {
		e.SetFmtAttr("label", "F")
	}

	vi.OnBackEdge = func(e *ds.GraphEdge[ds.Text]) {
		e.SetFmtAttr("label", "B")
	}

	vi.OnCrossEdge = func(e *ds.GraphEdge[ds.Text]) {
		e.SetFmtAttr("label", "C")
	}

	// annotate the input graph with the result of the DFS,
	// then export the annotated version
	if err := vi.Export(fOut); err != nil {
		panic(err)
	}
}
