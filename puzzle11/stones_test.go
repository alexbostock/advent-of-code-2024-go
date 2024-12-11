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

func TestCountStonesAfterNumBlinks(t *testing.T) {
	stones := []int{125, 17}
	got := CountStonesAfterNumBlinks(stones, 6)
	expected := 22

	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestTransformStonesOnBlink(t *testing.T) {
	stones := []int{0, 1, 10, 99, 999}
	expected := []int{1, 2024, 1, 0, 9, 9, 2021976}
	got := transformStonesOnBlink(stones)

	if !slices.Equal(got, expected) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}
