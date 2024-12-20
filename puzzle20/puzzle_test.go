package puzzle20

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseInput20(t *testing.T) {
	got := ParseInput20(strings.NewReader(`####
S #E
`))

	expected := Maze{
		4, 2,
		Coords{0, 1}, Coords{3, 1},
		map[Coords]bool{{0, 0}: true, {1, 0}: true, {2, 0}: true, {3, 0}: true, {2, 1}: true},
	}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected %+v, got %+v", expected, got)
	}
}

var exampleMaze = ParseInput20(strings.NewReader(`###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############
`))

func TestPathWithoutCheats(t *testing.T) {
	got := exampleMaze.pathWithoutCheats()
	if len(got)-1 != 84 {
		t.Errorf("expected path length 84, got %d", len(got)-1)
	}

	checks := []struct {
		offset   int
		expected Coords
	}{
		{0, Coords{1, 3}},
		{84, Coords{5, 7}},
		{83, Coords{4, 7}},
	}
	for _, check := range checks {
		if got[check.offset] != check.expected {
			t.Errorf("expected %+v at %dth position in path, got %+v", check.expected, check.offset, got[check.offset])
		}
	}
}

func TestNumCheatsSavingAtLeast(t *testing.T) {
	examples := []struct{ numCheatsAllowed, threshold, expected int }{
		{2, 64, 1},
		{2, 40, 2},
		{2, 20, 5},
		{2, 12, 8},
		{20, 76, 3},
		{20, 74, 7},
	}

	for _, example := range examples {
		got := exampleMaze.NumCheatsSavingAtLeast(example.numCheatsAllowed, example.threshold)
		if example.expected != got {
			t.Errorf("%d / %d: expected %d, got %d", example.numCheatsAllowed, example.threshold, example.expected, got)
		}
	}
}
