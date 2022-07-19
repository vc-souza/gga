package testutils

import "strconv"

// BenchItem is a ds.Item implementation used by benchmarks.
type BenchItem int

func (i BenchItem) Label() string { return strconv.Itoa(int(i)) }
