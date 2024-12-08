package main

import (
	"reflect"
	"slices"
	"strings"
	"testing"
)

var leftExample = []int{3, 4, 2, 1, 3, 3}
var rightExample = []int{4, 3, 5, 3, 9, 3}

func TestParseInput1(t *testing.T) {
	left, right, err := ParseInput1(strings.NewReader(`3   4
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

func TestParseInput3(t *testing.T) {
	examples := []struct {
		input    string
		expected []*Command
	}{
		{"", nil},
		{"mul(1,1)", []*Command{{Mul, 1, 1}}},
		{"mul(131,45)", []*Command{{Mul, 131, 45}}},
		{"mul(1234,7)", nil},
		{" mul(2,3)", []*Command{{Mul, 2, 3}}},
		{" mul(2, 3)", nil},
		{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))", []*Command{{Mul, 2, 4}, {Mul, 5, 5}, {Mul, 11, 8}, {Mul, 8, 5}}},
	}

	for _, example := range examples {
		got := parseInput3Str(example.input)
		gotExpected := len(got) == 0 && len(example.expected) == 0 || reflect.DeepEqual(got, example.expected)
		if !gotExpected {
			t.Errorf("%v: expected %v, got %v", example.input, example.expected, got)
		}
	}
}

func TestSumMuls(t *testing.T) {
	got := sumMuls([]*Command{{Mul, 2, 4}, {Mul, 5, 5}, {Mul, 11, 8}, {Mul, 8, 5}}, false)
	expected := 161
	if got != expected {
		t.Errorf("got %v, expected %v", got, expected)
	}
}

func TestSumMulsWithDoAndDont(t *testing.T) {
	commands := parseInput3Str("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))")
	got := sumMuls(commands, true)
	expected := 48
	if got != expected {
		t.Errorf("got %v, expected %v", got, expected)
	}
}

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

func TestParseInput6(t *testing.T) {
	got := ParseInput6(strings.NewReader(`....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
`))

	expected := &Map{
		height: 10,
		width:  10,
		obstacles: map[Coords]bool{
			{0, 4}: true,
			{1, 9}: true,
			{3, 2}: true,
			{4, 7}: true,
			{6, 1}: true,
			{7, 8}: true,
			{8, 0}: true,
			{9, 6}: true,
		},
		guardPosition:  Coords{6, 4},
		guardDirection: Up,
	}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestCountGuardPositionsVisited(t *testing.T) {
	area := ParseInput6(strings.NewReader(`....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
`))
	got := CountGuardPositionsVisited(area)
	if got != 41 {
		t.Errorf("expected 41, got %v", got)
	}
}

func TestCountPossibleNewObstaclesCausingLoops(t *testing.T) {
	area := ParseInput6(strings.NewReader(`....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
`))

	got := CountPossibleNewObstaclesCausingLoops(area)
	if got != 6 {
		t.Errorf("expected 6, got %v", got)
	}
}

func TestHasGuardLoop(t *testing.T) {
	area := ParseInput6(strings.NewReader(`....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
`))
	got := hasGuardLoop(area)
	if got {
		t.Error("expected false, got true")
	}

	area = ParseInput6(strings.NewReader(`....#.....
.........#
..........
..#.......
.......#..
..........
.#.#^.....
........#.
#.........
......#...
`))

	got = hasGuardLoop(area)
	if !got {
		t.Error("expected true, got false")
	}
}

func TestParseInput7(t *testing.T) {
	got := ParseInput7(strings.NewReader(`190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
`))

	expected := []CalibrationEquation{
		{190, []int{10, 19}},
		{3267, []int{81, 40, 27}},
		{83, []int{17, 5}},
		{156, []int{15, 6}},
		{7290, []int{6, 8, 6, 15}},
		{161011, []int{16, 10, 13}},
		{192, []int{17, 8, 14}},
		{21037, []int{9, 7, 18, 13}},
		{292, []int{11, 6, 16, 20}},
	}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestSumValidCalibrationEquations(t *testing.T) {
	equations := []CalibrationEquation{
		{190, []int{10, 19}},
		{3267, []int{81, 40, 27}},
		{83, []int{17, 5}},
		{156, []int{15, 6}},
		{7290, []int{6, 8, 6, 15}},
		{161011, []int{16, 10, 13}},
		{192, []int{17, 8, 14}},
		{292, []int{11, 6, 16, 20}},
	}

	got := SumValidCalibrationEquations(equations, false)
	if got != 3749 {
		t.Errorf("expected 3749, got %v", got)
	}
}

func TestCanSolveCalibrationEquation(t *testing.T) {
	examples := []struct {
		equation CalibrationEquation
		expected bool
	}{
		{CalibrationEquation{190, []int{10, 19}}, true},
		{CalibrationEquation{3267, []int{81, 40, 27}}, true},
		{CalibrationEquation{83, []int{17, 5}}, false},
		{CalibrationEquation{156, []int{15, 6}}, false},
		{CalibrationEquation{7290, []int{6, 8, 6, 15}}, false},
		{CalibrationEquation{161011, []int{16, 10, 13}}, false},
		{CalibrationEquation{192, []int{17, 8, 14}}, false},
		{CalibrationEquation{292, []int{11, 6, 16, 20}}, true},
	}

	for _, example := range examples {
		got := canSolveCalibrationEquation(example.equation, false)
		if got != example.expected {
			t.Errorf("%v: expected %v, got %v", example.equation, example.expected, got)
		}
	}
}

func TestCanSolveCalibrationEquationWithConcatenation(t *testing.T) {
	examples := []struct {
		equation CalibrationEquation
		expected bool
	}{
		{CalibrationEquation{190, []int{10, 19}}, true},
		{CalibrationEquation{3267, []int{81, 40, 27}}, true},
		{CalibrationEquation{83, []int{17, 5}}, false},
		{CalibrationEquation{156, []int{15, 6}}, true},

		{CalibrationEquation{6, []int{6}}, true},
		{CalibrationEquation{48, []int{6, 8}}, true},
		{CalibrationEquation{486, []int{6, 8, 6}}, true},
		{CalibrationEquation{7290, []int{6, 8, 6, 15}}, true},

		{CalibrationEquation{161011, []int{16, 10, 13}}, false},
		{CalibrationEquation{192, []int{17, 8, 14}}, true},
		{CalibrationEquation{292, []int{11, 6, 16, 20}}, true},
	}

	for _, example := range examples {
		got := canSolveCalibrationEquation(example.equation, true)
		if got != example.expected {
			t.Errorf("%v: expected %v, got %v", example.equation, example.expected, got)
		}
	}
}

func TestParseInput8(t *testing.T) {
	got := ParseInput8(strings.NewReader(`............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............
`))

	expected := Puzzle8Input{
		height: 12,
		width:  12,
		antennas: map[rune][]Coords{
			'0': {{1, 8}, {2, 5}, {3, 7}, {4, 4}},
			'A': {{5, 6}, {8, 8}, {9, 9}},
		},
	}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestCountDistinctAntinodes(t *testing.T) {
	got := CountDistinctAntinodes(Puzzle8Input{
		height: 12,
		width:  12,
		antennas: map[rune][]Coords{
			'0': {{1, 8}, {2, 5}, {3, 7}, {4, 4}},
			'A': {{5, 6}, {8, 8}, {9, 9}},
		},
	})

	if got != 14 {
		t.Errorf("expected 14, got %v", got)
	}
}

func TestCountDistinctAntinodesWithResonantHarmonics(t *testing.T) {
	got := CountDistinctAntinodesWithResonantHarmonics(Puzzle8Input{
		height: 12,
		width:  12,
		antennas: map[rune][]Coords{
			'0': {{1, 8}, {2, 5}, {3, 7}, {4, 4}},
			'A': {{5, 6}, {8, 8}, {9, 9}},
		},
	})

	if got != 34 {
		t.Errorf("expected 34, got %v", got)
	}
}
