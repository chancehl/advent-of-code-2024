package random

import (
	"slices"
	"testing"
)

// This is the dumbest fucking test in the world, but the math is on my side here.
//
// I believe the chances of not picking a specific option are (3/4)^1000 which is
// so astronomincally small that I won't have to worry about this failing.
func TestChoice(t *testing.T) {
	options := []string{"a", "b", "c", "d"}
	choices := []string{}

	for i := 0; i < 1000; i += 1 {
		choices = append(choices, Choice(options))
	}

	for _, option := range options {
		if !slices.Contains(choices, option) {
			t.Errorf("option %s was not in chosen values %s", option, choices)
		}
	}
}
