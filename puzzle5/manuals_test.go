package puzzle5

import (
	"reflect"
	"strings"
	"testing"
)

const puzzle5ExampleInput = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47
`

func TestParseInput5(t *testing.T) {
	got := ParseInput5(strings.NewReader(puzzle5ExampleInput))
	expected := Puzzle5Input{
		rules: []pageOrderingRule{
			{47, 53},
			{97, 13},
			{97, 61},
			{97, 47},
			{75, 29},
			{61, 13},
			{75, 53},
			{29, 13},
			{97, 29},
			{53, 29},
			{61, 53},
			{97, 53},
			{61, 29},
			{47, 13},
			{75, 47},
			{97, 75},
			{47, 61},
			{75, 61},
			{47, 29},
			{75, 13},
			{53, 13},
		},
		updates: [][]int{
			{75, 47, 61, 53, 29},
			{97, 61, 53, 29, 13},
			{75, 29, 13},
			{75, 97, 47, 61, 53},
			{61, 13, 29},
			{97, 13, 75, 29, 47},
		},
	}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestSumMiddlePagesOfCorrectlyOrderedUpdates(t *testing.T) {
	input := ParseInput5(strings.NewReader(puzzle5ExampleInput))
	got := SumMiddlePagesOfCorrectlyOrderedUpdates(input)
	expected := 143
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestLoadPageOrderingRules(t *testing.T) {
	got := loadPageOrderingRules([]pageOrderingRule{{11, 22}, {11, 44}, {33, 44}})
	expected := pageOrderingRules{
		rules: map[pageOrderingRule]bool{
			{11, 22}: true,
			{11, 44}: true,
			{33, 44}: true,
		},
		byBeforePage: map[int][]int{11: {0, 1}, 33: {2}},
		byAfterPage:  map[int][]int{22: {0}, 44: {1, 2}},
	}
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestIsValidUpdate(t *testing.T) {
	input := ParseInput5(strings.NewReader(puzzle5ExampleInput))
	rules := loadPageOrderingRules(input.rules)

	examples := []struct {
		update   []int
		expected bool
	}{
		{[]int{75, 47, 61, 53, 29}, true},
		{[]int{97, 61, 53, 29, 13}, true},
		{[]int{75, 29, 13}, true},
		{[]int{75, 97, 47, 61, 53}, false},
		{[]int{61, 13, 29}, false},
		{[]int{97, 13, 75, 29, 47}, false},
	}

	for _, example := range examples {
		got := isValidUpdate(example.update, rules)
		if got != example.expected {
			t.Errorf("%v: expected %v, got %v", example.update, example.expected, got)
		}
	}
}

func TestMiddlepageInUpdate(t *testing.T) {
	examples := []struct {
		update   []int
		expected int
	}{
		{[]int{75, 47, 61, 53, 29}, 61},
		{[]int{97, 61, 53, 29, 13}, 53},
		{[]int{75, 29, 13}, 29},
	}

	for _, example := range examples {
		got := middlePageInUpdate(example.update)
		if got != example.expected {
			t.Errorf("expected %v, got %v", example.expected, got)
		}
	}
}

func TestSumMiddlePagesOfFixedUpdates(t *testing.T) {
	input := ParseInput5(strings.NewReader(puzzle5ExampleInput))
	got := SumMiddlePagesOfFixedUpdates(input)
	expected := 123
	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestFixUpdate(t *testing.T) {
	input := ParseInput5(strings.NewReader(puzzle5ExampleInput))
	rules := loadPageOrderingRules(input.rules)

	examples := []struct {
		update   []int
		expected []int
	}{
		{[]int{75, 97, 47, 61, 53}, []int{97, 75, 47, 61, 53}},
		{[]int{61, 13, 29}, []int{61, 29, 13}},
		{[]int{97, 13, 75, 29, 47}, []int{97, 75, 47, 29, 13}},
	}

	for _, example := range examples {
		got := fixUpdate(example.update, rules)
		if !reflect.DeepEqual(got, example.expected) {
			t.Errorf("%v: expected %v, got %v", example.update, example.expected, got)
		}
	}
}
