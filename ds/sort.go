package ds

// TODO: docs (non-decreasing edge weights)
type ByEdgeWeight[T Item] []*GE[T]

func (b ByEdgeWeight[T]) Len() int           { return len(b) }
func (b ByEdgeWeight[T]) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByEdgeWeight[T]) Less(i, j int) bool { return b[i].Wt < b[j].Wt }
