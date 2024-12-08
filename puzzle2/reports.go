package puzzle2

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

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
