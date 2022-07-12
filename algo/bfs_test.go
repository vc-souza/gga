package algo

import (
	"math"
	"testing"

	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestBFS_directed(t *testing.T) {
	g := ds.NewDirectedGraph[ut.ID]()

	// CLRS 3rd Edition, 22.2-1
	v1 := ut.ID("1")
	v2 := ut.ID("2")
	v3 := ut.ID("3")
	v4 := ut.ID("4")
	v5 := ut.ID("5")
	v6 := ut.ID("6")

	g.AddVertex(&v1)
	g.AddVertex(&v2)
	g.AddVertex(&v3)
	g.AddVertex(&v4)
	g.AddVertex(&v5)
	g.AddVertex(&v6)

	g.AddUnweightedEdge(&v1, &v2)
	g.AddUnweightedEdge(&v1, &v4)
	g.AddUnweightedEdge(&v2, &v5)
	g.AddUnweightedEdge(&v3, &v5)
	g.AddUnweightedEdge(&v3, &v6)
	g.AddUnweightedEdge(&v4, &v2)
	g.AddUnweightedEdge(&v5, &v4)
	g.AddUnweightedEdge(&v6, &v6)

	tree, err := BFS(g, &v3)

	ut.AssertEqual(t, true, err == nil)

	ut.AssertEqual(t, true, math.IsInf(tree[&v1].Distance, 1))
	ut.AssertEqual(t, nil, tree[&v1].Parent)

	ut.AssertEqual(t, 3, tree[&v2].Distance)
	ut.AssertEqual(t, &v4, tree[&v2].Parent)

	ut.AssertEqual(t, 0, tree[&v3].Distance)
	ut.AssertEqual(t, nil, tree[&v3].Parent)

	ut.AssertEqual(t, 2, tree[&v4].Distance)
	ut.AssertEqual(t, &v5, tree[&v4].Parent)

	ut.AssertEqual(t, 1, tree[&v5].Distance)
	ut.AssertEqual(t, &v3, tree[&v5].Parent)

	ut.AssertEqual(t, 1, tree[&v6].Distance)
	ut.AssertEqual(t, &v3, tree[&v6].Parent)
}

func TestBFS_undirected(t *testing.T) {
	g := ds.NewUndirectedGraph[ut.ID]()

	// CLRS 3rd Edition, 22.2-2
	vR := ut.ID("r")
	vS := ut.ID("s")
	vT := ut.ID("t")
	vU := ut.ID("u")
	vV := ut.ID("v")
	vW := ut.ID("w")
	vX := ut.ID("x")
	vY := ut.ID("y")

	g.AddVertex(&vR)
	g.AddVertex(&vS)
	g.AddVertex(&vT)
	g.AddVertex(&vU)
	g.AddVertex(&vV)
	g.AddVertex(&vW)
	g.AddVertex(&vX)
	g.AddVertex(&vY)

	g.AddUnweightedEdge(&vR, &vS)
	g.AddUnweightedEdge(&vR, &vV)

	g.AddUnweightedEdge(&vS, &vR)
	g.AddUnweightedEdge(&vS, &vW)

	g.AddUnweightedEdge(&vT, &vU)
	g.AddUnweightedEdge(&vT, &vW)
	g.AddUnweightedEdge(&vT, &vX)

	g.AddUnweightedEdge(&vU, &vT)
	g.AddUnweightedEdge(&vU, &vX)
	g.AddUnweightedEdge(&vU, &vY)

	g.AddUnweightedEdge(&vV, &vR)

	g.AddUnweightedEdge(&vW, &vS)
	g.AddUnweightedEdge(&vW, &vT)
	g.AddUnweightedEdge(&vW, &vX)

	g.AddUnweightedEdge(&vX, &vT)
	g.AddUnweightedEdge(&vX, &vU)
	g.AddUnweightedEdge(&vX, &vW)
	g.AddUnweightedEdge(&vX, &vY)

	g.AddUnweightedEdge(&vY, &vU)
	g.AddUnweightedEdge(&vY, &vX)

	tree, err := BFS(g, &vU)

	ut.AssertEqual(t, true, err == nil)

	ut.AssertEqual(t, 4, tree[&vR].Distance)
	ut.AssertEqual(t, &vS, tree[&vR].Parent)

	ut.AssertEqual(t, 3, tree[&vS].Distance)
	ut.AssertEqual(t, &vW, tree[&vS].Parent)

	ut.AssertEqual(t, 1, tree[&vT].Distance)
	ut.AssertEqual(t, &vU, tree[&vT].Parent)

	ut.AssertEqual(t, 0, tree[&vU].Distance)
	ut.AssertEqual(t, nil, tree[&vU].Parent)

	ut.AssertEqual(t, 5, tree[&vV].Distance)
	ut.AssertEqual(t, &vR, tree[&vV].Parent)

	ut.AssertEqual(t, 2, tree[&vW].Distance)
	ut.AssertEqual(t, &vT, tree[&vW].Parent)

	ut.AssertEqual(t, 1, tree[&vX].Distance)
	ut.AssertEqual(t, &vU, tree[&vX].Parent)

	ut.AssertEqual(t, 1, tree[&vY].Distance)
	ut.AssertEqual(t, &vU, tree[&vY].Parent)
}
