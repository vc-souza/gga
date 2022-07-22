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

func (g *G2) RemoveEdge(src Item, dst Item) error {
	_, idx, ok := g.GetEdge(src, dst)

	if !ok {
		return ErrNoEdge
	}

	edges := &g.V[g.sat[src]].E

	copy((*edges)[idx:], (*edges)[idx+1:])

	// avoiding memory leaks
	(*edges)[len(*edges)-1] = GE2{}

	*edges = (*edges)[:len(*edges)-1]

	g.eCount--

	return nil
}
