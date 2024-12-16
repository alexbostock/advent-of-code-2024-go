package puzzle16

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseInput16(t *testing.T) {
	got := ParseInput16(strings.NewReader(`####
#S.E`))

	expected := &Maze{
		height: 2,
		width:  4,
		walls: map[Coords]bool{
			{0, 0}: true, {1, 0}: true, {2, 0}: true, {3, 0}: true,
			{0, 1}: true,
		},
		start: Coords{1, 1},
		end:   Coords{3, 1},
	}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestShortestPath(t *testing.T) {
	maze := ParseInput16(strings.NewReader(`###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############
`))

	score, positionsVisited := maze.ShortestPath()
	numPositionsVisited := len(positionsVisited)

	expectedScore := 7036
	expectedNumPositionsVisited := 45

	if expectedScore != score {
		t.Errorf("expected %v, got %v", expectedScore, score)
	}
	if expectedNumPositionsVisited != numPositionsVisited {
		t.Errorf("expected %v, got %v", expectedNumPositionsVisited, numPositionsVisited)
	}
}
