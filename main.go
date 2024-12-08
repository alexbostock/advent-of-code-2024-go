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
