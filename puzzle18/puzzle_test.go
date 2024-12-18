package puzzle18

import (
	"slices"
	"strings"
	"testing"
)

var exampleInput = `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0
`

func TestParseInput18(t *testing.T) {
	got := ParseInput18(strings.NewReader(exampleInput))

	expected := []Coords{
		{5, 4},
		{4, 2},
		{4, 5},
		{3, 0},
		{2, 1},
		{6, 3},
		{2, 4},
		{1, 5},
		{0, 6},
		{3, 3},
		{2, 6},
		{5, 1},
		{1, 2},
		{5, 5},
		{2, 5},
		{6, 5},
		{1, 4},
		{0, 4},
		{6, 4},
		{1, 1},
		{6, 1},
		{1, 0},
		{0, 5},
		{1, 6},
		{2, 0},
	}

	if !slices.Equal(expected, got) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestShortestPathAfterBytesFallen(t *testing.T) {
	examples := []struct {
		numBytes int
		expected int
	}{
		{12, 22},
		{21, -1},
	}

	bytes := ParseInput18(strings.NewReader(exampleInput))

	for _, example := range examples {
		got := ShortestPathAfterBytesFallen(bytes, 7, 7, example.numBytes)
		if example.expected != got {
			t.Errorf("expected %v, got %v", example.expected, got)
		}
	}
}

func TestFindFirstByteObstructingExit(t *testing.T) {
	bytes := ParseInput18(strings.NewReader(exampleInput))
	got := FindFirstByteObstructingExit(bytes, 7, 7)
	expected := "6,1"
	if expected != got {
		t.Errorf("expected %v, got %v", expected, got)
	}
}
