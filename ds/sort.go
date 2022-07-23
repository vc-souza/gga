package ds

/*
ByEdgeWeight implements the sort.Interface, enabling the sorting
of a list of graph edges in order of non-decreasing edge weights.
*/
type ByEdgeWeight []*GE

func (b ByEdgeWeight) Len() int           { return len(b) }
func (b ByEdgeWeight) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByEdgeWeight) Less(i, j int) bool { return b[i].Wt < b[j].Wt }
