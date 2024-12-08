package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	puzzle1()
	puzzle2()
	puzzle3()
	puzzle4()
	puzzle5()
}

func puzzle1() {
	fmt.Println("Puzzle 1")
	input, err := os.Open("./input/1.txt")
	if err != nil {
		panic(err)
	}
	left, right, err := ParseInput1(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(DiffLocations(left, right))
	fmt.Println(CalculateSimilarityScore(left, right))
}

func puzzle2() {
	fmt.Println("Puzzle 2")
	input, err := os.Open("./input/2.txt")
	if err != nil {
		panic(err)
	}
	reports, err := ParseInput2(input)
	fmt.Println(CountSafeReports(reports))
	fmt.Println(CountSafeReportsWithProblemDampener(reports))
}

func puzzle3() {
	fmt.Println("Puzzle 3")
	input, err := os.Open("./input/3.txt")
	if err != nil {
		panic(err)
	}
	commands := parseInput3(input)
	fmt.Println(sumMuls(commands, false))
	fmt.Println(sumMuls(commands, true))
}

func puzzle4() {
	fmt.Println("Puzzle 4")
	input, err := os.Open("./input/4.txt")
	if err != nil {
		panic(err)
	}
	wordSearch := ParseInput4(input)
	fmt.Println(CountXMASInWordSearch(wordSearch))
	fmt.Println(CountCrossMASInWordSearch(wordSearch))
}

func puzzle5() {
	fmt.Println("Puzzle 5")
	inputData, err := os.Open("./input/5.txt")
	if err != nil {
		panic(err)
	}
	input := ParseInput5(inputData)
	fmt.Println(SumMiddlePagesOfCorrectlyOrderedUpdates(input))
	fmt.Println(SumMiddlePagesOfFixedUpdates(input))
}

// Puzzle 1
func ParseInput1(data io.Reader) (left, right []int, err error) {
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		if len(tokens) < 2 {
			return nil, nil, errors.New("Malformatted line " + line)
		}

		leftVal, err := strconv.Atoi(tokens[0])
		if err != nil {
			return nil, nil, err
		}
		rightVal, err := strconv.Atoi(tokens[len(tokens)-1])
		if err != nil {
			return nil, nil, err
		}

		left = append(left, leftVal)
		right = append(right, rightVal)
	}

	return
}

func DiffLocations(left, right []int) int {
	slices.Sort(left)
	slices.Sort(right)
	distance := 0
	for index, leftVal := range left {
		rightVal := right[index]
		if leftVal > rightVal {
			distance += leftVal - rightVal
		} else {
			distance += rightVal - leftVal
		}
	}
	return distance
}

func CalculateSimilarityScore(left, right []int) int {
	slices.Sort(left)
	slices.Sort(right)
	score := 0
	rightCounts := make(map[int]int)
	for _, val := range right {
		rightCounts[val]++
	}

	for _, val := range left {
		score += val * rightCounts[val]
	}
	return score
}

// Puzzle 2
func ParseInput2(data io.Reader) (reports [][]int, err error) {
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")

		report := make([]int, len(tokens))
		for i, token := range tokens {
			val, err := strconv.Atoi(token)
			if err != nil {
				return nil, err
			}
			report[i] = val
		}
		reports = append(reports, report)
	}
	return
}

func CountSafeReports(
	reports [][]int) int {
	numSafe := 0
	for _, report := range reports {
		if IsSafe(report, -1) {
			numSafe++
		}
	}
	return numSafe
}

func IsSafe(report []int, dropIndex int) bool {
	hasIncrease := false
	hasDecrease := false
	prev := report[0]
	if dropIndex == 0 {
		prev = report[1]
	}
	firstRow := true
	for index, val := range report {
		if index == dropIndex {
			continue
		}
		if firstRow {
			firstRow = false
			continue
		}
		if val > prev {
			hasIncrease = true
			diff := val - prev
			if diff < 1 || diff > 3 {
				return false
			}
		} else {
			hasDecrease = true
			diff := prev - val
			if diff < 1 || diff > 3 {
				return false
			}
		}
		if hasIncrease && hasDecrease {
			return false
		}
		prev = val
	}
	return true
}

func CountSafeReportsWithProblemDampener(reports [][]int) int {
	numSafe := 0
	for _, report := range reports {
		if IsSafeWithProblemDampener(report) {
			numSafe++
		}
	}
	return numSafe
}

func IsSafeWithProblemDampener(report []int) bool {
	for i := -1; i < len(report); i++ {
		if IsSafe(report, i) {
			return true
		}
	}
	return false
}

// Puzzle 3
type Command struct {
	instruction Instruction
	op1         int
	op2         int
}

type Instruction int

const (
	NOP  Instruction = 0
	Do               = 1
	Dont             = 2
	Mul              = 3
)

