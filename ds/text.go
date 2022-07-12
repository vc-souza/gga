package ds

import (
	"fmt"
	"strconv"
	"strings"
)

/*
Text trivially implements the ds.Item interface, and it is meant for both
basic graphs where only strings are manipulated as vertices, and for unit
testing in all packages of this module.
*/
type Text string

func (t Text) Label() string {
	return string(t)
}

const (
	UndirectedGraphKey = "graph"
	DirectedGraphKey   = "digraph"
)

/*
ParseGraph produces a new graph from a string containing text in the the following grammar:
	Graph = GraphType ["\n" AdjEntries]
	GraphType = "graph" | "digraph"
	AdjEntries = AdjEntry {"\n" AdjEntry}
	AdjEntry = Vertex "#" [EdgeList]
	Vertex = all characters but "#", "\n", ":", ","
	EdgeList = Edge {"," Edge}
	Edge = Vertex [":" Weight]
	Weight = float

Sample (Undirected):
	graph
	a#b:4,h:8
	b#a:4,c:8,h:11
	c#b:8,d:7,i:2,f:4
	d#c:7,e:9,f:14
	e#d:9,f:10
	f#c:4,d:14,e:10,g:2
	g#f:2,h:1,i:6
	h#a:8,b:11,g:1,i:7
	i#c:2,g:6,h:7

Sample (Directed)
	digraph
	1#2,4
	2#5
	3#5,6
	4#2
	5#4
	6#6
*/
func ParseGraph(s string) (*Graph[Text], error) {
	addrs := make(map[string]*Text)
	var g *Graph[Text]

	vPtr := func(raw string) *Text {
		var res *Text

		if v, ok := addrs[raw]; ok {
			res = v
		} else {
			v := Text(raw)
			addrs[raw] = &v
			res = &v

			g.UnsafeAddVertex(res)
		}

		return res
	}

	bail := func(l string) error {
		return fmt.Errorf("%w: %s", ErrInvalidSer, l)
	}

	for _, l := range strings.Split(s, "\n") {
		l = strings.Trim(l, "\n")

		if len(l) == 0 {
			continue
		}

		if g == nil {
			switch l {

			case UndirectedGraphKey:
				g = NewUndirectedGraph[Text]()

			case DirectedGraphKey:
				g = NewDirectedGraph[Text]()

			default:
				return nil, bail(l)
			}

			continue
		}

		adj := strings.Split(l, "#")

		if len(adj) != 2 {
			return nil, bail(l)
		}

		src := vPtr(adj[0])

		for _, e := range strings.Split(adj[1], ",") {
			var wt float64

			edge := strings.Split(e, ":")

			if len(edge) < 1 || len(edge) > 2 {
				return nil, bail(l)
			}

			dst := vPtr(edge[0])

			if len(edge) == 2 {
				pWt, err := strconv.ParseFloat(edge[1], 64)

				if err != nil {
					return nil, err
				}

				wt = pWt
			}

			g.UnsafeAddWeightedEdge(src, dst, wt)
		}
	}

	return g, nil
}
