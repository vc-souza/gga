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
func (g *G2) AddVertex(i Item) (*GV2, error) {
	if idx, ok := g.sat[i]; ok {
		return &g.V[idx], ErrExists
	}

	g.V = append(g.V, GV2{Item: i})

	idx := len(g.V) - 1
	ptr := &g.V[idx]

	ptr.Index = idx
	g.sat[i] = idx

	g.vCount++

	return ptr, nil
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
