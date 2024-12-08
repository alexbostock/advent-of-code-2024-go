package puzzle8

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseInput8(t *testing.T) {
	got := ParseInput8(strings.NewReader(`............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............
`))

	expected := Puzzle8Input{
		height: 12,
		width:  12,
		antennas: map[rune][]Coords{
			'0': {{1, 8}, {2, 5}, {3, 7}, {4, 4}},
			'A': {{5, 6}, {8, 8}, {9, 9}},
		},
	}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestCountDistinctAntinodes(t *testing.T) {
	got := CountDistinctAntinodes(Puzzle8Input{
		height: 12,
		width:  12,
		antennas: map[rune][]Coords{
			'0': {{1, 8}, {2, 5}, {3, 7}, {4, 4}},
			'A': {{5, 6}, {8, 8}, {9, 9}},
		},
	})

	if got != 14 {
		t.Errorf("expected 14, got %v", got)
	}
}

func TestCountDistinctAntinodesWithResonantHarmonics(t *testing.T) {
	got := CountDistinctAntinodesWithResonantHarmonics(Puzzle8Input{
		height: 12,
		width:  12,
		antennas: map[rune][]Coords{
			'0': {{1, 8}, {2, 5}, {3, 7}, {4, 4}},
			'A': {{5, 6}, {8, 8}, {9, 9}},
		},
	})

	if got != 34 {
		t.Errorf("expected 34, got %v", got)
	}
}
