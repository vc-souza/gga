package algo

import (
	"errors"
	"testing"

	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestTSort_directed(t *testing.T) {
	// for _, classify := range []bool{true, false} {
	// 	t.Run(fmt.Sprintf("classify: %v", classify), func(t *testing.T) {
	// 		g, vars, err := ds.NewTextParser().Parse(ut.UDGSimple)

	// 		ut.AssertEqual(t, true, err == nil)

	// 		v1 := vars["1"]
	// 		v2 := vars["2"]
	// 		v3 := vars["3"]
	// 		v4 := vars["4"]
	// 		v5 := vars["5"]
	// 		v6 := vars["6"]

	// 		fst, tps, err := DFS(g, classify)

	// 		ut.AssertEqual(t, true, err == nil)

	// 		if classify {
	// 			ut.AssertEqual(t, 1, len(tps.Forward))
	// 			ut.AssertEqual(t, 2, len(tps.Back))
	// 			ut.AssertEqual(t, 1, len(tps.Cross))
	// 		} else {
	// 			ut.AssertEqual(t, 0, len(tps.Forward))
	// 			ut.AssertEqual(t, 0, len(tps.Back))
	// 			ut.AssertEqual(t, 0, len(tps.Cross))
	// 		}

	// 		ut.AssertEqual(t, 1, fst[v1].Discovery)
	// 		ut.AssertEqual(t, 8, fst[v1].Finish)
	// 		ut.AssertEqual(t, nil, fst[v1].Parent)

	// 		ut.AssertEqual(t, 2, fst[v2].Discovery)
	// 		ut.AssertEqual(t, 7, fst[v2].Finish)
	// 		ut.AssertEqual(t, v1, fst[v2].Parent)

	// 		ut.AssertEqual(t, 9, fst[v3].Discovery)
	// 		ut.AssertEqual(t, 12, fst[v3].Finish)
	// 		ut.AssertEqual(t, nil, fst[v3].Parent)

	// 		ut.AssertEqual(t, 4, fst[v4].Discovery)
	// 		ut.AssertEqual(t, 5, fst[v4].Finish)
	// 		ut.AssertEqual(t, v5, fst[v4].Parent)

	// 		ut.AssertEqual(t, 3, fst[v5].Discovery)
	// 		ut.AssertEqual(t, 6, fst[v5].Finish)
	// 		ut.AssertEqual(t, v2, fst[v5].Parent)

	// 		ut.AssertEqual(t, 10, fst[v6].Discovery)
	// 		ut.AssertEqual(t, 11, fst[v6].Finish)
	// 		ut.AssertEqual(t, v3, fst[v6].Parent)
	// 	})
	// }
}

func TestTSort_undirected(t *testing.T) {
	g, _, err := ds.NewTextParser().Parse(ut.UUGSimple)

	ut.AssertEqual(t, true, err == nil)

	_, err = TSort(g)

	ut.AssertEqual(t, true, err != nil)
	ut.AssertEqual(t, true, errors.Is(err, ds.ErrUndefOp))
}
