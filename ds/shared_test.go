package ds

import (
	"testing"

	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestCut(t *testing.T) {
	cases := []struct {
		desc   string
		slice  []int
		calls  []int
		expect []int
	}{
		{
			desc:   "empty slice",
			slice:  []int{},
			calls:  []int{0},
			expect: []int{},
		},
		{
			desc:   "nil slice",
			slice:  nil,
			calls:  []int{0},
			expect: nil,
		},
		{
			desc:   "index too low",
			slice:  []int{1},
			calls:  []int{-1},
			expect: []int{1},
		},
		{
			desc:   "index too high",
			slice:  []int{1},
			calls:  []int{1},
			expect: []int{1},
		},
		{
			desc:   "first index",
			slice:  []int{1, 2, 3},
			calls:  []int{0},
			expect: []int{2, 3},
		},
		{
			desc:   "middle index",
			slice:  []int{1, 2, 3},
			calls:  []int{1},
			expect: []int{1, 3},
		},
		{
			desc:   "last index",
			slice:  []int{1, 2, 3},
			calls:  []int{2},
			expect: []int{1, 2},
		},
		{
			desc:   "remove all",
			slice:  []int{1, 2, 3},
			calls:  []int{1, 1, 0},
			expect: []int{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			s := tc.slice

			for _, idx := range tc.calls {
				Cut(&s, idx)
			}

			ut.Equal(t, len(tc.expect), len(s))

			for i := 0; i < len(tc.expect); i++ {
				ut.Equal(t, tc.expect[i], s[i])
			}
		})
	}
}
