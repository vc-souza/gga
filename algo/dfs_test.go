package algo

import (
	"fmt"
	"testing"

	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestDFS_directed(t *testing.T) {
	for _, classify := range []bool{true, false} {
		t.Run(fmt.Sprintf("classify: %v", classify), func(t *testing.T) {
			g, vars, err := ds.NewTextParser().Parse(ut.UDGSimple)

			ut.AssertEQ(t, true, err == nil)

			v1 := vars["1"]
			v2 := vars["2"]
			v3 := vars["3"]
			v4 := vars["4"]
			v5 := vars["5"]
			v6 := vars["6"]

			fst, tps, err := DFS(g, classify)

			ut.AssertEQ(t, true, err == nil)

			if classify {
				ut.AssertEQ(t, 1, len(tps.Forward))
				ut.AssertEQ(t, 2, len(tps.Back))
				ut.AssertEQ(t, 1, len(tps.Cross))
			} else {
				ut.AssertEQ(t, 0, len(tps.Forward))
				ut.AssertEQ(t, 0, len(tps.Back))
				ut.AssertEQ(t, 0, len(tps.Cross))
			}

			ut.AssertEQ(t, 1, fst[v1].Discovery)
			ut.AssertEQ(t, 8, fst[v1].Finish)
			ut.AssertEQ(t, nil, fst[v1].Parent)

			ut.AssertEQ(t, 2, fst[v2].Discovery)
			ut.AssertEQ(t, 7, fst[v2].Finish)
			ut.AssertEQ(t, v1, fst[v2].Parent)

			ut.AssertEQ(t, 9, fst[v3].Discovery)
			ut.AssertEQ(t, 12, fst[v3].Finish)
			ut.AssertEQ(t, nil, fst[v3].Parent)

			ut.AssertEQ(t, 4, fst[v4].Discovery)
			ut.AssertEQ(t, 5, fst[v4].Finish)
			ut.AssertEQ(t, v5, fst[v4].Parent)

			ut.AssertEQ(t, 3, fst[v5].Discovery)
			ut.AssertEQ(t, 6, fst[v5].Finish)
			ut.AssertEQ(t, v2, fst[v5].Parent)

			ut.AssertEQ(t, 10, fst[v6].Discovery)
			ut.AssertEQ(t, 11, fst[v6].Finish)
			ut.AssertEQ(t, v3, fst[v6].Parent)
		})
	}
}

func TestDFS_undirected(t *testing.T) {
	for _, classify := range []bool{true, false} {
		t.Run(fmt.Sprintf("classify: %v", classify), func(t *testing.T) {
			g, vars, err := ds.NewTextParser().Parse(ut.UUGSimple)

			ut.AssertEQ(t, true, err == nil)

			vR := vars["r"]
			vS := vars["s"]
			vT := vars["t"]
			vU := vars["u"]
			vV := vars["v"]
			vW := vars["w"]
			vX := vars["x"]
			vY := vars["y"]

			fst, tps, err := DFS(g, classify)

			ut.AssertEQ(t, true, err == nil)

			ut.AssertEQ(t, 0, len(tps.Forward))
			ut.AssertEQ(t, 0, len(tps.Cross))

			if classify {
				// undirected graph, so two for each edge
				ut.AssertEQ(t, 6, len(tps.Back))
			} else {
				ut.AssertEQ(t, 0, len(tps.Back))
			}

			ut.AssertEQ(t, 1, fst[vR].Discovery)
			ut.AssertEQ(t, 16, fst[vR].Finish)
			ut.AssertEQ(t, nil, fst[vR].Parent)

			ut.AssertEQ(t, 2, fst[vS].Discovery)
			ut.AssertEQ(t, 13, fst[vS].Finish)
			ut.AssertEQ(t, vR, fst[vS].Parent)

			ut.AssertEQ(t, 4, fst[vT].Discovery)
			ut.AssertEQ(t, 11, fst[vT].Finish)
			ut.AssertEQ(t, vW, fst[vT].Parent)

			ut.AssertEQ(t, 5, fst[vU].Discovery)
			ut.AssertEQ(t, 10, fst[vU].Finish)
			ut.AssertEQ(t, vT, fst[vU].Parent)

			ut.AssertEQ(t, 14, fst[vV].Discovery)
			ut.AssertEQ(t, 15, fst[vV].Finish)
			ut.AssertEQ(t, vR, fst[vV].Parent)

			ut.AssertEQ(t, 3, fst[vW].Discovery)
			ut.AssertEQ(t, 12, fst[vW].Finish)
			ut.AssertEQ(t, vS, fst[vW].Parent)

			ut.AssertEQ(t, 6, fst[vX].Discovery)
			ut.AssertEQ(t, 9, fst[vX].Finish)
			ut.AssertEQ(t, vU, fst[vX].Parent)

			ut.AssertEQ(t, 7, fst[vY].Discovery)
			ut.AssertEQ(t, 8, fst[vY].Finish)
			ut.AssertEQ(t, vX, fst[vY].Parent)
		})
	}
}
