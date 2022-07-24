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

	Index int
}

// TODO: docs
type GEIdx struct {
	V int
	E int
}

func (e *GE) String() string {
	return fmt.Sprintf(
		"@%d %d -> %d <%.2f>",
		e.Index,
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

	b.WriteString(fmt.Sprintf("Vertex '%s' @%d adj [", v.Item.Label(), v.Index))

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
		b.WriteString("Graph")
	} else {
		b.WriteString("Digraph")
	}

	b.WriteString(fmt.Sprintf(" |V| = %d", g.VertexCount()))
	b.WriteString(fmt.Sprintf(" |E| = %d\n", g.EdgeCount()))

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
func (g *G) GetVertex(i Item) (int, bool) {
	idx, ok := g.sat[i]
	return idx, ok
}

// TODO: docs
func (g *G) AddVertex(i Item) (int, error) {
	if _, ok := g.sat[i]; ok {
		return 0, ErrExists
	}

	g.V = append(g.V, GV{Item: i})

	idx := len(g.V) - 1

	g.V[idx].Index = idx
	g.sat[i] = idx

	g.vCount++

	return idx, nil
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
			dec := 0

			for j := range g.V[i].E {
				edge := &g.V[i].E[j]

				if edge.Dst == iDel {
					toDel = append(toDel, j)
					dec++
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

				// TODO: explain fix
				edge.Index -= dec
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
func (g *G) GetEdge(src Item, dst Item) (int, int, bool) {
	iSrc, ok := g.GetVertex(src)

	if !ok {
		return 0, 0, false
	}

	iDst, ok := g.sat[dst]

	if !ok {
		return 0, 0, false
	}

	for j := range g.V[iSrc].E {
		if g.V[iSrc].E[j].Dst == iDst {
			return iSrc, j, true
		}
	}

	return 0, 0, false
}

// TODO: docs
func (g *G) AddEdge(src Item, dst Item, wt float64) (int, error) {
	if g.Undirected() && src == dst {
		return 0, ErrInvLoop
	}

	iSrc, ok := g.GetVertex(src)

	if !ok {
		return 0, ErrNoVtx
	}

	iDst, ok := g.GetVertex(dst)

	if !ok {
		return 0, ErrNoVtx
	}

	for j := range g.V[iSrc].E {
		if g.V[iSrc].E[j].Dst == iDst {
			return 0, ErrExists
		}
	}

	g.V[iSrc].E = append(
		g.V[iSrc].E,
		GE{
			Index: len(g.V[iSrc].E),
			Src:   iSrc,
			Dst:   iDst,
			Wt:    wt,
		},
	)

	g.eCount++

	return len(g.V[iSrc].E) - 1, nil
}

// TODO: docs
func (g *G) RemoveEdge(src Item, dst Item) error {
	vIdx, idx, ok := g.GetEdge(src, dst)

	if !ok {
		return ErrNoEdge
	}

	Cut(&g.V[vIdx].E, idx)

	for i := idx; i < len(g.V[vIdx].E); i++ {
		g.V[vIdx].E[i].Index--
	}

	g.eCount--

	return nil
}

// TODO: docs
func (g G) Accept(v GraphVisitor) {
	v.VisitGraphStart(g)

	for i := range g.V {
		v.VisitVertex(g, i)

		for j := range g.V[i].E {
			v.VisitEdge(g, i, j)
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
		for _, edge := range g.V[i].E {
			res.AddEdge(
				g.V[edge.Dst].Item,
				g.V[edge.Src].Item,
				edge.Wt,
			)
		}
	}

	return res, nil
}
