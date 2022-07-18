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

func input() *ds.G[ds.Text] {
	g, _, err := ds.Parse(ut.UDGSimple + "\n7#")

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
	g := input()

	exportStart(g)

	fst, edges, err := algo.DFS(g, true)

	if err != nil {
		panic(err)
	}

	vi := viz.NewDFSViz(g, fst, edges, viz.Themes.LightBreeze)

	vi.OnTreeVertex = func(v *ds.GV[ds.Text], n *algo.DFNode[ds.Text]) {
		v.SetFmtAttr("label", fmt.Sprintf(`%s | { d = %d | f = %d }`, v.Label(), n.Discovery, n.Finish))
	}

	vi.OnRootVertex = func(v *ds.GV[ds.Text], n *algo.DFNode[ds.Text]) {
		v.SetFmtAttr("penwidth", "1.7")
		v.SetFmtAttr("color", "#000000")
	}

	vi.OnTreeEdge = func(e *ds.GE[ds.Text]) {
		e.SetFmtAttr("penwidth", "3.0")
	}

	vi.OnForwardEdge = func(e *ds.GE[ds.Text]) {
		e.SetFmtAttr("label", "F")
	}

	vi.OnBackEdge = func(e *ds.GE[ds.Text]) {
		e.SetFmtAttr("label", "B")
	}

	vi.OnCrossEdge = func(e *ds.GE[ds.Text]) {
		e.SetFmtAttr("label", "C")
	}

	exportEnd(vi)
}
