package testutils

import (
	"embed"
)

//go:embed testdata
var fixFS embed.FS

// // Basic Unweighted Directed Graph.
var UDGBasic = LoadFixture("testdata/graphs/clrs_22_2_a.gga")

// // Basic Unweighted Undirected Graph
var UUGBasic = LoadFixture("testdata/graphs/clrs_22_3.gga")

// LoadFixture loads a fixture from a file
func LoadFixture(path string) string {
	bs, err := fixFS.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return string(bs)
}
