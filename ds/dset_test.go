package ds

import (
	"testing"

	ut "github.com/vc-souza/gga/internal/testutils"
)

func TestDSetMakeSet(t *testing.T) {
	i := 10

	dset := NewDSet[*int]()

	dset.MakeSet(&i)

	ut.Equal(t, &i, dset.FindSet(&i))
}

func TestDSetFindSet_same_set(t *testing.T) {
	vars := []int{0, 1}

	dset := NewDSet[*int]()

	dset.MakeSet(&vars[0])
	dset.MakeSet(&vars[1])

	dset.Union(&vars[0], &vars[1])

	ut.True(t, dset.FindSet(&vars[0]) == dset.FindSet(&vars[1]))
}

func TestDSetFindSet_different_set(t *testing.T) {
	vars := []int{0, 1}

	dset := NewDSet[*int]()

	dset.MakeSet(&vars[0])
	dset.MakeSet(&vars[1])

	ut.True(t, dset.FindSet(&vars[0]) != dset.FindSet(&vars[1]))
}

func TestDSetUnion(t *testing.T) {
	vars := []int{0, 1, 2, 3, 4}
	dset := NewDSet[*int]()

	// {0} {1} {2} {3} {4}
	for i := range vars {
		dset.MakeSet(&vars[i])
	}

	// {0, 1}
	dset.Union(&vars[0], &vars[1])

	// {2, 3}
	dset.Union(&vars[2], &vars[3])

	// {2, 3, 4}
	dset.Union(&vars[2], &vars[4])

	// {0, 1, 2, 3, 4}
	dset.Union(&vars[0], &vars[4])

	// same set, so all should have
	// the same representative
	target := dset.FindSet(&vars[0])

	for i := range vars {
		ut.Equal(t, target, dset.FindSet(&vars[i]))
	}
}
