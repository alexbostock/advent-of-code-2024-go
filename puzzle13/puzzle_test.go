package puzzle13

import (
	"slices"
	"strings"
	"testing"
)

func TestParseInput13(t *testing.T) {
	got, err := ParseInput13(strings.NewReader(`Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176
`))

	expected := []Machine{
		{
			a:     Coords{94, 34},
			b:     Coords{22, 67},
			prize: Coords{8400, 5400},
		},
		{
			a:     Coords{26, 66},
			b:     Coords{67, 21},
			prize: Coords{12748, 12176},
		},
	}

	if err != nil {
		t.Errorf("expect nil error, got %v", err)
	}

	if !slices.Equal(expected, got) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestSearchMinimumTokensToWinAllPrizes(t *testing.T) {
	got := SearchMinimumTokensToWinAllPrizes([]Machine{
		{Coords{94, 34}, Coords{22, 67}, Coords{8400, 5400}},
		{Coords{26, 66}, Coords{67, 21}, Coords{12748, 12176}},
		{Coords{17, 86}, Coords{84, 37}, Coords{7870, 6450}},
		{Coords{69, 23}, Coords{27, 71}, Coords{18641, 10279}},
	})

	expected := 480

	if expected != got {
		t.Errorf("expected %v, got %v", expected, got)
	}
}
