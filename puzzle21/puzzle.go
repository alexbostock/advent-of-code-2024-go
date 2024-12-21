package puzzle21

import (
	"bufio"
	"fmt"
	"io"
	"unicode"
)

// final keypad: always need ([navigate to button in fewest key presses] [press a])*
// fastest way to navigate is always [(up or down)* (right or left)*] or [(right or left)* (up or down)*]

type Instruction int

const (
	A Instruction = iota
	Up
	Down
	Left
	Right
)

func (i Instruction) String() string {
	switch i {
	case A:
		return "A"
	case Up:
		return "^"
	case Down:
		return "v"
	case Left:
		return "<"
	case Right:
		return ">"
	}
	panic("unrecognised instruction")
}

type vector struct{ x, y int }

func ParseInput21(input io.Reader) (codes []string) {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		codes = append(codes, scanner.Text())
	}
	return
}

func TotalComplexityOfCodes(codes []string, numIntermediateRobots int) (total int) {
	for _, code := range codes {
		total += complexityOfCode(code, numIntermediateRobots)
	}
	return
}

func complexityOfCode(code string, numIntermediateRobots int) int {
	numericPart := 0
	for _, char := range []rune(code) {
		if unicode.IsDigit(char) {
			digit := int(char) - '0'
			numericPart = 10*numericPart + digit
		}
	}
	return len(shortestPathToInputCode(code, numIntermediateRobots)) * numericPart
}

func shortestPathToInputCode(code string, numIntermediateRobots int) []Instruction {
	var shortestCandidate []Instruction
	// for _, candidateNumeric := range candidateNumericKeypadSequencesToInputCode([]rune(code)) {
	// 	for _, candidateFirstDirectional := range candidateDirectionalKeypadSequencesToInputInstructions(candidateNumeric) {
	// 		for _, candidateSecondDirectional := range candidateDirectionalKeypadSequencesToInputInstructions(candidateFirstDirectional) {
	// 			if shortestCandidate == nil || len(candidateSecondDirectional) < len(shortestCandidate) {
	// 				shortestCandidate = candidateSecondDirectional
	// 			}
	// 		}
	// 	}
	// }
	// return shortestCandidate

	candidates := make(chan []Instruction, 1)
	go generateCandidates(candidates, code, numIntermediateRobots)
	for candidate := range candidates {
		if shortestCandidate == nil || len(candidate) < len(shortestCandidate) {
			shortestCandidate = candidate
		}
	}
	return shortestCandidate
}

func generateCandidates(output chan []Instruction, code string, numIntermediateRobots int) {
	if numIntermediateRobots == 0 {
		for _, candidateNumeric := range candidateNumericKeypadSequencesToInputCode([]rune(code)) {
			output <- candidateNumeric
		}
	} else {
		nestedCandidates := make(chan []Instruction, 1)
		go generateCandidates(nestedCandidates, code, numIntermediateRobots-1)
		for nestedCandidate := range nestedCandidates {
			for _, candidate := range candidateDirectionalKeypadSequencesToInputInstructions(nestedCandidate) {
				output <- candidate
			}
		}
	}
	close(output)
}

func candidateNumericKeypadSequencesToInputCode(code []rune) [][]Instruction {
	startPosition := positionOfKeyOnNumericKeypad('A')
	positions := make([]vector, len(code))
	for index, char := range code {
		positions[index] = positionOfKeyOnNumericKeypad(char)
	}
	moveVectors := make([]vector, len(code))
	for index, position := range positions {
		prev := startPosition
		if index > 0 {
			prev = positions[index-1]
		}
		moveVectors[index] = vector{position.x - prev.x, position.y - prev.y}
	}

	candidates := allCandidateKeypadSequences(moveVectors)
	var filteredCandidates [][]Instruction
	forbiddenPosition := vector{0, 3}
	for _, candidate := range candidates {
		if isValidSequence(positionOfKeyOnNumericKeypad('A'), candidate, forbiddenPosition) {
			filteredCandidates = append(filteredCandidates, candidate)
		}
	}
	return filteredCandidates
}

