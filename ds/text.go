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

const (
	InvalidVertexRunes = "#\n:,"
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

	wrapError := func(err error, s string) error {
		return fmt.Errorf("%w: %s", err, s)
	}

	parseType := func(raw string) error {
		switch raw {

		case UndirectedGraphKey:
			g = NewUndirectedGraph[Text]()

		case DirectedGraphKey:
			g = NewDirectedGraph[Text]()

		default:
			return ErrInvalidSer
		}

		return nil
	}

	parseVertex := func(raw string) (*Text, error) {
		if len(raw) == 0 {
			return nil, ErrInvalidSer
		}

		if strings.ContainsAny(raw, InvalidVertexRunes) {
			return nil, ErrInvalidSer
		}

		var res *Text

		if v, ok := addrs[raw]; ok {
			res = v
		} else {
			v := Text(raw)
			addrs[raw] = &v
			res = &v

			g.UnsafeAddVertex(res)
		}

		return res, nil
	}

	parseEdge := func(src *Text, raw string) error {
		if len(raw) == 0 {
			return ErrInvalidSer
		}

		var wt float64

		edge := strings.Split(raw, ":")

		if len(edge) < 1 || len(edge) > 2 {
			return ErrInvalidSer
		}

		dst, err := parseVertex(edge[0])

		if err != nil {
			return err
		}

		if len(edge) == 2 {
			pWt, err := strconv.ParseFloat(edge[1], 64)

			if err != nil {
				return err
			}

			wt = pWt
		}

		g.UnsafeAddWeightedEdge(src, dst, wt)

		return nil
	}

	parseEdgeList := func(src *Text, raw string) error {
		for _, e := range strings.Split(raw, ",") {
			err := parseEdge(src, e)

			if err != nil {
				return err
			}
		}

		return nil
	}

	parseAdjEntry := func(raw string) error {
		adj := strings.Split(raw, "#")

		if len(adj) != 2 {
			return ErrInvalidSer
		}

		src, err := parseVertex(adj[0])

		if err != nil {
			return err
		}

		err = parseEdgeList(src, adj[1])

		if err != nil {
			return err
		}

		return nil
	}

	for _, l := range strings.Split(s, "\n") {
		l = strings.Trim(l, "\n\t")

		if len(l) == 0 {
			continue
		}

		if g == nil {
			err := parseType(l)

			if err != nil {
				return nil, wrapError(err, l)
			}

			continue
		}

		err := parseAdjEntry(l)

		if err != nil {
			return nil, wrapError(err, l)
		}
	}

	return g, nil
}
