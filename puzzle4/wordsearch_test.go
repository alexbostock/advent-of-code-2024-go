package puzzle4

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseInput4(t *testing.T) {
	got := ParseInput4(strings.NewReader(`..X...
.SAMX.
.A..A.
XMAS.S
.X....
`))

	expected := WordSearch{
		chars: [][]rune{
			{'.', '.', 'X', '.', '.', '.'},
			{'.', 'S', 'A', 'M', 'X', '.'},
			{'.', 'A', '.', '.', 'A', '.'},
			{'X', 'M', 'A', 'S', '.', 'S'},
			{'.', 'X', '.', '.', '.', '.'},
		},
		height: 5,
		width:  6,
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestCountXMASInWordSearch(t *testing.T) {
	wordSearch := ParseInput4(strings.NewReader(`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`))

	got := CountXMASInWordSearch(wordSearch)

	if got != 18 {
		t.Errorf("expected 18, got %v", got)
	}
}

func TestCountCrossMASInWordSearch(t *testing.T) {
	wordSearch := ParseInput4(strings.NewReader(`.M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
..........`))

	got := CountCrossMASInWordSearch(wordSearch)

	if got != 9 {
		t.Errorf("expected 9, got %v", got)
	}
}
