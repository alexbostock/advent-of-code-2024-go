package puzzle14

import (
	"slices"
	"strings"
	"testing"
)

func TestParseInput14(t *testing.T) {
	got := ParseInput14(strings.NewReader(`p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3
`))

	expected := []Robot{
		{Vector{0, 4}, Vector{3, -3}},
		{Vector{6, 3}, Vector{-1, -3}},
		{Vector{10, 3}, Vector{-1, 2}},
		{Vector{2, 0}, Vector{2, -1}},
		{Vector{0, 0}, Vector{1, 3}},
		{Vector{3, 0}, Vector{-2, -2}},
		{Vector{7, 6}, Vector{-1, -3}},
		{Vector{3, 0}, Vector{-1, -2}},
		{Vector{9, 3}, Vector{2, 3}},
		{Vector{7, 3}, Vector{-1, 2}},
		{Vector{2, 4}, Vector{2, -3}},
		{Vector{9, 5}, Vector{-3, -3}},
	}

	if !slices.Equal(expected, got) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestStateAfterSeconds(t *testing.T) {
	startPos := []Robot{{Vector{0, 4}, Vector{3, -3}},
		{Vector{6, 3}, Vector{-1, -3}},
		{Vector{10, 3}, Vector{-1, 2}},
		{Vector{2, 0}, Vector{2, -1}},
		{Vector{0, 0}, Vector{1, 3}},
		{Vector{3, 0}, Vector{-2, -2}},
		{Vector{7, 6}, Vector{-1, -3}},
		{Vector{3, 0}, Vector{-1, -2}},
		{Vector{9, 3}, Vector{2, 3}},
		{Vector{7, 3}, Vector{-1, 2}},
		{Vector{2, 4}, Vector{2, -3}},
		{Vector{9, 5}, Vector{-3, -3}},
	}

	got := StateAfterSeconds(startPos, 100, 11, 7)
	safetyFactor := SafetyFactor(got, 11, 7)

	if safetyFactor != 12 {
		t.Errorf("expected safety factor 12, got %v (%v)", got, safetyFactor)
	}
}
