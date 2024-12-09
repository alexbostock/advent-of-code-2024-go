package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"adventofcode2024/puzzle1"
	"adventofcode2024/puzzle2"
	"adventofcode2024/puzzle3"
	"adventofcode2024/puzzle4"
	"adventofcode2024/puzzle5"
	"adventofcode2024/puzzle6"
	"adventofcode2024/puzzle7"
	"adventofcode2024/puzzle8"
	"adventofcode2024/puzzle9"
)

func main() {
	puzzles := []struct {
		puzzleNum int
		solve     func(input io.Reader)
	}{
		{1, p1},
		{2, p2},
		{3, p3},
		{4, p4},
		{5, p5},
		{6, p6},
		{7, p7},
		{8, p8},
		{9, p9},
	}

	puzzleNum := ""
	if len(os.Args) > 1 {
		puzzleNum = os.Args[1]
	}

	for _, puzzle := range puzzles {
		if puzzleNum == "" || puzzleNum == strconv.Itoa(puzzle.puzzleNum) {
			solvePuzzle(puzzle.puzzleNum, puzzle.solve)
		}
	}
}

func solvePuzzle(puzzleNum int, solve func(input io.Reader)) {
	fmt.Printf("Puzzle %v\n", puzzleNum)
	input, err := os.Open(fmt.Sprintf("./input/%v.txt", puzzleNum))
	if err != nil {
		panic(err)
	}
	solve(input)
	fmt.Println()
}

func p1(input io.Reader) {
	left, right, err := puzzle1.ParseInput1(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(puzzle1.DiffLocations(left, right))
	fmt.Println(puzzle1.CalculateSimilarityScore(left, right))
}

func p2(input io.Reader) {
	reports, err := puzzle2.ParseInput2(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(puzzle2.CountSafeReports(reports))
	fmt.Println(puzzle2.CountSafeReportsWithProblemDampener(reports))
}

func p3(input io.Reader) {
	commands := puzzle3.ParseInput3(input)
	fmt.Println(puzzle3.SumMuls(commands, false))
	fmt.Println(puzzle3.SumMuls(commands, true))
}

func p4(input io.Reader) {
	wordSearch := puzzle4.ParseInput4(input)
	fmt.Println(puzzle4.CountXMASInWordSearch(wordSearch))
	fmt.Println(puzzle4.CountCrossMASInWordSearch(wordSearch))
}

func p5(inputData io.Reader) {
	input := puzzle5.ParseInput5(inputData)
	fmt.Println(puzzle5.SumMiddlePagesOfCorrectlyOrderedUpdates(input))
	fmt.Println(puzzle5.SumMiddlePagesOfFixedUpdates(input))
}

func p6(input io.Reader) {
	area := puzzle6.ParseInput6(input)
	fmt.Println(puzzle6.CountGuardPositionsVisited(area.Clone()))
	fmt.Println(puzzle6.CountPossibleNewObstaclesCausingLoops(area))
}

func p7(input io.Reader) {
	equations := puzzle7.ParseInput7(input)
	fmt.Println(puzzle7.SumValidCalibrationEquations(equations, false))
	fmt.Println(puzzle7.SumValidCalibrationEquations(equations, true))
}

func p8(input io.Reader) {
	area := puzzle8.ParseInput8(input)
	fmt.Println(puzzle8.CountDistinctAntinodes(area))
	fmt.Println(puzzle8.CountDistinctAntinodesWithResonantHarmonics(area))
}

func p9(input io.Reader) {
	fileSystem := puzzle9.ParseInput9(input)
	puzzle9.MoveBlocks(fileSystem.Blocks)
	fmt.Println(puzzle9.ComputeChecksum(fileSystem.Blocks))
	fmt.Println(puzzle9.ComputeChecksumFiles(puzzle9.MoveFiles(fileSystem.Files)))
}
