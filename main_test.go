package main

import (
	"testing"

	"github.com/matryer/is"
)

func TestGameOfLife(t *testing.T) {
	cases := map[string]struct {
		in  []string
		exp []string
	}{
		"dead cell with 0 neighbour => 1 dead cell": {
			in: []string{
				".",
			},
			exp: []string{
				".",
			},
		},

		"1 cell alive with 0 neighbour => 1 dead cell": {
			in: []string{
				"*",
			},
			exp: []string{
				".",
			},
		},

		"1 living cell with 1 neighbour => 2 dead cells": {
			in: []string{
				"**",
			},
			exp: []string{
				"..",
			},
		},

		"2 living cells with only 1 neighbour => all 3 dead cells": {
			in: []string{
				"**.",
			},
			exp: []string{
				"...",
			},
		},

		"1 living cell with 2 neighbours => 2 dead cells and 1 living cell": {
			in: []string{
				"***",
			},
			exp: []string{
				".*.",
			},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			got := GameOfLife(c.in)
			is := is.New(t)
			is.Equal(got, c.exp)
		})
	}
}
