package viz

import (
	"github.com/vc-souza/gga/algo"
	"github.com/vc-souza/gga/ds"
)

/*
DFSViz formats and exports a graph after the execution of the DFS algorithm.
The output of the algorithm is traversed, and hooks are provided so that
custom formatting can be applied to the graph, its vertices and edges.
*/
type DFSViz struct {
	ThemedGraphViz

	Forest algo.DFForest
	Edges  *algo.EdgeTypes

	// OnTreeVertex is called for every vertex in the graph.
	OnTreeVertex func(int, algo.DFNode)

	// OnRootVertex is called when the root of a DF tree is found.
	OnRootVertex func(int, algo.DFNode)

	// OnTreeEdge is called when a tree edge is found.
	OnTreeEdge func(int, int)

	// OnForwardEdge is called when a forward edge is found.
	OnForwardEdge func(int, int)

	// OnBackEdge is called when a back edge is found.
	OnBackEdge func(int, int)

	// OnCrossEdge is called when a cross edge is found.
	OnCrossEdge func(int, int)
}

// NewDFSViz initializes a new DFSViz with NOOP hooks.
func NewDFSViz(g *ds.G, f algo.DFForest, e *algo.EdgeTypes, t Theme) *DFSViz {
	res := &DFSViz{}

	res.Forest = f
	res.Edges = e

	res.Graph = g
	res.Theme = t

	res.OnTreeVertex = func(int, algo.DFNode) {}
	res.OnRootVertex = func(int, algo.DFNode) {}

	res.OnTreeEdge = func(int, int) {}
	res.OnForwardEdge = func(int, int) {}
	res.OnBackEdge = func(int, int) {}
	res.OnCrossEdge = func(int, int) {}

	return res
}

// Traverse iterates over the results of a DFS execution, calling its hooks when appropriate.
func (vi *DFSViz) Traverse() error {
	for v, node := range vi.Forest {
		vi.OnTreeVertex(v, node)

		if node.Parent == -1 {
			vi.OnRootVertex(v, node)
			continue
		}

		v, e, ok := vi.Graph.GetEdge(
			vi.Graph.V[node.Parent].Item,
			vi.Graph.V[v].Item,
		)

		if !ok {
			return ds.ErrNoEdge
		}

		vi.OnTreeEdge(v, e)

		if vi.Graph.Directed() {
			continue
		}

		v, e, ok = vi.Graph.GetEdge(
			vi.Graph.V[v].Item,
			vi.Graph.V[node.Parent].Item,
		)

		if !ok {
			return ds.ErrNoRevEdge
		}

		vi.OnTreeEdge(v, e)
	}

	for _, e := range vi.Edges.Forward {
		vi.OnForwardEdge(e.Src, e.Index)
	}

	for _, e := range vi.Edges.Back {
		vi.OnBackEdge(e.Src, e.Index)
	}

	for _, e := range vi.Edges.Cross {
		vi.OnCrossEdge(e.Src, e.Index)
	}

	return nil
}
