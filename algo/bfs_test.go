package algo

import (
	"math"
	"testing"

	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestBFS_directed(t *testing.T) {
	g, vars, err := ds.NewTextParser().Parse(ut.UDGSimple)

	ut.Equal(t, true, err == nil)

	v1 := vars["1"]
	v2 := vars["2"]
	v3 := vars["3"]
	v4 := vars["4"]
	v5 := vars["5"]
	v6 := vars["6"]

	tree, err := BFS(g, v3)

	ut.Equal(t, true, err == nil)

	ut.Equal(t, true, math.IsInf(tree[v1].Distance, 1))
	ut.Equal(t, nil, tree[v1].Parent)

	ut.Equal(t, 3, tree[v2].Distance)
	ut.Equal(t, v4, tree[v2].Parent)

	ut.Equal(t, 0, tree[v3].Distance)
	ut.Equal(t, nil, tree[v3].Parent)

	ut.Equal(t, 2, tree[v4].Distance)
	ut.Equal(t, v5, tree[v4].Parent)

	ut.Equal(t, 1, tree[v5].Distance)
	ut.Equal(t, v3, tree[v5].Parent)

	ut.Equal(t, 1, tree[v6].Distance)
	ut.Equal(t, v3, tree[v6].Parent)
}

func TestBFS_undirected(t *testing.T) {
	g, vars, err := ds.NewTextParser().Parse(ut.UUGSimple)

	ut.Equal(t, true, err == nil)

	vR := vars["r"]
	vS := vars["s"]
	vT := vars["t"]
	vU := vars["u"]
	vV := vars["v"]
	vW := vars["w"]
	vX := vars["x"]
	vY := vars["y"]

	tree, err := BFS(g, vU)

	ut.Equal(t, true, err == nil)

	ut.Equal(t, 4, tree[vR].Distance)
	ut.Equal(t, vS, tree[vR].Parent)

	ut.Equal(t, 3, tree[vS].Distance)
	ut.Equal(t, vW, tree[vS].Parent)

	ut.Equal(t, 1, tree[vT].Distance)
	ut.Equal(t, vU, tree[vT].Parent)

	ut.Equal(t, 0, tree[vU].Distance)
	ut.Equal(t, nil, tree[vU].Parent)

	ut.Equal(t, 5, tree[vV].Distance)
	ut.Equal(t, vR, tree[vV].Parent)

	ut.Equal(t, 2, tree[vW].Distance)
	ut.Equal(t, vT, tree[vW].Parent)

	ut.Equal(t, 1, tree[vX].Distance)
	ut.Equal(t, vU, tree[vX].Parent)

	ut.Equal(t, 1, tree[vY].Distance)
	ut.Equal(t, vU, tree[vY].Parent)
}
