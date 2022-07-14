package algo

import (
	"math"
	"testing"

	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestBFS_directed(t *testing.T) {
	g, vars, err := ds.NewTextParser().Parse(ut.UDGBasic)

	ut.AssertEqual(t, true, err == nil)

	v1 := vars["1"]
	v2 := vars["2"]
	v3 := vars["3"]
	v4 := vars["4"]
	v5 := vars["5"]
	v6 := vars["6"]

	tree, err := BFS(g, v3)

	ut.AssertEqual(t, true, err == nil)

	ut.AssertEqual(t, true, math.IsInf(tree[v1].Distance, 1))
	ut.AssertEqual(t, nil, tree[v1].Parent)

	ut.AssertEqual(t, 3, tree[v2].Distance)
	ut.AssertEqual(t, v4, tree[v2].Parent)

	ut.AssertEqual(t, 0, tree[v3].Distance)
	ut.AssertEqual(t, nil, tree[v3].Parent)

	ut.AssertEqual(t, 2, tree[v4].Distance)
	ut.AssertEqual(t, v5, tree[v4].Parent)

	ut.AssertEqual(t, 1, tree[v5].Distance)
	ut.AssertEqual(t, v3, tree[v5].Parent)

	ut.AssertEqual(t, 1, tree[v6].Distance)
	ut.AssertEqual(t, v3, tree[v6].Parent)
}

func TestBFS_undirected(t *testing.T) {
	g, vars, err := ds.NewTextParser().Parse(ut.UUGBasic)

	ut.AssertEqual(t, true, err == nil)

	vR := vars["r"]
	vS := vars["s"]
	vT := vars["t"]
	vU := vars["u"]
	vV := vars["v"]
	vW := vars["w"]
	vX := vars["x"]
	vY := vars["y"]

	tree, err := BFS(g, vU)

	ut.AssertEqual(t, true, err == nil)

	ut.AssertEqual(t, 4, tree[vR].Distance)
	ut.AssertEqual(t, vS, tree[vR].Parent)

	ut.AssertEqual(t, 3, tree[vS].Distance)
	ut.AssertEqual(t, vW, tree[vS].Parent)

	ut.AssertEqual(t, 1, tree[vT].Distance)
	ut.AssertEqual(t, vU, tree[vT].Parent)

	ut.AssertEqual(t, 0, tree[vU].Distance)
	ut.AssertEqual(t, nil, tree[vU].Parent)

	ut.AssertEqual(t, 5, tree[vV].Distance)
	ut.AssertEqual(t, vR, tree[vV].Parent)

	ut.AssertEqual(t, 2, tree[vW].Distance)
	ut.AssertEqual(t, vT, tree[vW].Parent)

	ut.AssertEqual(t, 1, tree[vX].Distance)
	ut.AssertEqual(t, vU, tree[vX].Parent)

	ut.AssertEqual(t, 1, tree[vY].Distance)
	ut.AssertEqual(t, vU, tree[vY].Parent)
}
