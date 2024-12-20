package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"adventofcode2024/puzzle1"
	"adventofcode2024/puzzle10"
	"adventofcode2024/puzzle11"
	"adventofcode2024/puzzle12"
	"adventofcode2024/puzzle13"
	"adventofcode2024/puzzle14"
	"adventofcode2024/puzzle15"
	"adventofcode2024/puzzle16"
	"adventofcode2024/puzzle17"
	"adventofcode2024/puzzle18"
	"adventofcode2024/puzzle19"
	"adventofcode2024/puzzle2"
	"adventofcode2024/puzzle20"
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
		solve     func(input io.ReadSeeker)
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
		{10, p10},
		{11, p11},
		{12, p12},
		{13, p13},
		{14, p14},
		{15, p15},
		{16, p16},
		{17, p17},
		{18, p18},
		{19, p19},
		{20, p20},
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

func solvePuzzle(puzzleNum int, solve func(input io.ReadSeeker)) {
	fmt.Printf("Puzzle %v\n", puzzleNum)
	input, err := os.Open(fmt.Sprintf("./input/%v.txt", puzzleNum))
	if err != nil {
		panic(err)
	}
	solve(input)
	fmt.Println()
}

func p1(input io.ReadSeeker) {
	left, right, err := puzzle1.ParseInput1(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(puzzle1.DiffLocations(left, right))
	fmt.Println(puzzle1.CalculateSimilarityScore(left, right))
}

func p2(input io.ReadSeeker) {
	reports, err := puzzle2.ParseInput2(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(puzzle2.CountSafeReports(reports))
	fmt.Println(puzzle2.CountSafeReportsWithProblemDampener(reports))
}

func p3(input io.ReadSeeker) {
	commands := puzzle3.ParseInput3(input)
	fmt.Println(puzzle3.SumMuls(commands, false))
	fmt.Println(puzzle3.SumMuls(commands, true))
}

func p4(input io.ReadSeeker) {
	wordSearch := puzzle4.ParseInput4(input)
	fmt.Println(puzzle4.CountXMASInWordSearch(wordSearch))
	fmt.Println(puzzle4.CountCrossMASInWordSearch(wordSearch))
}

func p5(inputData io.ReadSeeker) {
	input := puzzle5.ParseInput5(inputData)
	fmt.Println(puzzle5.SumMiddlePagesOfCorrectlyOrderedUpdates(input))
	fmt.Println(puzzle5.SumMiddlePagesOfFixedUpdates(input))
}

func p6(input io.ReadSeeker) {
	area := puzzle6.ParseInput6(input)
	fmt.Println(puzzle6.CountGuardPositionsVisited(area.Clone()))
	fmt.Println(puzzle6.CountPossibleNewObstaclesCausingLoops(area))
}

func p7(input io.ReadSeeker) {
	equations := puzzle7.ParseInput7(input)
	fmt.Println(puzzle7.SumValidCalibrationEquations(equations, false))
	fmt.Println(puzzle7.SumValidCalibrationEquations(equations, true))
}

func p8(input io.ReadSeeker) {
	area := puzzle8.ParseInput8(input)
	fmt.Println(puzzle8.CountDistinctAntinodes(area))
	fmt.Println(puzzle8.CountDistinctAntinodesWithResonantHarmonics(area))
}

func p9(input io.ReadSeeker) {
	fileSystem := puzzle9.ParseInput9(input)
	puzzle9.MoveBlocks(fileSystem.Blocks)
	fmt.Println(puzzle9.ComputeChecksum(fileSystem.Blocks))
	fmt.Println(puzzle9.ComputeChecksumFiles(puzzle9.MoveFiles(fileSystem.Files)))
}

func p10(input io.ReadSeeker) {
	area, err := puzzle10.ParseInput10(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(puzzle10.SumScoresOfTrailheads(area))
	fmt.Println(puzzle10.CountDistinctTrails(area))
}

func p11(input io.ReadSeeker) {
	stones, err := puzzle11.ParseInput11(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(puzzle11.CountAllStonesAfterNumBlinks(stones, 25))
	fmt.Println(puzzle11.CountAllStonesAfterNumBlinks(stones, 75))
}

func p12(input io.ReadSeeker) {
	area := puzzle12.ParseInput12(input)
	fmt.Println(puzzle12.CostFences(area))
	fmt.Println(puzzle12.CostFencesNumSides(area))
}

func p13(input io.ReadSeeker) {
	machines, err := puzzle13.ParseInput13(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(puzzle13.SearchMinimumTokensToWinAllPrizes(machines))
	fmt.Println(puzzle13.SearchMinimumTokensToWinAllPrizesWithPrizeError(machines))
}

func p14(input io.ReadSeeker) {
	robots := puzzle14.ParseInput14(input)
	after100Seconds := puzzle14.StateAfterSeconds(robots, 100, 101, 103)
	fmt.Println(puzzle14.SafetyFactor(after100Seconds, 101, 103))
	puzzle14.PrintEachState(robots, 50000, 101, 103)
}

func p15(input io.ReadSeeker) {
	warehouse := puzzle15.ParseInput15(input)
	warehouse.ExecuteInstructions()
	fmt.Println(warehouse.SumGPSCoordsOfBoxes())

	input.Seek(0, io.SeekStart)

	wideWarehouse := puzzle15.ParseInput15Wide(input)
	wideWarehouse.ExecuteInstructions()
	fmt.Println(wideWarehouse.SumGPSCoordsOfBoxes())
}

func p16(input io.ReadSeeker) {
	maze := puzzle16.ParseInput16(input)
	cost, positionsSeen := maze.ShortestPath()
	fmt.Println(cost)
	fmt.Println(len(positionsSeen))
}

func p17(input io.ReadSeeker) {
	state := puzzle17.ParseInput17(input)
	_, output := puzzle17.ExecuteProgram(state.Clone())
	fmt.Println(output)
	fmt.Println(puzzle17.FindRegAValueWhichMakesQuine(state))
}

func p18(input io.ReadSeeker) {
	bytes := puzzle18.ParseInput18(input)
	fmt.Println(puzzle18.ShortestPathAfterBytesFallen(bytes, 71, 71, 1024))
	fmt.Println(puzzle18.FindFirstByteObstructingExit(bytes, 71, 71))
}

func p19(input io.ReadSeeker) {
	towelsAndDesigns := puzzle19.ParseInput19(input)
	fmt.Println(puzzle19.CountPossibleDesigns(towelsAndDesigns))
	fmt.Println(puzzle19.SumPossibleWaysOfMakingDesigns(towelsAndDesigns))
}

func p20(input io.ReadSeeker) {
	maze := puzzle20.ParseInput20(input)
	fmt.Println(maze.NumCheatsSavingAtLeast(2, 100))
	fmt.Println(maze.NumCheatsSavingAtLeast(20, 100))
}
