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
	puzzle6()
	puzzle7()
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

func puzzle6() {
	fmt.Println("Puzzle 6")
	input, err := os.Open("./input/6.txt")
	if err != nil {
		panic(err)
	}
	area := ParseInput6(input)
	areaCopy := &Map{
		height:    area.height,
		width:     area.width,
		obstacles: area.obstacles,
		guardPosition: Coords{
			area.guardPosition.i,
			area.guardPosition.j,
		},
		guardDirection: area.guardDirection,
	}
	fmt.Println(CountGuardPositionsVisited(area))
	fmt.Println(CountPossibleNewObstaclesCausingLoops(areaCopy))
}

func puzzle7() {
	fmt.Println("Puzzle 7")
	input, err := os.Open("./input/7.txt")
	if err != nil {
		panic(err)
	}
	equations := ParseInput7(input)
	fmt.Println(SumValidCalibrationEquations(equations, false))
	fmt.Println(SumValidCalibrationEquations(equations, true))
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

// Puzzle 6
type Map struct {
	height         int
	width          int
	obstacles      map[Coords]bool
	guardPosition  Coords
	guardDirection Direction
}
type Coords struct {
	i int
	j int
}
type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

func ParseInput6(data io.Reader) *Map {
	scanner := bufio.NewScanner(data)

	area := Map{
		obstacles: make(map[Coords]bool),
	}

	currentLineNum := -1
	for scanner.Scan() {
		currentLineNum++
		area.height = currentLineNum + 1
		line := scanner.Text()
		cells := strings.Split(line, "")
		for indexInLine, cell := range cells {
			area.width = indexInLine + 1
			if cell == "#" {
				area.obstacles[Coords{currentLineNum, indexInLine}] = true
			} else if cell == "^" {
				area.guardPosition = Coords{currentLineNum, indexInLine}
				area.guardDirection = Up
			}
		}
	}

	return &area
}

func CountGuardPositionsVisited(area *Map) int {
	positionsVisited := make(map[Coords]bool)
	for area.guardPosition.i >= 0 && area.guardPosition.i < area.height && area.guardPosition.j >= 0 && area.guardPosition.j < area.width {
		positionsVisited[area.guardPosition] = true
		nextPosition := oneStepForward(area.guardPosition, area.guardDirection)
		if area.obstacles[nextPosition] {
			area.guardDirection = turnRight(area.guardDirection)
		} else {
			area.guardPosition = nextPosition
		}
	}
	return len(positionsVisited)
}

func oneStepForward(position Coords, direction Direction) Coords {
	switch direction {
	case Up:
		return Coords{position.i - 1, position.j}
	case Down:
		return Coords{position.i + 1, position.j}
	case Right:
		return Coords{position.i, position.j + 1}
	case Left:
		return Coords{position.i, position.j - 1}
	}
	panic("Unexpected direction")
}

func turnRight(direction Direction) Direction {
	switch direction {
	case Up:
		return Right
	case Right:
		return Down
	case Down:
		return Left
	case Left:
		return Up
	}
	panic("Unexpected direction")
}

func CountPossibleNewObstaclesCausingLoops(area *Map) int {
	initialGuardPosition := Coords{area.guardPosition.i, area.guardPosition.j}
	initialGuardDirection := area.guardDirection
	count := 0
	for i := 0; i < area.height; i++ {
		for j := 0; j < area.width; j++ {
			pos := Coords{i, j}
			if area.obstacles[pos] {
				continue
			}
			if pos == initialGuardPosition {
				continue
			}

			area.obstacles[pos] = true
			if hasGuardLoop(area) {
				count++
			}

			delete(area.obstacles, pos)
			area.guardPosition.i = initialGuardPosition.i
			area.guardPosition.j = initialGuardPosition.j
			area.guardDirection = initialGuardDirection
		}
	}
	return count
}

type PositionAndDirection struct {
	position  Coords
	direction Direction
}

func hasGuardLoop(area *Map) bool {
	visitedPositions := make(map[PositionAndDirection]bool)
	for area.guardPosition.i >= 0 && area.guardPosition.i < area.height && area.guardPosition.j >= 0 && area.guardPosition.j < area.width {
		guardState := PositionAndDirection{area.guardPosition, area.guardDirection}
		if visitedPositions[guardState] {
			return true
		}
		visitedPositions[guardState] = true
		nextPosition := oneStepForward(area.guardPosition, area.guardDirection)
		if area.obstacles[nextPosition] {
			area.guardDirection = turnRight(area.guardDirection)
		} else {
			area.guardPosition = nextPosition
		}
	}
	return false
}

// Puzzle 7
type CalibrationEquation struct {
	lhs int
	rhs []int
}

func ParseInput7(data io.Reader) []CalibrationEquation {
	var equations []CalibrationEquation
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()
		sides := strings.Split(line, ": ")
		lhs, err := strconv.Atoi(sides[0])
		if err != nil {
			panic(err)
		}
		rhsTokens := strings.Split(sides[1], " ")
		var rhs []int
		for _, token := range rhsTokens {
			num, err := strconv.Atoi(token)
			if err != nil {
				panic(err)
			}
			rhs = append(rhs, num)
		}
		equations = append(equations, CalibrationEquation{lhs, rhs})
	}
	return equations
}

func SumValidCalibrationEquations(equations []CalibrationEquation, withConcatenation bool) int {
	sum := 0
	for _, equation := range equations {
		if canSolveCalibrationEquation(equation, withConcatenation) {
			sum += equation.lhs
		}
	}
	return sum
}

func canSolveCalibrationEquation(equation CalibrationEquation, withConcatenation bool) bool {
	if len(equation.rhs) == 1 {
		return equation.lhs == equation.rhs[0]
	}
	lastTermInRHS := equation.rhs[len(equation.rhs)-1]
	allOtherTerms := equation.rhs[:len(equation.rhs)-1]
	concatOperand, canConcat := reverseConcat(equation.lhs, lastTermInRHS)

	return equation.lhs%lastTermInRHS == 0 && canSolveCalibrationEquation(CalibrationEquation{equation.lhs / lastTermInRHS, allOtherTerms}, withConcatenation) ||
		canSolveCalibrationEquation(CalibrationEquation{equation.lhs - lastTermInRHS, allOtherTerms}, withConcatenation) ||
		withConcatenation && canConcat && canSolveCalibrationEquation(CalibrationEquation{concatOperand, allOtherTerms}, withConcatenation)
}

func reverseConcat(term, suffix int) (prefix int, isValid bool) {
	if term <= suffix {
		return 0, false
	}
	factor := smallestMultipleOfTenGreaterThan(suffix, 1)
	if term%factor == suffix {
		return (term - suffix) / factor, true
	} else {
		return 0, false
	}
}

func smallestMultipleOfTenGreaterThan(num, multipleOfTen int) int {
	if num < multipleOfTen {
		return multipleOfTen
	} else {
		return smallestMultipleOfTenGreaterThan(num, multipleOfTen*10)
	}
}
