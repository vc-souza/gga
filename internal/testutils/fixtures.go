package testutils

import (
	"embed"
)

//go:embed testdata
var fixFS embed.FS

// Unweighted Directed Graph with simple layout.
var UDGSimple = LoadFixture("testdata/graphs/clrs_22_2_a.gga")

// Unweighted Directed Graph with non-trivial dependencies.
var UDGDeps = LoadFixture("testdata/graphs/clrs_22_6.gga")

// Unweighted Directed Graph with complex dependencies.
var UDGClx = LoadFixture("testdata/graphs/clrs_22_8.gga")

// Unweighted Directed Graph containing dress order data.
var UDGDress = LoadFixture("testdata/graphs/clrs_22_7.gga")

// Unweighted Undirected Graph with simple layout.
var UUGSimple = LoadFixture("testdata/graphs/clrs_22_3.gga")

// LoadFixture loads a fixture from a file
func LoadFixture(path string) string {
	bs, err := fixFS.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return string(bs)
}
