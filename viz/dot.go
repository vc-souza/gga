package viz

import (
	"fmt"
	"strings"
)

type DotAttrs map[string]string

func (a DotAttrs) String() string {
	if len(a) == 0 {
		return ""
	}

	s := []string{"[\n"}

	for k, v := range a {
		s = append(s, fmt.Sprintf("%s=%s", k, v))
	}

	s = append(s, "\n]")

	return strings.Join(s, " ")
}

// TODO: one master interface embedded by all?
// TODO: one interface per type? (graph, subgraph, edge, node/vertex)
// TODO: visitor?
