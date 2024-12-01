package main

import (
	"reflect"
	"slices"
	"strings"
	"testing"
)

var leftExample = []int{3, 4, 2, 1, 3, 3}
var rightExample = []int{4, 3, 5, 3, 9, 3}

func TestParseInput(t *testing.T) {
	left, right, err := ParseInput(strings.NewReader(`3   4
4   3
2   5
1   3
3   9
3   3
`))

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if !reflect.DeepEqual(left, leftExample) {
		t.Errorf("expected %v, got %v", leftExample, left)
	}

	if !reflect.DeepEqual(right, rightExample) {
		t.Errorf("expected %v, got %v", rightExample, right)
	}
}

func TestDiffLocations(t *testing.T) {
	distance := DiffLocations(slices.Clone(leftExample), slices.Clone(rightExample))
	if distance != 11 {
		t.Errorf("expected 11, got %d", distance)
	}
}

func TestCalculateSimilarityScore(t *testing.T) {
	score := CalculateSimilarityScore(slices.Clone(leftExample), slices.Clone(rightExample))
	if score != 31 {
		t.Errorf("expected 31, got %d", score)
	}
}
