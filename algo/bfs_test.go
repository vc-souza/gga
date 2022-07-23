package algo

import (
	"math"
	"testing"

	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestBFS_directed(t *testing.T) {
	g, _, err := ds.Parse(ut.UDGSimple)

	ut.Nil(t, err)

	tree, err := BFS(g, 2)

	ut.Nil(t, err)

	for i := range tree {
		switch i {
		case 0:
			ut.True(t, math.IsInf(tree[i].Distance, 1))
			ut.Equal(t, -1, tree[i].Parent)
		case 1:
			ut.Equal(t, 3, tree[i].Distance)
			ut.Equal(t, 3, tree[i].Parent)
		case 2:
			ut.Equal(t, 0, tree[i].Distance)
			ut.Equal(t, -1, tree[i].Parent)
		case 3:
			ut.Equal(t, 2, tree[i].Distance)
			ut.Equal(t, 4, tree[i].Parent)
		case 4:
		case 5:
			ut.Equal(t, 1, tree[i].Distance)
			ut.Equal(t, 2, tree[i].Parent)
		}
	}
}

func TestBFS_undirected(t *testing.T) {
	g, _, err := ds.Parse(ut.UUGSimple)

	ut.Nil(t, err)

	tree, err := BFS(g, 3)

	ut.Nil(t, err)

	for i := range tree {
		switch i {
		case 0:
			ut.Equal(t, 4, tree[i].Distance)
			ut.Equal(t, 1, tree[i].Parent)
		case 1:
			ut.Equal(t, 3, tree[i].Distance)
			ut.Equal(t, 5, tree[i].Parent)
		case 2:
			ut.Equal(t, 1, tree[i].Distance)
			ut.Equal(t, 3, tree[i].Parent)
		case 3:
			ut.Equal(t, 0, tree[i].Distance)
			ut.Equal(t, -1, tree[i].Parent)
		case 4:
			ut.Equal(t, 5, tree[i].Distance)
			ut.Equal(t, 0, tree[i].Parent)
		case 5:
			ut.Equal(t, 2, tree[i].Distance)
			ut.Equal(t, 2, tree[i].Parent)
		case 6:
			ut.Equal(t, 1, tree[i].Distance)
			ut.Equal(t, 3, tree[i].Parent)
		case 7:
			ut.Equal(t, 1, tree[i].Distance)
			ut.Equal(t, 3, tree[i].Parent)
		}
	}
}
