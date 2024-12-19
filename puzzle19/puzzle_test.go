package puzzle19

import (
	"reflect"
	"strings"
	"testing"
)

var exampleInput = `r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb
`

func TestParseInput19(t *testing.T) {
	got := ParseInput19(strings.NewReader(exampleInput))
	expected := Input{
		towels: []string{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"},
		designs: []string{
			"brwrr",
			"bggr",
			"gbbr",
			"rrbgbr",
			"ubwu",
			"bwurrg",
			"brgr",
			"bbrgwb",
		},
	}
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected %+v, got %+v", expected, got)
	}
}

func TestCountPossibleDesigns(t *testing.T) {
	input := ParseInput19(strings.NewReader(exampleInput))
	got := CountPossibleDesigns(input)
	expected := 6
	if expected != got {
		t.Errorf("expected %d, got %d", expected, got)
	}
}

func TestSumPossibleWaysOfMakingDesigns(t *testing.T) {
	input := ParseInput19(strings.NewReader(exampleInput))
	got := SumPossibleWaysOfMakingDesigns(input)
	expected := 16
	if expected != got {
		t.Errorf("expected %d, got %d", expected, got)
	}
}
