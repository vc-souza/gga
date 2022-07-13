package algo

import "github.com/vc-souza/gga/ds"

// TODO: docs
type DFSNode[V ds.Item] struct {
	// TODO: docs
	Discovery int

	// TODO: docs
	Finish int

	// TODO: docs
	Color int

	// TODO: docs
	Parent *V
}

// TODO: docs
type DFSForest[V ds.Item] map[*V]*DFSNode[V]

// TODO: docs
func DFS[V ds.Item](g *ds.Graph[V]) (DFSForest[V], *EdgeTypes[V], error) {
	fst := DFSForest[V]{}
	tps := &EdgeTypes[V]{}

	//

	return fst, tps, nil
}
