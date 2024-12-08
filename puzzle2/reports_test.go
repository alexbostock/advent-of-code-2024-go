package puzzle2

import (
	"reflect"
	"slices"
	"strings"
	"testing"
)

func TestParseInput2(t *testing.T) {
	reports, err := ParseInput2(strings.NewReader(`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`))

	expected := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if !reflect.DeepEqual(reports, expected) {
		t.Errorf("expected %v, got %v", expected, reports)
	}
}

func TestIsSafe(t *testing.T) {
	examples := []struct {
		report   []int
		expected bool
	}{
		{[]int{7, 6, 4, 2, 1}, true},
		{[]int{1, 2, 7, 8, 9}, false},
		{[]int{9, 7, 6, 2, 1}, false},
		{[]int{1, 3, 2, 4, 5}, false},
		{[]int{8, 6, 4, 4, 1}, false},
		{[]int{1, 3, 6, 7, 9}, true},
	}

	for _, example := range examples {
		got := IsSafe(example.report, -1)
		if got != example.expected {
			t.Errorf("expected %v, got %v", example.expected, got)
		}
	}
}

func TestIsSafeDropIndex(t *testing.T) {
	examples := []struct {
		report    []int
		dropIndex int
		expected  bool
	}{
		{[]int{6, 1, 7}, -1, false},
		{[]int{6, 1, 7}, 0, false},
		{[]int{6, 1, 7}, 1, true},
		{[]int{6, 1, 7}, 2, false},
	}

	for _, example := range examples {
		got := IsSafe(example.report, example.dropIndex)
		if got != example.expected {
			t.Errorf("expected %v, got %v", example.expected, got)
		}
	}
}

func TestIsSafeWithProblemDampener(t *testing.T) {
	examples := []struct {
		report   []int
		expected bool
	}{

		{[]int{7, 6, 4, 2, 1}, true},
		{[]int{1, 2, 7, 8, 9}, false},
		{[]int{9, 7, 6, 2, 1}, false},
		{[]int{1, 3, 2, 4, 5}, true},
		{[]int{8, 6, 4, 4, 1}, true},
		{[]int{1, 3, 6, 7, 9}, true},
		{[]int{1, 6, 7}, true},
		{[]int{1, 7, 6}, true},
		{[]int{6, 1, 7}, true},
		{[]int{6, 7, 1}, true},
		{[]int{7, 1, 6}, true},
		{[]int{7, 6, 1}, true},
	}

	for _, example := range examples {
		reportUnmutated := slices.Clone(example.report)
		got := IsSafeWithProblemDampener(example.report)
		if got != example.expected {
			t.Errorf("%v: expected %v, got %v", reportUnmutated, example.expected, got)
		}
	}
}
