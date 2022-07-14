package algo

import (
	"testing"

	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestDFS_directed(t *testing.T) {
	g, vars, err := ds.NewTextParser().Parse(ut.UDGSimple)

	ut.AssertEqual(t, true, err == nil)

	v1 := vars["1"]
	v2 := vars["2"]
	v3 := vars["3"]
	v4 := vars["4"]
	v5 := vars["5"]
	v6 := vars["6"]

	fst, tps, err := DFS(g)

	ut.AssertEqual(t, true, err == nil)

	ut.AssertEqual(t, 1, len(tps.Forward))
	ut.AssertEqual(t, 2, len(tps.Back))
	ut.AssertEqual(t, 1, len(tps.Cross))

	ut.AssertEqual(t, 1, fst[v1].Discovery)
	ut.AssertEqual(t, 8, fst[v1].Finish)
	ut.AssertEqual(t, nil, fst[v1].Parent)

	ut.AssertEqual(t, 2, fst[v2].Discovery)
	ut.AssertEqual(t, 7, fst[v2].Finish)
	ut.AssertEqual(t, v1, fst[v2].Parent)

	ut.AssertEqual(t, 9, fst[v3].Discovery)
	ut.AssertEqual(t, 12, fst[v3].Finish)
	ut.AssertEqual(t, nil, fst[v3].Parent)

	ut.AssertEqual(t, 4, fst[v4].Discovery)
	ut.AssertEqual(t, 5, fst[v4].Finish)
	ut.AssertEqual(t, v5, fst[v4].Parent)

	ut.AssertEqual(t, 3, fst[v5].Discovery)
	ut.AssertEqual(t, 6, fst[v5].Finish)
	ut.AssertEqual(t, v2, fst[v5].Parent)

	ut.AssertEqual(t, 10, fst[v6].Discovery)
	ut.AssertEqual(t, 11, fst[v6].Finish)
	ut.AssertEqual(t, v3, fst[v6].Parent)
}

func TestDFS_undirected(t *testing.T) {
	g, vars, err := ds.NewTextParser().Parse(ut.UUGSimple)

	ut.AssertEqual(t, true, err == nil)

	vR := vars["r"]
	vS := vars["s"]
	vT := vars["t"]
	vU := vars["u"]
	vV := vars["v"]
	vW := vars["w"]
	vX := vars["x"]
	vY := vars["y"]

	fst, tps, err := DFS(g)

	ut.AssertEqual(t, true, err == nil)

	ut.AssertEqual(t, 0, len(tps.Forward))
	ut.AssertEqual(t, 0, len(tps.Cross))

	// undirected graph, so two for each edge
	ut.AssertEqual(t, 6, len(tps.Back))

	ut.AssertEqual(t, 1, fst[vR].Discovery)
	ut.AssertEqual(t, 16, fst[vR].Finish)
	ut.AssertEqual(t, nil, fst[vR].Parent)

	ut.AssertEqual(t, 2, fst[vS].Discovery)
	ut.AssertEqual(t, 13, fst[vS].Finish)
	ut.AssertEqual(t, vR, fst[vS].Parent)

	ut.AssertEqual(t, 4, fst[vT].Discovery)
	ut.AssertEqual(t, 11, fst[vT].Finish)
	ut.AssertEqual(t, vW, fst[vT].Parent)

	ut.AssertEqual(t, 5, fst[vU].Discovery)
	ut.AssertEqual(t, 10, fst[vU].Finish)
	ut.AssertEqual(t, vT, fst[vU].Parent)

	ut.AssertEqual(t, 14, fst[vV].Discovery)
	ut.AssertEqual(t, 15, fst[vV].Finish)
	ut.AssertEqual(t, vR, fst[vV].Parent)

	ut.AssertEqual(t, 3, fst[vW].Discovery)
	ut.AssertEqual(t, 12, fst[vW].Finish)
	ut.AssertEqual(t, vS, fst[vW].Parent)

	ut.AssertEqual(t, 6, fst[vX].Discovery)
	ut.AssertEqual(t, 9, fst[vX].Finish)
	ut.AssertEqual(t, vU, fst[vX].Parent)

	ut.AssertEqual(t, 7, fst[vY].Discovery)
	ut.AssertEqual(t, 8, fst[vY].Finish)
	ut.AssertEqual(t, vX, fst[vY].Parent)
}
