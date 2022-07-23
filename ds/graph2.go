package ds

import (
	"fmt"
	"strings"
)

// TODO: docs
type GE2 struct {
	Formattable

	Src int
	Dst int
	Wt  float64
}

func (e *GE2) String() string {
	return fmt.Sprintf(
		"%d -> %d <%.2f>",
		e.Src,
		e.Dst,
		e.Wt,
	)
}

// TODO: docs
type GV2 struct {
	Formattable

	Index int
	Item  Item
	E     []GE2
}

func (v *GV2) String() string {
	b := strings.Builder{}
	es := []string{}

	b.WriteString(fmt.Sprintf("Vertex '%s' i:<%d> adj:<", v.Item.Label(), v.Index))

	for i := range v.E {
		es = append(es, v.E[i].String())
	}

	b.WriteString(strings.Join(es, ", "))
	b.WriteString(">\n")

	return b.String()
}

// TODO: docs
type G2 struct {
	V []GV2

	sat    map[Item]int
	dir    bool
	eCount int
	vCount int
}

// TODO: docs
func NewG2(dir bool) *G2 {
	g := &G2{}

	g.V = make([]GV2, 0)

	g.sat = map[Item]int{}
	g.dir = dir

	return g
}

func (g *G2) String() string {
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
func (g *G2) Directed() bool {
	return g.dir
}

// TODO: docs
func (g *G2) Undirected() bool {
	return !g.dir
}

// TODO: docs
func (g *G2) VertexCount() int {
	return g.vCount
}

// TODO: docs
func (g *G2) EdgeCount() int {
	return g.eCount
}

// TODO: docs
func (g *G2) GetVertex(i Item) (*GV2, int, bool) {
	idx, ok := g.sat[i]

	if !ok {
		return nil, -1, false
	}

	return &g.V[idx], idx, true
}

// TODO: docs
func (g *G2) AddVertex(i Item) (*GV2, error) {
	if idx, ok := g.sat[i]; ok {
		return &g.V[idx], ErrExists
	}

	g.V = append(g.V, GV2{Item: i})

	idx := len(g.V) - 1

	g.V[idx].Index = idx
	g.sat[i] = idx

	g.vCount++

	return &g.V[idx], nil
}

// TODO: docs
func (g *G2) RemoveVertex(i Item) error {
	iDel, ok := g.sat[i]

	if !ok {
		return ErrNoVtx
	}

	eCount := len(g.V[iDel].E)

	for i := range g.V {
		if i == iDel {
			continue
		}

		toDel := []int{}

		for j := range g.V[i].E {
			if g.V[i].E[j].Dst == iDel {
				toDel = append(toDel, j)
				continue
			}

			// TODO: explain fix
			if g.V[i].E[j].Src > iDel {
				g.V[i].E[j].Src--
			}

			// TODO: explain fix
			if g.V[i].E[j].Dst > iDel {
				g.V[i].E[j].Dst--
			}
		}

		for _, eIdx := range toDel {
			Cut(&g.V[i].E, eIdx)
			eCount++
		}
	}

	Cut(&g.V, iDel)
	delete(g.sat, i)

	// TODO: explain fix
	for i := iDel; i < len(g.V); i++ {
		g.V[i].Index = i
		g.sat[g.V[i].Item] = i
	}

	g.eCount -= eCount
	g.vCount--

	return nil
}

// TODO: docs
func (g *G2) GetEdge(src Item, dst Item) (*GE2, int, bool) {
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
func (g *G2) AddEdge(src Item, dst Item, wt float64) (*GE2, error) {
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

	vSrc.E = append(vSrc.E, GE2{Src: vSrc.Index, Dst: vDst.Index, Wt: wt})

	g.eCount++

	return &(vSrc.E[len(vSrc.E)-1]), nil
}

// TODO: docs
func (g *G2) RemoveEdge(src Item, dst Item) error {
	_, idx, ok := g.GetEdge(src, dst)

	if !ok {
		return ErrNoEdge
	}

	Cut(&g.V[g.sat[src]].E, idx)

	g.eCount--

	return nil
}

// TODO: docs
func Cut[T any](s *[]T, idx int) {
	if idx < 0 || idx >= len(*s) {
		return
	}

	copy((*s)[idx:], (*s)[idx+1:])

	// avoiding memory leak by assigning the
	// zero value to the duplicated position
	var zero T
	(*s)[len(*s)-1] = zero

	*s = (*s)[:len(*s)-1]
}
