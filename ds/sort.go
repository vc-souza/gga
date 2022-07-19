package ds

import "sort"

// TODO: docs
func SortEdgesByWeight[V Item](es []*GE[V], asc bool) {
	sort.Slice(es, func(i, j int) bool {
		if asc {
			return es[i].Wt < es[j].Wt
		} else {
			return es[i].Wt > es[j].Wt
		}
	})
}
