package ds

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestTextLabel(t *testing.T) {
	ut.AssertEqual(t, "test", Text("test").Label())
}

func TestTextParser(t *testing.T) {
	cases := []struct {
		desc      string
		addType   bool
		input     string
		err       string
		vertOrder string
		vertCount int
		edgeCount int
	}{
		{
			desc:    "unweighted",
			addType: true,
			input: `
			a#b,c
			b#a
			c#a
			d#
			`,
			vertOrder: "a,b,c,d",
			vertCount: 4,
			edgeCount: 4,
		},
		{
			desc:    "weighted",
			addType: true,
			input: `
			a#b:10,c:5
			b#a:10
			c#a:5
			`,
			vertOrder: "a,b,c",
			vertCount: 3,
			edgeCount: 4,
		},
		{
			desc:    "bad weight",
			err:     "weight: bad value",
			addType: true,
			input: `
			a#b:10&&,c:5
			b#a:10&&
			c#a:5
			`,
		},
		{
			desc:    "no graph type",
			err:     "graph type: bad name",
			addType: false,
			input: `
			a#b,c
			b#a
			c#a
			`,
		},
		{
			desc:    "bad graph type",
			err:     "graph type: bad name",
			addType: false,
			input: `
			random
			a#b,c
			b#a
			c#a
			`,
		},
		{
			desc:    "bad adj list",
			err:     "adjacency list: wrong item count",
			addType: true,
			input: `
			a#b,c#
			b#a
			c#a
			`,
		},
		{
			desc:    "bad vertex name",
			err:     "vertex: bad name",
			addType: true,
			input: `
			a::#b,c
			b#a::
			c#a::
			`,
		},
		{
			desc:    "empty vertex name",
			err:     "vertex: empty name",
			addType: true,
			input: `
			#b,c
			`,
		},
		{
			desc:    "bad edge name",
			err:     "edge: wrong item count",
			addType: true,
			input: `
			a#x:10:*,c
			b#
			c#a
			`,
		},
		{
			desc:    "empty edge name",
			err:     "edge: empty",
			addType: true,
			input: `
			a#,c
			b#
			c#a
			`,
		},
		{
			desc:    "undeclared vertex",
			err:     "edge: unknown destination",
			addType: true,
			input: `
			a#b,c
			b#a,d
			c#a
			`,
		},
	}

	for _, tc := range cases {
		for _, gType := range []string{undirectedGraphKey, directedGraphKey} {
			t.Run(tagGraphTest(gType, tc.desc), func(t *testing.T) {
				var input string

				if tc.addType {
					input = fmt.Sprintf("%s\n%s", gType, tc.input)
				} else {
					input = tc.input
				}

				g, vars, err := Parse(input)

				if len(tc.err) != 0 {
					ut.AssertTrue(t, errors.As(err, new(ErrInvalidSer)))
					ut.AssertTrue(t, strings.Contains(err.Error(), tc.err))
					return
				}

				if gType == undirectedGraphKey {
					ut.AssertTrue(t, g.Undirected())
				} else {
					ut.AssertTrue(t, g.Directed())
				}

				ut.AssertEqual(t, tc.vertCount, g.VertexCount())
				ut.AssertEqual(t, tc.vertCount, len(vars))

				// expected order by the test case
				expectVerts := strings.Split(tc.vertOrder, ",")

				// making sure the test case is sane
				ut.AssertEqual(t, tc.vertCount, len(expectVerts))

				// actual order in the graph (insertion)
				for i := 0; i < tc.vertCount; i++ {
					ut.AssertEqual(t, expectVerts[i], g.V[i].Label())
				}

				ut.AssertEqual(t, tc.edgeCount, g.EdgeCount())
			})
		}
	}
}
