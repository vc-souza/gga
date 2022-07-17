package ds

import (
	"testing"

	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestRemoveFromPointersSlice(t *testing.T) {
	var v1 = 1
	var v2 = 2
	var v3 = 3

	cases := []struct {
		desc   string
		slice  []*int
		calls  []int
		expect []*int
	}{
		{
			desc:   "empty slice",
			slice:  []*int{},
			calls:  []int{0},
			expect: []*int{},
		},
		{
			desc:   "nil slice",
			slice:  nil,
			calls:  []int{0},
			expect: nil,
		},
		{
			desc:   "index too low",
			slice:  []*int{&v1},
			calls:  []int{-1},
			expect: []*int{&v1},
		},
		{
			desc:   "index too high",
			slice:  []*int{&v1},
			calls:  []int{1},
			expect: []*int{&v1},
		},
		{
			desc:   "first index",
			slice:  []*int{&v1, &v2, &v3},
			calls:  []int{0},
			expect: []*int{&v2, &v3},
		},
		{
			desc:   "middle index",
			slice:  []*int{&v1, &v2, &v3},
			calls:  []int{1},
			expect: []*int{&v1, &v3},
		},
		{
			desc:   "last index",
			slice:  []*int{&v1, &v2, &v3},
			calls:  []int{2},
			expect: []*int{&v1, &v2},
		},
		{
			desc:   "remove all",
			slice:  []*int{&v1, &v2, &v3},
			calls:  []int{2, 1, 0},
			expect: []*int{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			s := tc.slice

			for _, idx := range tc.calls {
				s = RemoveFromPointersSlice(s, idx)
			}

			ut.Equal(t, len(tc.expect), len(s))

			for i := 0; i < len(tc.expect); i++ {
				ut.Equal(t, tc.expect[i], s[i])
			}
		})
	}
}
