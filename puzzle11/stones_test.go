package puzzle11

import (
	"slices"
	"strings"
	"testing"
)

func TestParseInput(t *testing.T) {
	got, err := ParseInput11(strings.NewReader(`0 1 10 99 999
`))
	expected := []int{0, 1, 10, 99, 999}

	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	if !slices.Equal(got, expected) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestCountAllStonesAfterNumBlinks(t *testing.T) {
	stones := []int{125, 17}
	got := CountAllStonesAfterNumBlinks(stones, 6)
	expected := 22

	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestTransformStoneOnBlink(t *testing.T) {
	examples := []struct {
		stone    int
		expected []int
	}{
		{0, []int{1}},
		{1, []int{2024}},
		{10, []int{1, 0}},
		{99, []int{9, 9}},
		{999, []int{2021976}},
	}

	for _, example := range examples {
		got := transformStoneOnBlink(example.stone)
		if !slices.Equal(got, example.expected) {
			t.Errorf("expected %v, got %v", example.expected, got)
		}
	}
}
