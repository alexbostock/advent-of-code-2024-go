package puzzle7

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

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
