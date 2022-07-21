package algo

import (
	"fmt"
	"testing"

	"github.com/vc-souza/gga/ds"
	ut "github.com/vc-souza/gga/internal/testutils"
)

func sccBenchGen(n int) *ds.G[ut.BenchItem] {
	g := ds.NewDirectedGraph[ut.BenchItem]()

	for i := 0; i < n; i++ {
		v := ut.BenchItem(i)
		g.UnsafeAddVertex(&v)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			g.UnsafeAddWeightedEdge(
				g.V[i].Ptr,
				g.V[j].Ptr,
				0,
			)
		}
	}

	return g
}

func BenchmarkSCC(b *testing.B) {
	for _, size := range []int{16, 256, 1024} {
		b.Run(fmt.Sprintf("kosaraju-%d", size), func(b *testing.B) {
			g := sccBenchGen(size)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				SCCKosaraju(g)
			}
		})

		b.Run(fmt.Sprintf("tarjan-%d", size), func(b *testing.B) {
			g := sccBenchGen(size)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				SCCTarjan(g)
			}
		})
	}
}
