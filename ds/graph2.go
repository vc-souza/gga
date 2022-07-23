package ds

// TODO: docs
type GE2 struct {
	Formattable

	Src int
	Dst int
	Wt  float64
}

// TODO: docs
type GV2 struct {
	Formattable

	Item  Item
	E     []GE2
	Index int
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

			// TODO: explain
			if g.V[i].E[j].Src > iDel {
				g.V[i].E[j].Src--
			}

			// TODO: explain
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
