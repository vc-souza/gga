package ds

import (
	"errors"
	"fmt"
	"testing"

	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestParseGraph(t *testing.T) {
	cases := []struct {
		desc      string
		addType   bool
		input     string
		err       error
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
			`,
			err:       nil,
			vertCount: 3,
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
			err:       nil,
			vertCount: 3,
			edgeCount: 4,
		},
		{
			desc:    "no graph type",
			addType: false,
			input: `
			a#b,c
			b#a
			c#a
			`,
			err: ErrInvalidSer,
		},
		{
			desc:    "bad graph type",
			addType: false,
			input: `
			random
			a#b,c
			b#a
			c#a
			`,
			err: ErrInvalidSer,
		},
		{
			desc:    "bad adj list",
			addType: true,
			input: `
			a#b,c#
			b#a
			c#a
			`,
			err: ErrInvalidSer,
		},
		{
			desc:    "bad vertex name",
			addType: true,
			input: `
			a::#b,c
			b#a::
			c#a::
			`,
			err: ErrInvalidSer,
		},
		{
			desc:    "empty vertex name",
			addType: true,
			input: `
			#b,c
			`,
			err: ErrInvalidSer,
		},
		{
			desc:    "bad edge name",
			addType: true,
			input: `
			a#x:10:*,c
			b#
			c#a
			`,
			err: ErrInvalidSer,
		},
		{
			desc:    "empty edge name",
			addType: true,
			input: `
			a#,c
			b#
			c#a
			`,
			err: ErrInvalidSer,
		},
	}

	for _, tc := range cases {
		for _, gType := range []string{UndirectedGraphKey, DirectedGraphKey} {
			t.Run(tagGraphTest(gType, tc.desc), func(t *testing.T) {
				var input string

				if tc.addType {
					input = fmt.Sprintf("%s\n%s", gType, tc.input)
				} else {
					input = tc.input
				}

				g, err := ParseGraph(input)

				if tc.err != nil {
					ut.AssertEqual(t, true, errors.Is(err, tc.err))
					return
				}

				if gType == UndirectedGraphKey {
					ut.AssertEqual(t, true, g.Undirected())
				} else {
					ut.AssertEqual(t, true, g.Directed())
				}

				ut.AssertEqual(t, tc.vertCount, g.VertexCount())

				if g.Directed() {
					ut.AssertEqual(t, tc.edgeCount, g.EdgeCount())
				} else {
					ut.AssertEqual(t, tc.edgeCount/2, g.EdgeCount())
				}
			})
		}
	}
}