func candidateDirectionalKeypadSequencesToInputInstructions(outputInstructions []Instruction) [][]Instruction {
	startPosition := positionOfKeyOnDirectionalKeypad(A)
	positions := make([]vector, len(outputInstructions))
	for index, instruction := range outputInstructions {
		positions[index] = positionOfKeyOnDirectionalKeypad(instruction)
	}
	moveVectors := make([]vector, len(outputInstructions))
	for index, position := range positions {
		prev := startPosition
		if index > 0 {
			prev = positions[index-1]
		}
		moveVectors[index] = vector{position.x - prev.x, position.y - prev.y}
	}

	candidates := allCandidateKeypadSequences(moveVectors)
	var filteredCandidates [][]Instruction
	forbiddenPosition := vector{0, 0}
	for _, candidate := range candidates {
		if isValidSequence(positionOfKeyOnDirectionalKeypad(A), candidate, forbiddenPosition) {
			filteredCandidates = append(filteredCandidates, candidate)
		}
	}
	return filteredCandidates
}

func allCandidateKeypadSequences(moveVectors []vector) [][]Instruction {
	if len(moveVectors) == 0 {
		return [][]Instruction{nil}
	}
	var candidateSequencesForCurrentNumber [][]Instruction
	nextMove := moveVectors[0]

	var xSequence []Instruction
	for i := 0; i < abs(nextMove.x); i++ {
		if nextMove.x > 0 {
			xSequence = append(xSequence, Right)
		} else {
			xSequence = append(xSequence, Left)
		}
	}
	var ySequence []Instruction
	for i := 0; i < abs(nextMove.y); i++ {
		if nextMove.y > 0 {
			ySequence = append(ySequence, Down)
		} else {
			ySequence = append(ySequence, Up)
		}
	}

	if len(xSequence) > 0 && len(ySequence) > 0 {
		candidateSequencesForCurrentNumber = [][]Instruction{
			append(xSequence, ySequence...),
			append(ySequence, xSequence...),
		}
	} else if len(xSequence) > 0 {
		candidateSequencesForCurrentNumber = [][]Instruction{xSequence}
	} else {
		candidateSequencesForCurrentNumber = [][]Instruction{ySequence}
	}

	for index, sequence := range candidateSequencesForCurrentNumber {
		candidateSequencesForCurrentNumber[index] = append(sequence, A)
	}

	onwardSequences := allCandidateKeypadSequences(moveVectors[1:])

	var candidateSequences [][]Instruction
	for _, currentNumberSequence := range candidateSequencesForCurrentNumber {
		for _, onwardSequence := range onwardSequences {
			candidateSequence := append(currentNumberSequence, onwardSequence...)
			candidateSequences = append(candidateSequences, candidateSequence)
		}
	}
	return candidateSequences
}

func positionOfKeyOnNumericKeypad(key rune) vector {
	switch key {
	case '7':
		return vector{0, 0}
	case '8':
		return vector{1, 0}
	case '9':
		return vector{2, 0}
	case '4':
		return vector{0, 1}
	case '5':
		return vector{1, 1}
	case '6':
		return vector{2, 1}
	case '1':
		return vector{0, 2}
	case '2':
		return vector{1, 2}
	case '3':
		return vector{2, 2}
	case '0':
		return vector{1, 3}
	case 'A':
		return vector{2, 3}
	}
	panic(fmt.Sprintf("unexpected numeric key: %#v", key))
}

func positionOfKeyOnDirectionalKeypad(key Instruction) vector {
	switch key {
	case Up:
		return vector{1, 0}
	case A:
		return vector{2, 0}
	case Left:
		return vector{0, 1}
	case Down:
		return vector{1, 1}
	case Right:
		return vector{2, 1}
	}
	panic(fmt.Sprintf("unexpected directional key: %#v", key))
}

func isValidSequence(startPosition vector, instructions []Instruction, forbiddenPosition vector) bool {
	position := startPosition
	for _, instruction := range instructions {
		switch instruction {
		case Up:
			position.y--
		case Down:
			position.y++
		case Left:
			position.x--
		case Right:
			position.x++
		}
		if position == forbiddenPosition {
			return false
		}
	}
	return true
}

func abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}
