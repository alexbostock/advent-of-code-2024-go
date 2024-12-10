package puzzle10

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseInput10(t *testing.T) {
	got, err := ParseInput10(strings.NewReader(`0123
1234
8765
9876
`))
	expected := [][]int{
		{0, 1, 2, 3},
		{1, 2, 3, 4},
		{8, 7, 6, 5},
		{9, 8, 7, 6},
	}

	if err != nil {
		t.Error("expected nil error")
	}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestSumScoresOfTrailheads(t *testing.T) {
	area := [][]int{
		{8, 9, 0, 1, 0, 1, 2, 3},
		{7, 8, 1, 2, 1, 8, 7, 4},
		{8, 7, 4, 3, 0, 9, 6, 5},
		{9, 6, 5, 4, 9, 8, 7, 4},
		{4, 5, 6, 7, 8, 9, 0, 3},
		{3, 2, 0, 1, 9, 0, 1, 2},
		{0, 1, 3, 2, 9, 8, 0, 1},
		{1, 0, 4, 5, 6, 7, 3, 2},
	}

	got := SumScoresOfTrailheads(area)
	expected := 36

	if expected != got {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestScoreTrailhead(t *testing.T) {
	area := [][]int{
		{8, 9, 0, 1, 0, 1, 2, 3},
		{7, 8, 1, 2, 1, 8, 7, 4},
		{8, 7, 4, 3, 0, 9, 6, 5},
		{9, 6, 5, 4, 9, 8, 7, 4},
		{4, 5, 6, 7, 8, 9, 0, 3},
		{3, 2, 0, 1, 9, 0, 1, 2},
		{0, 1, 3, 2, 9, 8, 0, 1},
		{1, 0, 4, 5, 6, 7, 3, 2},
	}

	examples := []struct {
		pos   Coords
		score int
	}{
		{Coords{0, 1}, 1},
		{Coords{0, 0}, 1},
		{Coords{2, 1}, 2},

		{Coords{0, 2}, 5},
		{Coords{0, 4}, 6},
		{Coords{2, 4}, 5},
		{Coords{4, 6}, 3},
		{Coords{5, 2}, 1},
		{Coords{5, 5}, 3},
		{Coords{6, 0}, 5},
		{Coords{6, 6}, 3},
		{Coords{7, 1}, 5},
	}

	for _, example := range examples {
		got := scoreTrailhead(area, example.pos)

		if got != example.score {
			t.Errorf("%v: expected %v, got %v", example.pos, example.score, got)
		}
	}
}

func TestCountDistinctTrails(t *testing.T) {
	area := [][]int{
		{8, 9, 0, 1, 0, 1, 2, 3},
		{7, 8, 1, 2, 1, 8, 7, 4},
		{8, 7, 4, 3, 0, 9, 6, 5},
		{9, 6, 5, 4, 9, 8, 7, 4},
		{4, 5, 6, 7, 8, 9, 0, 3},
		{3, 2, 0, 1, 9, 0, 1, 2},
		{0, 1, 3, 2, 9, 8, 0, 1},
		{1, 0, 4, 5, 6, 7, 3, 2},
	}

	expected := 81
	got := CountDistinctTrails(area)

	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}
}
