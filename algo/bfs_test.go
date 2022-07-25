package algo

import (
	"math"
	"testing"

	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestBFS_directed(t *testing.T) {
	g, idx, err := ds.Parse(ut.UDGSimple)

	ut.Nil(t, err)

	tree, err := BFS(g, 2)

	ut.Nil(t, err)

	for i := range tree {
		switch i {
		case idx("1"):
			ut.True(t, math.IsInf(tree[i].Distance, 1))
			ut.Equal(t, -1, tree[i].Parent)
		case idx("2"):
			ut.Equal(t, 3, tree[i].Distance)
			ut.Equal(t, idx("4"), tree[i].Parent)
		case idx("3"):
			ut.Equal(t, 0, tree[i].Distance)
			ut.Equal(t, -1, tree[i].Parent)
		case idx("4"):
			ut.Equal(t, 2, tree[i].Distance)
			ut.Equal(t, idx("5"), tree[i].Parent)
		case idx("5"):
		case idx("6"):
			ut.Equal(t, 1, tree[i].Distance)
			ut.Equal(t, idx("3"), tree[i].Parent)
		}
	}
}

func TestBFS_undirected(t *testing.T) {
	g, idx, err := ds.Parse(ut.UUGSimple)

	ut.Nil(t, err)

	tree, err := BFS(g, 3)

	ut.Nil(t, err)

	for i := range tree {
		switch i {
		case idx("r"):
			ut.Equal(t, 4, tree[i].Distance)
			ut.Equal(t, idx("s"), tree[i].Parent)
		case idx("s"):
			ut.Equal(t, 3, tree[i].Distance)
			ut.Equal(t, idx("w"), tree[i].Parent)
		case idx("t"):
			ut.Equal(t, 1, tree[i].Distance)
			ut.Equal(t, idx("u"), tree[i].Parent)
		case idx("u"):
			ut.Equal(t, 0, tree[i].Distance)
			ut.Equal(t, -1, tree[i].Parent)
		case idx("v"):
			ut.Equal(t, 5, tree[i].Distance)
			ut.Equal(t, idx("r"), tree[i].Parent)
		case idx("w"):
			ut.Equal(t, 2, tree[i].Distance)
			ut.Equal(t, idx("t"), tree[i].Parent)
		case idx("x"):
		case idx("y"):
			ut.Equal(t, 1, tree[i].Distance)
			ut.Equal(t, idx("u"), tree[i].Parent)
		}
	}
}
