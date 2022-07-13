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

var soloVertex = "7#"

var (
	defaultGraphFmt = ds.FmtAttrs{
		"layout":  "dot",
		"nodesep": "0.8",
		"ranksep": "0.5",
		"pad":     "0.2",
	}

	defaultVertexFmt = ds.FmtAttrs{
		"shape":     "Mrecord",
		"style":     "filled",
		"fillcolor": "#7289da",
		"fontcolor": "#ffffff",
		"color":     "#ffffff",
		"penwidth":  "1.1",
	}

	defaultEdgeFmt = ds.FmtAttrs{
		"penwidth": "1.2",
	}
)

func buildDFSInput() *ds.Graph[ds.Text] {
	g, _, err := ds.NewTextParser().Parse(ut.BasicUDG + soloVertex)

	if err != nil {
		panic(err)
	}

	return g
}

func main() {
	g := buildDFSInput()
	ex := viz.NewExporter(g)

	ex.DefaultGraphFmt = defaultGraphFmt
	ex.DefaultVertexFmt = defaultVertexFmt
	ex.DefaultEdgeFmt = defaultEdgeFmt

	fIn, err := os.Create(fileIn)

	if err != nil {
		panic(err)
	}

	defer fIn.Close()

	// export the input graph
	ex.Export(fIn)

	fst, edges, err := algo.DFS(g)

	if err != nil {
		panic(err)
	}

	fmt.Println("DFS Forest...")
	for vtx, node := range fst {
		s := fmt.Sprintf("%s %d/%d", vtx.Label(), node.Discovery, node.Finish)

		if node.Parent != nil {
			s = fmt.Sprintf("%s p=%s", s, node.Parent.Label())
		}

		fmt.Println(s)
	}

	fmt.Println("Forward edges...")
	for _, e := range edges.Forward {
		fmt.Printf("%s -> %s\n", e.Src.Label(), e.Dst.Label())
	}

	fmt.Println("Back edges...")
	for _, e := range edges.Back {
		fmt.Printf("%s -> %s\n", e.Src.Label(), e.Dst.Label())
	}

	fmt.Println("Cross edges...")
	for _, e := range edges.Cross {
		fmt.Printf("%s -> %s\n", e.Src.Label(), e.Dst.Label())
	}
}
