//go:build !test

package main

import (
	"fmt"
	"os"

	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
	"github.com/vc-souza/gga/viz"
)

func export(src string) {
	g, _, err := ds.Parse(src)

	if err != nil {
		panic(err)
	}

	viz.NewExporter().Export(g, os.Stdout)
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
		export(ut.UDGSimple)
	case "2":
		export(ut.UUGSimple)
	default:
		panic("invalid choice!")
	}
}
