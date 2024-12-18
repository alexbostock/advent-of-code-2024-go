package puzzle17

import (
	"bufio"
	"io"
	"math"
	"strconv"
	"strings"
)

type State struct {
	pc      int
	regA    int
	regB    int
	regC    int
	program []int
}

func ParseInput17(data io.Reader) State {
	scanner := bufio.NewScanner(data)
	var parsed State
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if line[0:10] == "Register A" {
			val, err := strconv.Atoi(line[12:])
			if err != nil {
				panic(err)
			}
			parsed.regA = val
		}
		if line[0:10] == "Register B" {
			val, err := strconv.Atoi(line[12:])
			if err != nil {
				panic(err)
			}
			parsed.regB = val
		}
		if line[0:10] == "Register C" {
			val, err := strconv.Atoi(line[12:])
			if err != nil {
				panic(err)
			}
			parsed.regC = val
		}
		if line[0:9] == "Program: " {
			tokens := strings.Split(line[9:], ",")
			for _, token := range tokens {
				val, err := strconv.Atoi(token)
				if err != nil {
					panic(err)
				}
				parsed.program = append(parsed.program, val)
			}
		}
	}
	return parsed
}

func (state State) Clone() State {
	return State{
		pc:      state.pc,
		regA:    state.regA,
		regB:    state.regB,
		regC:    state.regC,
		program: state.program,
	}
}

func ExecuteProgram(state State) (output []int, outputStr string) {
	var outputStrs []string

	for state.pc = 0; state.pc < len(state.program); state.pc += 2 {
		opcode := state.program[state.pc]
		operand := state.program[state.pc+1]
		comboOperand := getComboOperand(state, operand)

		switch opcode {
		case 0:
			state.regA = state.regA / int(math.Pow(2, float64(comboOperand)))
		case 1:
			state.regB = state.regB ^ operand
		case 2:
			state.regB = comboOperand % 8
		case 3:
			if state.regA != 0 {
				state.pc = operand - 2
			}
		case 4:
			state.regB = state.regB ^ state.regC
		case 5:
			valToOutput := comboOperand % 8
			output = append(output, valToOutput)
			outputStrs = append(outputStrs, strconv.Itoa(valToOutput))
		case 6:
			state.regB = state.regA / int(math.Pow(2, float64(comboOperand)))
		case 7:
			state.regC = state.regA / int(math.Pow(2, float64(comboOperand)))
		}
	}

	return output, strings.Join(outputStrs, ",")
}

func getComboOperand(state State, operand int) int {
	if operand <= 3 {
		return operand
	}
	switch operand {
	case 4:
		return state.regA
	case 5:
		return state.regB
	case 6:
		return state.regC
	}

	return 0 // combo operand should never be used in this case
}

func FindRegAValueWhichMakesQuine(initialState State) int {
	// Observations:
	// Changing regA by 1 changes first output
	// Changing regA by 8 changes second output
	// Changing regA by 64 changes third output
	// etc.
	// (in each case, there is a chance of no change; changing regA by 8**n is necessary, not sufficient)

	regA := 0
	for outputOffset := len(initialState.program) - 1; outputOffset >= 0; outputOffset-- {
		interval := int(math.Pow(8, float64(outputOffset)))
		for {
			state := initialState.Clone()
			state.regA = regA
			output, _ := ExecuteProgram(state)
			if len(output) == len(initialState.program) && output[outputOffset] == initialState.program[outputOffset] {
				break
			}
			regA += interval
		}
	}
	return regA
}
