//go:build !test

package main

type letter string

func (l letter) Label() string {
	return string(l)
}
