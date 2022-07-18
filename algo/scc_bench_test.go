package algo

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/vc-souza/gga/ds"
)

type id int

func (i id) Label() string { return strconv.Itoa(int(i)) }

func BenchmarkSCC(b *testing.B) {
	gen := func(n int) *ds.Graph[id] {
		g := ds.NewDirectedGraph[id]()

		for i := 0; i < n; i++ {
			v := id(i)
			g.UnsafeAddVertex(&v)
		}

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				g.UnsafeAddWeightedEdge(
					g.Verts[i].Val,
					g.Verts[j].Val,
					0,
				)
			}
		}

		return g
	}

	inputs := []struct {
		g *ds.Graph[id]
		n int
	}{
		{
			gen(16),
			16,
		},
		{
			gen(256),
			256,
		},
		{
			gen(1024),
			1024,
		},
	}

	for _, in := range inputs {
		b.Run(fmt.Sprintf("kosaraju-%d", in.n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SCCKosaraju(in.g)
			}
		})

		b.Run(fmt.Sprintf("tarjan-%d", in.n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SCCTarjan(in.g)
			}
		})
	}
}
