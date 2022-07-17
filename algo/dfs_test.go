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

			ut.Equal(t, true, err == nil)

			v1 := vars["1"]
			v2 := vars["2"]
			v3 := vars["3"]
			v4 := vars["4"]
			v5 := vars["5"]
			v6 := vars["6"]

			fst, tps, err := DFS(g, classify)

			ut.Equal(t, true, err == nil)

			if classify {
				ut.Equal(t, 1, len(tps.Forward))
				ut.Equal(t, 2, len(tps.Back))
				ut.Equal(t, 1, len(tps.Cross))
			} else {
				ut.Equal(t, 0, len(tps.Forward))
				ut.Equal(t, 0, len(tps.Back))
				ut.Equal(t, 0, len(tps.Cross))
			}

			ut.Equal(t, 1, fst[v1].Discovery)
			ut.Equal(t, 8, fst[v1].Finish)
			ut.Equal(t, nil, fst[v1].Parent)

			ut.Equal(t, 2, fst[v2].Discovery)
			ut.Equal(t, 7, fst[v2].Finish)
			ut.Equal(t, v1, fst[v2].Parent)

			ut.Equal(t, 9, fst[v3].Discovery)
			ut.Equal(t, 12, fst[v3].Finish)
			ut.Equal(t, nil, fst[v3].Parent)

			ut.Equal(t, 4, fst[v4].Discovery)
			ut.Equal(t, 5, fst[v4].Finish)
			ut.Equal(t, v5, fst[v4].Parent)

			ut.Equal(t, 3, fst[v5].Discovery)
			ut.Equal(t, 6, fst[v5].Finish)
			ut.Equal(t, v2, fst[v5].Parent)

			ut.Equal(t, 10, fst[v6].Discovery)
			ut.Equal(t, 11, fst[v6].Finish)
			ut.Equal(t, v3, fst[v6].Parent)
		})
	}
}

func TestDFS_undirected(t *testing.T) {
	for _, classify := range []bool{true, false} {
		t.Run(fmt.Sprintf("classify: %v", classify), func(t *testing.T) {
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

			fst, tps, err := DFS(g, classify)

			ut.Equal(t, true, err == nil)

			ut.Equal(t, 0, len(tps.Forward))
			ut.Equal(t, 0, len(tps.Cross))

			if classify {
				// undirected graph, so two for each edge
				ut.Equal(t, 6, len(tps.Back))
			} else {
				ut.Equal(t, 0, len(tps.Back))
			}

			ut.Equal(t, 1, fst[vR].Discovery)
			ut.Equal(t, 16, fst[vR].Finish)
			ut.Equal(t, nil, fst[vR].Parent)

			ut.Equal(t, 2, fst[vS].Discovery)
			ut.Equal(t, 13, fst[vS].Finish)
			ut.Equal(t, vR, fst[vS].Parent)

			ut.Equal(t, 4, fst[vT].Discovery)
			ut.Equal(t, 11, fst[vT].Finish)
			ut.Equal(t, vW, fst[vT].Parent)

			ut.Equal(t, 5, fst[vU].Discovery)
			ut.Equal(t, 10, fst[vU].Finish)
			ut.Equal(t, vT, fst[vU].Parent)

			ut.Equal(t, 14, fst[vV].Discovery)
			ut.Equal(t, 15, fst[vV].Finish)
			ut.Equal(t, vR, fst[vV].Parent)

			ut.Equal(t, 3, fst[vW].Discovery)
			ut.Equal(t, 12, fst[vW].Finish)
			ut.Equal(t, vS, fst[vW].Parent)

			ut.Equal(t, 6, fst[vX].Discovery)
			ut.Equal(t, 9, fst[vX].Finish)
			ut.Equal(t, vU, fst[vX].Parent)

			ut.Equal(t, 7, fst[vY].Discovery)
			ut.Equal(t, 8, fst[vY].Finish)
			ut.Equal(t, vX, fst[vY].Parent)
		})
	}
}
