package algo

import (
	"fmt"
	"testing"

	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

func ccBenchGen(n int) *ds.G {
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
				0,
			)
		}
	}

	return g
}

func BenchmarkCC(b *testing.B) {
	for _, size := range []int{16, 256, 1024} {
		b.Run(fmt.Sprintf("dfs-%d", size), func(b *testing.B) {
			g := ccBenchGen(size)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				CCDFS(g)
			}
		})

		b.Run(fmt.Sprintf("union-find-%d", size), func(b *testing.B) {
			g := ccBenchGen(size)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				CCUnionFind(g)
			}
		})
	}
}