func parseInput3(input io.Reader) []*Command {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, input)
	if err != nil {
		panic(err)
	}
	return parseInput3Str(buf.String())
}

func parseInput3Str(input string) []*Command {
	if len(input) == 0 {
		return nil
	}
	var answer []*Command
	for offset := 0; offset < len(input); {
		if input[offset] == 'm' {
			mul, newOffset := parseMul(input, offset)
			if mul != nil {
				answer = append(answer, mul)
			}
			if newOffset == offset {
				offset++
			} else {
				offset = newOffset
			}
		} else if input[offset] == 'd' {
			doOrDont, newOffset := parseDoOrDont(input, offset)
			if doOrDont == Do {
				answer = append(answer, &Command{Do, 0, 0})
			} else if doOrDont == Dont {
				answer = append(answer, &Command{Dont, 0, 0})
			}
			offset = newOffset
		} else {
			offset++
		}
	}
	return answer
}

func parseMul(input string, offset int) (mul *Command, newOffset int) {
	if len(input) < offset+6 {
		return nil, offset + 1
	}
	if input[offset:offset+4] != "mul(" {
		return nil, offset + 1
	}
	offset += 4
	op1, offsetAfterOp1 := parseInt(input, offset)
	if offsetAfterOp1 == offset {
		return nil, offset
	}
	offset = offsetAfterOp1
	if input[offset] != ',' {
		return nil, offset + 1
	}
	offset++
	op2, offsetAfterOp2 := parseInt(input, offset)
	if offsetAfterOp2 == offset {
		return nil, offset
	}
	offset = offsetAfterOp2
	if input[offset] != ')' {
		return nil, offset
	}
	return &Command{Mul, op1, op2}, offset + 1
}

func parseInt(input string, offset int) (num int, newOffset int) {
	newOffset = offset
	for i := 0; i < 3; i++ {
		if offset+i >= len(input) {
			return
		}
		digit, err := strconv.Atoi(string(input[offset+i]))
		if err != nil {
			// Not a digit
			return
		}
		num = 10*num + digit
		newOffset++
	}
	return
}

func parseDoOrDont(input string, offset int) (instruction Instruction, newOffset int) {
	if len(input) < offset+2 {
		return NOP, offset + 2
	}
	if input[offset:offset+2] != "do" {
		return NOP, offset + 2
	}
	offset += 2
	if len(input) >= offset+2 && input[offset:offset+2] == "()" {
		return Do, offset + 2
	}
	if len(input) >= offset+5 && input[offset:offset+5] == "n't()" {
		return Dont, offset + 2
	}
	return NOP, offset + 1
}

func sumMuls(commands []*Command, observeDosAndDonts bool) int {
	sum := 0
	mulEnabled := true
	for _, command := range commands {
		switch command.instruction {
		case Do:
			mulEnabled = true
		case Dont:
			mulEnabled = false
		case Mul:
			if !observeDosAndDonts || mulEnabled {
				sum += command.op1 * command.op2
			}
		}
	}
	return sum
}

// Puzzle 4
type WordSearch struct {
	chars  [][]rune
	height int
	width  int
}

func ParseInput4(data io.Reader) WordSearch {
	var chars [][]rune
	scanner := bufio.NewScanner(data)
	height := 0
	width := 0
	for scanner.Scan() {
		chars = append(chars, []rune(scanner.Text()))
		height++
		width = len(scanner.Text())
	}
	return WordSearch{
		chars, height, width,
	}
}

func CountXMASInWordSearch(wordSearch WordSearch) int {
	count := 0
	for i := 0; i < wordSearch.height; i++ {
		for j := 0; j < wordSearch.width; j++ {
			patterns := [][]struct {
				i int
				j int
			}{
				{{i, j}, {i + 1, j}, {i + 2, j}, {i + 3, j}},
				{{i, j}, {i - 1, j}, {i - 2, j}, {i - 3, j}},
				{{i, j}, {i, j + 1}, {i, j + 2}, {i, j + 3}},
				{{i, j}, {i, j - 1}, {i, j - 2}, {i, j - 3}},
				{{i, j}, {i + 1, j + 1}, {i + 2, j + 2}, {i + 3, j + 3}},
				{{i, j}, {i + 1, j - 1}, {i + 2, j - 2}, {i + 3, j - 3}},
				{{i, j}, {i - 1, j + 1}, {i - 2, j + 2}, {i - 3, j + 3}},
				{{i, j}, {i - 1, j - 1}, {i - 2, j - 2}, {i - 3, j - 3}},
			}
			for _, pattern := range patterns {
				if hasXMASInWordSearchAtPositions(wordSearch, pattern) {
					count++
				}
			}
		}
	}
	return count
}

