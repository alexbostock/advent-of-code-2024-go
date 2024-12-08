package main

import (
	"fmt"
	"os"

	"adventofcode2024/puzzle1"
	"adventofcode2024/puzzle2"
	"adventofcode2024/puzzle3"
	"adventofcode2024/puzzle4"
	"adventofcode2024/puzzle5"
	"adventofcode2024/puzzle6"
	"adventofcode2024/puzzle7"
	"adventofcode2024/puzzle8"
)

func main() {
	solvePuzzle1()
	solvePuzzle2()
	solvePuzzle3()
	solvePuzzle4()
	solvePuzzle5()
	solvePuzzle6()
	solvePuzzle7()
	solvePuzzle8()
}

func solvePuzzle1() {
	fmt.Println("Puzzle 1")
	input, err := os.Open("./input/1.txt")
	if err != nil {
		panic(err)
	}
	left, right, err := puzzle1.ParseInput1(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(puzzle1.DiffLocations(left, right))
	fmt.Println(puzzle1.CalculateSimilarityScore(left, right))
}

func solvePuzzle2() {
	fmt.Println("Puzzle 2")
	input, err := os.Open("./input/2.txt")
	if err != nil {
		panic(err)
	}
	reports, err := puzzle2.ParseInput2(input)
	fmt.Println(puzzle2.CountSafeReports(reports))
	fmt.Println(puzzle2.CountSafeReportsWithProblemDampener(reports))
}

func solvePuzzle3() {
	fmt.Println("Puzzle 3")
	input, err := os.Open("./input/3.txt")
	if err != nil {
		panic(err)
	}
	commands := puzzle3.ParseInput3(input)
	fmt.Println(puzzle3.SumMuls(commands, false))
	fmt.Println(puzzle3.SumMuls(commands, true))
}

func solvePuzzle4() {
	fmt.Println("Puzzle 4")
	input, err := os.Open("./input/4.txt")
	if err != nil {
		panic(err)
	}
	wordSearch := puzzle4.ParseInput4(input)
	fmt.Println(puzzle4.CountXMASInWordSearch(wordSearch))
	fmt.Println(puzzle4.CountCrossMASInWordSearch(wordSearch))
}

func solvePuzzle5() {
	fmt.Println("Puzzle 5")
	inputData, err := os.Open("./input/5.txt")
	if err != nil {
		panic(err)
	}
	input := puzzle5.ParseInput5(inputData)
	fmt.Println(puzzle5.SumMiddlePagesOfCorrectlyOrderedUpdates(input))
	fmt.Println(puzzle5.SumMiddlePagesOfFixedUpdates(input))
}

func solvePuzzle6() {
	fmt.Println("Puzzle 6")
	input, err := os.Open("./input/6.txt")
	if err != nil {
		panic(err)
	}
	area := puzzle6.ParseInput6(input)
	fmt.Println(puzzle6.CountGuardPositionsVisited(area.Clone()))
	fmt.Println(puzzle6.CountPossibleNewObstaclesCausingLoops(area))
}

func solvePuzzle7() {
	fmt.Println("Puzzle 7")
	input, err := os.Open("./input/7.txt")
	if err != nil {
		panic(err)
	}
	equations := puzzle7.ParseInput7(input)
	fmt.Println(puzzle7.SumValidCalibrationEquations(equations, false))
	fmt.Println(puzzle7.SumValidCalibrationEquations(equations, true))
}

func solvePuzzle8() {
	fmt.Println("Puzzle 8")
	input, err := os.Open("./input/8.txt")
	if err != nil {
		panic(err)
	}
	area := puzzle8.ParseInput8(input)
	fmt.Println(puzzle8.CountDistinctAntinodes(area))
	fmt.Println(puzzle8.CountDistinctAntinodesWithResonantHarmonics(area))
}
