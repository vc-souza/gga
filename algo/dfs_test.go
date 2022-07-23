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
			g, idx, err := ds.Parse(ut.UDGSimple)

			ut.Nil(t, err)

			fst, tps, err := DFS(g, classify)

			ut.Nil(t, err)

			if classify {
				ut.Equal(t, 1, len(tps.Forward))
				ut.Equal(t, 2, len(tps.Back))
				ut.Equal(t, 1, len(tps.Cross))
			} else {
				ut.Equal(t, 0, len(tps.Forward))
				ut.Equal(t, 0, len(tps.Back))
				ut.Equal(t, 0, len(tps.Cross))
			}

			for i := range fst {
				switch i {
				case idx("1"):
					ut.Equal(t, 1, fst[i].Discovery)
					ut.Equal(t, 8, fst[i].Finish)
					ut.Equal(t, -1, fst[i].Parent)
				case idx("2"):
					ut.Equal(t, 2, fst[i].Discovery)
					ut.Equal(t, 7, fst[i].Finish)
					ut.Equal(t, idx("1"), fst[i].Parent)
				case idx("3"):
					ut.Equal(t, 9, fst[i].Discovery)
					ut.Equal(t, 12, fst[i].Finish)
					ut.Equal(t, -1, fst[i].Parent)
				case idx("4"):
					ut.Equal(t, 4, fst[i].Discovery)
					ut.Equal(t, 5, fst[i].Finish)
					ut.Equal(t, idx("5"), fst[i].Parent)
				case idx("5"):
					ut.Equal(t, 3, fst[i].Discovery)
					ut.Equal(t, 6, fst[i].Finish)
					ut.Equal(t, idx("2"), fst[i].Parent)
				case idx("6"):
					ut.Equal(t, 10, fst[i].Discovery)
					ut.Equal(t, 11, fst[i].Finish)
					ut.Equal(t, idx("3"), fst[i].Parent)
				}
			}
		})
	}
}

func TestDFS_undirected(t *testing.T) {
	for _, classify := range []bool{true, false} {
		t.Run(fmt.Sprintf("classify: %v", classify), func(t *testing.T) {
			g, idx, err := ds.Parse(ut.UUGSimple)

			ut.Nil(t, err)

			fst, tps, err := DFS(g, classify)

			ut.Nil(t, err)

			ut.Equal(t, 0, len(tps.Forward))
			ut.Equal(t, 0, len(tps.Cross))

			if classify {
				// undirected graph, so two for each edge
				ut.Equal(t, 6, len(tps.Back))
			} else {
				ut.Equal(t, 0, len(tps.Back))
			}

			for i := range fst {
				switch i {
				case idx("r"):
					ut.Equal(t, 1, fst[i].Discovery)
					ut.Equal(t, 16, fst[i].Finish)
					ut.Equal(t, -1, fst[i].Parent)
				case idx("s"):
					ut.Equal(t, 2, fst[i].Discovery)
					ut.Equal(t, 13, fst[i].Finish)
					ut.Equal(t, idx("r"), fst[i].Parent)
				case idx("t"):
					ut.Equal(t, 4, fst[i].Discovery)
					ut.Equal(t, 11, fst[i].Finish)
					ut.Equal(t, idx("w"), fst[i].Parent)
				case idx("u"):
					ut.Equal(t, 5, fst[i].Discovery)
					ut.Equal(t, 10, fst[i].Finish)
					ut.Equal(t, idx("t"), fst[i].Parent)
				case idx("v"):
					ut.Equal(t, 14, fst[i].Discovery)
					ut.Equal(t, 15, fst[i].Finish)
					ut.Equal(t, idx("r"), fst[i].Parent)
				case idx("w"):
					ut.Equal(t, 3, fst[i].Discovery)
					ut.Equal(t, 12, fst[i].Finish)
					ut.Equal(t, idx("s"), fst[i].Parent)
				case idx("x"):
					ut.Equal(t, 6, fst[i].Discovery)
					ut.Equal(t, 9, fst[i].Finish)
					ut.Equal(t, idx("u"), fst[i].Parent)
				case idx("y"):
					ut.Equal(t, 7, fst[i].Discovery)
					ut.Equal(t, 8, fst[i].Finish)
					ut.Equal(t, idx("x"), fst[i].Parent)
				}
			}
		})
	}
}
