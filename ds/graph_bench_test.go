package ds

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

type idx int

func (i idx) Label() string {
	return strconv.Itoa(int(i))
}

var gens = []struct {
	f    func() *Graph[idx]
	name string
}{
	{
		NewUndirectedGraph[idx],
		"graph",
	},
	{
		NewDirectedGraph[idx],
		"digraph",
	},
}

var sizes = []int{
	int(math.Pow(32, 1)),
	int(math.Pow(32, 2)),
	int(math.Pow(32, 3)),
}

func tagGraphBench(gType, mode string, size int) string {
	return fmt.Sprintf("%s_%s_%d", gType, mode, size)
}

func BenchmarkGraphAddingVertex(b *testing.B) {
	// probably not significant, but worth a try:
	// the higher the load of the hash table / map,
	// the higher the chance of a collision:
	// the actual time (not amortized) of the
	// map lookup might be higher than O(1),
	// depending on the type of collision
	// resolution used by the map.
	addVerts := func(g *Graph[idx], size int) {
		for i := 0; i < size; i++ {
			item := idx(i)
			g.UnsafeAddVertex(&item)
		}
	}

	for _, gen := range gens {
		for _, size := range sizes {
			b.Run(tagGraphBench(gen.name, "safe", size), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					b.StopTimer()

					g := gen.f()
					v := idx(size)

					addVerts(g, size)

					b.StartTimer()

					// has an extra map lookup
					g.AddVertex(&v)
				}
			})

			b.Run(tagGraphBench(gen.name, "unsafe", size), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					b.StopTimer()

					g := gen.f()
					v := idx(size)

					addVerts(g, size)

					b.StartTimer()

					g.UnsafeAddVertex(&v)
				}
			})
		}
	}
}

func BenchmarkGraphAddingEdges(b *testing.B) {
	// to build the worst case, we need a line of vertices
	addVerts := func(g *Graph[idx], size int) []idx {
		verts := make([]idx, size)

		for i := 0; i < size; i++ {
			verts[i] = idx(i)
			g.UnsafeAddVertex(&verts[i])
		}

		return verts
	}

	// the worst-case scenario: one vertex v has len(Adj[v]) = O(|E|)
	addEdges := func(g *Graph[idx], verts []idx) {
		for i := 1; i < len(verts)-1; i++ {
			g.UnsafeAddWeightedEdge(&verts[0], &verts[i], 0)
		}
	}

	for _, gen := range gens {
		for _, size := range sizes {
			b.Run(tagGraphBench(gen.name, "safe", size), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					b.StopTimer()

					g := gen.f()

					verts := addVerts(g, size)
					addEdges(g, verts)

					b.StartTimer()

					// O(E) checks
					g.AddWeightedEdge(&verts[0], &verts[len(verts)-1], 0)
				}
			})

			b.Run(tagGraphBench(gen.name, "unsafe", size), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					b.StopTimer()

					g := gen.f()

					verts := addVerts(g, size)
					addEdges(g, verts)

					b.StartTimer()

					g.UnsafeAddWeightedEdge(&verts[0], &verts[len(verts)-1], 0)
				}
			})
		}
	}
}
