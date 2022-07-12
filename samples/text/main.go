//go:build !test

package main

import (
	"fmt"
	"os"

	"github.com/vc-souza/gga/ds"
	"github.com/vc-souza/gga/viz"
)

var undirectedInput = `
graph
a#b:4,h:8
b#a:4,c:8,h:11
c#b:8,d:7,i:2,f:4
d#c:7,e:9,f:14
e#d:9,f:10
f#c:4,d:14,e:10,g:2
g#f:2,h:1,i:6
h#a:8,b:11,g:1,i:7
i#c:2,g:6,h:7
j#
`

var directedInput = `
digraph
1#2,4
2#5
3#5,6
4#2
5#4
6#6
7#
`

func export(src string) {
	g, _, err := new(ds.TextParser).Parse(src)

	if err != nil {
		panic(err)
	}

	viz.NewExporter(g).Export(os.Stdout)
}

func main() {
	var opt string

	fmt.Println("Pick a graph type:")
	fmt.Println("1. Digraph")
	fmt.Println("2. Graph")

	_, err := fmt.Scanln(&opt)

	if err != nil {
		panic(err)
	}

	switch opt {
	case "1":
		export(directedInput)
	case "2":
		export(undirectedInput)
	default:
		panic("invalid choice!")
	}
}
