package puzzle12

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseInput12(t *testing.T) {
	got := ParseInput12(strings.NewReader(`AAAA
BBCD
BBCC
EEEC
`))
	expected := Area{
		height: 4,
		width:  4,
		plots: [][]rune{
			{'A', 'A', 'A', 'A'},
			{'B', 'B', 'C', 'D'},
			{'B', 'B', 'C', 'C'},
			{'E', 'E', 'E', 'C'},
		},
	}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestCostFences(t *testing.T) {
	got := CostFences(ParseInput12(strings.NewReader(`AAAA
BBCD
BBCC
EEEC
`)))

	expected := 140

	if expected != got {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestCostFencesNumSides(t *testing.T) {
	examples := []struct {
		input    string
		expected int
	}{
		{
			`AAAA
BBCD
BBCC
EEEC
`, 80,
		},
		{
			`OOOOO
OXOXO
OOOOO
OXOXO
OOOOO
`, 436,
		},
		{
			`EEEEE
EXXXX
EEEEE
EXXXX
EEEEE
`, 236,
		},
		{
			`AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA
`, 368,
		},
	}

	for _, example := range examples {
		got := CostFencesNumSides(ParseInput12(strings.NewReader(example.input)))
		if got != example.expected {
			t.Errorf("expected %v, got %v", example.expected, got)
		}
	}
}
