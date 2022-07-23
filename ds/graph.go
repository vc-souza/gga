package ds

import (
	"fmt"
	"strings"
)

// TODO: docs
type GE struct {
	Formattable

	Src int
	Dst int
	Wt  float64
}

func (e *GE) String() string {
	return fmt.Sprintf(
		"%d -> %d <%.2f>",
		e.Src,
		e.Dst,
		e.Wt,
	)
}

// TODO: docs
type GV struct {
	Formattable

	Index int
	Item  Item
	E     []GE
}

func (v *GV) String() string {
	b := strings.Builder{}
	es := []string{}

	b.WriteString(fmt.Sprintf("Vertex '%s' @<%d> adj [", v.Item.Label(), v.Index))

	for i := range v.E {
		es = append(es, v.E[i].String())
	}

	b.WriteString(strings.Join(es, ", "))
	b.WriteString("]\n")

	return b.String()
}

// TODO: docs
type G struct {
	V []GV

	sat    map[Item]int
	dir    bool
	eCount int
	vCount int
}

// TODO: docs
func newG(dir bool) *G {
	g := &G{}

	g.V = make([]GV, 0)

	g.sat = map[Item]int{}
	g.dir = dir

	return g
}

// TODO: docs
func NewGraph() *G {
	return newG(false)
}

// TODO: docs
func NewDigraph() *G {
	return newG(true)
}

func (g *G) String() string {
	b := strings.Builder{}

	if g.Undirected() {
		b.WriteString("Undirected Graph\n")
	} else {
		b.WriteString("Directed Graph\n")
	}

	b.WriteString(fmt.Sprintf("%d map entries\n", len(g.sat)))

	for i := range g.V {
		b.WriteString(g.V[i].String())
	}

	return b.String()
}

// TODO: docs
func (g *G) Directed() bool {
	return g.dir
}

// TODO: docs
func (g *G) Undirected() bool {
	return !g.dir
}

// TODO: docs
func (g *G) VertexCount() int {
	return g.vCount
}

// TODO: docs
func (g *G) GetVertex(i Item) (*GV, int, bool) {
	idx, ok := g.sat[i]

	if !ok {
		return nil, -1, false
	}

	return &g.V[idx], idx, true
}

// TODO: docs
func (g *G) FromIndex(idx int) *GV {
	return &g.V[idx]
}

// TODO: docs
func (g *G) AddVertex(i Item) (*GV, error) {
	if idx, ok := g.sat[i]; ok {
		return &g.V[idx], ErrExists
	}

	g.V = append(g.V, GV{Item: i})

	idx := len(g.V) - 1

	g.V[idx].Index = idx
	g.sat[i] = idx

	g.vCount++

	return &g.V[idx], nil
}

// TODO: docs
func (g *G) RemoveVertex(i Item) error {
	iDel, ok := g.sat[i]

	if !ok {
		return ErrNoVtx
	}

	fixEdges := func() {
		for i := range g.V {
			if i == iDel {
				g.eCount -= len(g.V[i].E)
				continue
			}

			toDel := []int{}

			for j := range g.V[i].E {
				edge := &g.V[i].E[j]

				if edge.Dst == iDel {
					toDel = append(toDel, j)
					continue
				}

				// TODO: explain fix
				if edge.Src > iDel {
					edge.Src--
				}

				// TODO: explain fix
				if edge.Dst > iDel {
					edge.Dst--
				}
			}

			for _, eIdx := range toDel {
				Cut(&g.V[i].E, eIdx)
				g.eCount--
			}
		}
	}

	deleteVertex := func() {
		Cut(&g.V, iDel)
		delete(g.sat, i)
		g.vCount--
	}

	fixVertices := func() {
		for i := iDel; i < len(g.V); i++ {
			g.V[i].Index = i
			g.sat[g.V[i].Item] = i
		}
	}

	fixEdges()
	deleteVertex()
	fixVertices()

	return nil
}

// TODO: docs
func (g *G) EdgeCount() int {
	return g.eCount
}

// TODO: docs
func (g *G) GetEdge(src Item, dst Item) (*GE, int, bool) {
	vSrc, _, ok := g.GetVertex(src)

	if !ok {
		return nil, -1, false
	}

	iDst, ok := g.sat[dst]

	if !ok {
		return nil, -1, false
	}

	for i := range vSrc.E {
		if vSrc.E[i].Dst == iDst {
			return &vSrc.E[i], i, true
		}
	}

	return nil, -1, false
}

// TODO: docs
func (g *G) AddEdge(src Item, dst Item, wt float64) (*GE, error) {
	if g.Undirected() && src == dst {
		return nil, ErrInvLoop
	}

	vSrc, _ := g.AddVertex(src)
	vDst, _ := g.AddVertex(dst)

	for i := range vSrc.E {
		if vSrc.E[i].Dst == vDst.Index {
			return nil, ErrExists
		}
	}

	vSrc.E = append(vSrc.E, GE{Src: vSrc.Index, Dst: vDst.Index, Wt: wt})

	g.eCount++

	return &(vSrc.E[len(vSrc.E)-1]), nil
}

// TODO: docs
func (g *G) RemoveEdge(src Item, dst Item) error {
	_, idx, ok := g.GetEdge(src, dst)

	if !ok {
		return ErrNoEdge
	}

	Cut(&g.V[g.sat[src]].E, idx)

	g.eCount--

	return nil
}

// TODO: docs
func (g *G) Accept(v GraphVisitor) {
	v.VisitGraphStart(g)

	for i := range g.V {
		v.VisitVertex(&g.V[i])

		for j := range g.V[i].E {
			v.VisitEdge(&g.V[i].E[j])
		}
	}

	v.VisitGraphEnd(g)
}

// TODO: docs (right place? maybe algo?)
func (g *G) Transpose() (*G, error) {
	if g.Undirected() {
		return nil, ErrUndirected
	}

	res := NewDigraph()

	for i := range g.V {
		res.AddVertex(g.V[i].Item)
	}

	for i := range g.V {
		for j := range g.V[i].E {
			edge := &g.V[i].E[j]

			res.AddEdge(
				g.V[edge.Dst].Item,
				g.V[edge.Src].Item,
				edge.Wt,
			)
		}
	}

	return res, nil
}