func hasXMASInWordSearchAtPositions(wordSearch WordSearch, coords []struct {
	i int
	j int
}) bool {
	for pos, coord := range coords {
		if coord.i < 0 || coord.i >= wordSearch.height || coord.j < 0 || coord.j >= wordSearch.width {
			return false
		}
		expected := rune("XMAS"[pos])
		if wordSearch.chars[coord.i][coord.j] != expected {
			return false
		}
	}
	return true
}

func CountCrossMASInWordSearch(wordSearch WordSearch) int {
	count := 0
	for i := 0; i < wordSearch.height; i++ {
		for j := 0; j < wordSearch.width; j++ {
			if hasCrossMASInWordSearchCentredAt(wordSearch, i, j) {
				count++
			}
		}
	}
	return count
}

func hasCrossMASInWordSearchCentredAt(wordSearch WordSearch, i, j int) bool {
	if i-1 < 0 || i+1 >= wordSearch.height || j-1 < 0 || j+1 >= wordSearch.width {
		return false
	}
	topLeft := wordSearch.chars[i-1][j-1]
	topRight := wordSearch.chars[i-1][j+1]
	bottomLeft := wordSearch.chars[i+1][j-1]
	bottomRight := wordSearch.chars[i+1][j+1]
	centre := wordSearch.chars[i][j]
	diagonalOneOkay := topLeft == 'M' && bottomRight == 'S' || topLeft == 'S' && bottomRight == 'M'
	diagonalTwoOkay := topRight == 'M' && bottomLeft == 'S' || topRight == 'S' && bottomLeft == 'M'
	return centre == 'A' && diagonalOneOkay && diagonalTwoOkay
}

// Puzzle 5
type Puzzle5Input struct {
	rules   []pageOrderingRule
	updates [][]int
}
type pageOrderingRule struct {
	before int
	after  int
}

func ParseInput5(data io.Reader) Puzzle5Input {
	var rules []pageOrderingRule
	var updates [][]int

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if strings.Contains(line, "|") {
			tokens := strings.Split(line, "|")
			before, err := strconv.Atoi(tokens[0])
			if err != nil {
				panic(err)
			}
			after, err := strconv.Atoi(tokens[1])
			if err != nil {
				panic(err)
			}
			rules = append(rules, struct {
				before int
				after  int
			}{before, after})
		} else {
			tokens := strings.Split(line, ",")
			var update []int
			for _, token := range tokens {
				num, err := strconv.Atoi(token)
				if err != nil {
					panic(err)
				}
				update = append(update, num)
			}
			updates = append(updates, update)
		}
	}

	return Puzzle5Input{rules, updates}
}

func SumMiddlePagesOfCorrectlyOrderedUpdates(input Puzzle5Input) int {
	rules := loadPageOrderingRules(input.rules)
	sum := 0
	for _, update := range input.updates {
		if isValidUpdate(update, rules) {
			sum += middlePageInUpdate(update)
		}
	}
	return sum
}

type pageOrderingRules struct {
	rules        map[pageOrderingRule]bool
	byBeforePage map[int][]int
	byAfterPage  map[int][]int
}

func loadPageOrderingRules(rules []pageOrderingRule) pageOrderingRules {
	loadedRules := pageOrderingRules{
		rules:        make(map[pageOrderingRule]bool),
		byBeforePage: make(map[int][]int),
		byAfterPage:  make(map[int][]int),
	}
	for index, rule := range rules {
		loadedRules.rules[rule] = true
		loadedRules.byBeforePage[rule.before] = append(loadedRules.byBeforePage[rule.before], index)
		loadedRules.byAfterPage[rule.after] = append(loadedRules.byAfterPage[rule.after], index)
	}
	return loadedRules
}

func isValidUpdate(update []int, rules pageOrderingRules) bool {
	rulesClosed := make(map[int]bool)
	for _, page := range update {
		for _, ruleID := range rules.byAfterPage[page] {
			rulesClosed[ruleID] = true
		}
		for _, ruleID := range rules.byBeforePage[page] {
			if rulesClosed[ruleID] {
				return false
			}
		}
	}
	return true
}

func middlePageInUpdate(update []int) int {
	return update[(len(update)+1)/2-1]
}

func SumMiddlePagesOfFixedUpdates(input Puzzle5Input) int {
	rules := loadPageOrderingRules(input.rules)
	sum := 0
	for _, update := range input.updates {
		if !isValidUpdate(update, rules) {
			fixed := fixUpdate(update, rules)
			sum += middlePageInUpdate(fixed)
		}
	}
	return sum
}

func fixUpdate(update []int, rules pageOrderingRules) []int {
	fixed := slices.Clone(update)
	slices.SortStableFunc(fixed, func(page1, page2 int) int {
		if rules.rules[pageOrderingRule{page1, page2}] {
			return -1
		} else if rules.rules[pageOrderingRule{page2, page1}] {
			return 1
		}
		return 0
	})
	return fixed
}
