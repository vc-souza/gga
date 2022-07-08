package testutils

// ID is a test struct that implements the ds.Item interface.
type ID string

func (i ID) Label() string {
	return string(i)
}
