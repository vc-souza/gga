package ds

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	UndirectedGraphKey = "graph"
	DirectedGraphKey   = "digraph"
)

const (
	InvalidVertexRunes = "#\n:,"
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

type ErrInvalidSer struct {
	Reason error
}

func (e ErrInvalidSer) Error() string {
	return fmt.Sprintf("invalid serialization: %s", e.Reason.Error())
}

func (e ErrInvalidSer) Unwrap() error {
	return e.Reason
}

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
	vars    map[string]*Text
	graph   *Graph[Text]
	pending map[*Text]string
}

// NewTextParser creates and returns a new TextParser.
func NewTextParser() *TextParser {
	return &TextParser{}
}

func (p *TextParser) parseGraphType(raw string) error {
	switch raw {

	case UndirectedGraphKey:
		p.graph = NewUndirectedGraph[Text]()

	case DirectedGraphKey:
		p.graph = NewDirectedGraph[Text]()

	default:
		return errors.New("graph type: bad name")
	}

	return nil
}

func (p *TextParser) parseVertex(raw string) (*Text, error) {
	if len(raw) == 0 {
		return nil, errors.New("vertex: empty name")
	}

	if strings.ContainsAny(raw, InvalidVertexRunes) {
		return nil, errors.New("vertex: bad name")
	}

	var res *Text

	if v, ok := p.vars[raw]; ok {
		res = v
	} else {
		v := Text(raw)
		p.vars[raw] = &v
		res = &v

		p.graph.UnsafeAddVertex(res)
	}

	return res, nil
}

func (p *TextParser) parseEdge(src *Text, raw string) error {
	if len(raw) == 0 {
		return errors.New("edge: empty")
	}

	var wt float64

	edge := strings.Split(raw, ":")

	if len(edge) < 1 || len(edge) > 2 {
		return errors.New("edge: wrong item count")
	}

	dst, ok := p.vars[edge[0]]

	if !ok {
		return errors.New("edge: unknown destination")
	}

	if len(edge) == 2 {
		pWt, err := strconv.ParseFloat(edge[1], 64)

		if err != nil {
			return fmt.Errorf("weight: bad value %w", err)
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
		if err := p.parseEdge(src, e); err != nil {
			return err
		}
	}

	return nil
}

func (p *TextParser) parseAdjEntry(raw string) error {
	adj := strings.Split(raw, "#")

	if len(adj) != 2 {
		return errors.New("adjacency list: wrong item count")
	}

	src, err := p.parseVertex(adj[0])

	if err != nil {
		return err
	}

	// postponing the processing of adjacency lists
	// so that vertices can be added in input order
	p.pending[src] = adj[1]

	return nil
}

// Parse parses the input string, generating a new graph.
func (p *TextParser) Parse(s string) (*Graph[Text], map[string]*Text, error) {
	p.vars = make(map[string]*Text)
	p.pending = make(map[*Text]string)
	p.graph = nil

	for _, l := range strings.Split(s, "\n") {
		l = strings.Trim(l, "\n\t")

		if len(l) == 0 {
			continue
		}

		if p.graph == nil {
			if err := p.parseGraphType(l); err != nil {
				return nil, nil, ErrInvalidSer{err}
			}

			continue
		}

		if err := p.parseAdjEntry(l); err != nil {
			return nil, nil, ErrInvalidSer{err}
		}
	}

	for src, raw := range p.pending {
		if err := p.parseEdgeList(src, raw); err != nil {
			return nil, nil, ErrInvalidSer{err}
		}
	}

	return p.graph, p.vars, nil
}
