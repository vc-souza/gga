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
TextParser produces a new graph from text in the following grammar:

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
type TextParser struct {
	addrs map[string]*Text
	graph *Graph[Text]
}

func (p *TextParser) wrapErr(err error, s string) error {
	return fmt.Errorf("%w: %s", err, s)
}

func (p *TextParser) parseGraphType(raw string) error {
	switch raw {

	case UndirectedGraphKey:
		p.graph = NewUndirectedGraph[Text]()

	case DirectedGraphKey:
		p.graph = NewDirectedGraph[Text]()

	default:
		return p.wrapErr(ErrInvalidSer, "graph type: bad name")
	}

	return nil
}

func (p *TextParser) parseVertex(raw string) (*Text, error) {
	if len(raw) == 0 {
		return nil, p.wrapErr(ErrInvalidSer, "vertex: empty name")
	}

	if strings.ContainsAny(raw, InvalidVertexRunes) {
		return nil, p.wrapErr(ErrInvalidSer, "vertex: bad name")
	}

	var res *Text

	if v, ok := p.addrs[raw]; ok {
		res = v
	} else {
		v := Text(raw)
		p.addrs[raw] = &v
		res = &v

		p.graph.UnsafeAddVertex(res)
	}

	return res, nil
}

func (p *TextParser) parseEdge(src *Text, raw string) error {
	if len(raw) == 0 {
		return p.wrapErr(ErrInvalidSer, "edge: empty")
	}

	var wt float64

	edge := strings.Split(raw, ":")

	if len(edge) < 1 || len(edge) > 2 {
		return p.wrapErr(ErrInvalidSer, "edge: wrong item count")
	}

	dst, err := p.parseVertex(edge[0])

	if err != nil {
		return err
	}

	if len(edge) == 2 {
		pWt, err := strconv.ParseFloat(edge[1], 64)

		if err != nil {
			return p.wrapErr(err, "weight: bad value")
		}

		wt = pWt
	}

	p.graph.UnsafeAddWeightedEdge(src, dst, wt)

	return nil
}

func (p *TextParser) parseEdgeList(src *Text, raw string) error {
	if len(raw) == 0 {
		return nil
	}

	for _, e := range strings.Split(raw, ",") {
		err := p.parseEdge(src, e)

		if err != nil {
			return err
		}
	}

	return nil
}

func (p *TextParser) parseAdjEntry(raw string) error {
	adj := strings.Split(raw, "#")

	if len(adj) != 2 {
		return p.wrapErr(ErrInvalidSer, "adjacency list: wrong item count")
	}

	src, err := p.parseVertex(adj[0])

	if err != nil {
		return err
	}

	err = p.parseEdgeList(src, adj[1])

	if err != nil {
		return err
	}

	return nil
}

// Parse parses the input string, generating a new graph.
func (p *TextParser) Parse(s string) (*Graph[Text], error) {
	p.addrs = make(map[string]*Text)
	p.graph = nil

	for _, l := range strings.Split(s, "\n") {
		l = strings.Trim(l, "\n\t")

		if len(l) == 0 {
			continue
		}

		if p.graph == nil {
			err := p.parseGraphType(l)

			if err != nil {
				return nil, p.wrapErr(err, l)
			}

			continue
		}

		err := p.parseAdjEntry(l)

		if err != nil {
			return nil, p.wrapErr(err, l)
		}
	}

	return p.graph, nil
}
