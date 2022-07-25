package algo

import (
	"fmt"
	"testing"

	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

func mstBenchGen(n int) *ds.G {
	g := ds.NewGraph()

	for i := 0; i < n; i++ {
		v := ut.BenchItem(i)
		g.AddVertex(&v)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}

			g.AddEdge(
				g.V[i].Item,
				g.V[j].Item,
				float64(i+j),
			)
		}
	}

	return g
}

func BenchmarkMST(b *testing.B) {
	for _, size := range []int{16, 256, 1024} {
		b.Run(fmt.Sprintf("kruskal-%d", size), func(b *testing.B) {
			g := mstBenchGen(size)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				MSTKruskal(g)
			}
		})

		b.Run(fmt.Sprintf("prim-%d", size), func(b *testing.B) {
			g := mstBenchGen(size)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				MSTPrim(g)
			}
		})
	}
}
