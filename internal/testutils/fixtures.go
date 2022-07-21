package testutils

import (
	"embed"
)

//go:embed testdata
var fixFS embed.FS

// Unweighted Directed Graph with simple layout.
var UDGSimple = loadFixture("testdata/graphs/clrs_22_2_a.gga")

// Unweighted Directed Graph with non-trivial dependencies.
var UDGDeps = loadFixture("testdata/graphs/clrs_22_6.gga")

// Unweighted Directed Graph with complex dependencies.
var UDGClx = loadFixture("testdata/graphs/clrs_22_8.gga")

// Unweighted Directed Graph containing dress order data.
var UDGDress = loadFixture("testdata/graphs/clrs_22_7.gga")

// Unweighted Undirected Graph with simple layout.
var UUGSimple = loadFixture("testdata/graphs/clrs_22_3.gga")

// Unweighted Undirected Graph that is not fully connected.
var UUGDisc = loadFixture("testdata/graphs/clrs_21_1_a.gga")

// Weighted Undirected Graph with simple layout.
var WUGSimple = loadFixture("testdata/graphs/clrs_23_1.gga")

// loadFixture loads a fixture from a file
func loadFixture(path string) string {
	bs, err := fixFS.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return string(bs)
}
