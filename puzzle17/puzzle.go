package puzzle17

import (
	"bufio"
	"fmt"
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

func ExecuteProgram(state State, expectedOutput []int) (output []int, outputStr string, matchesExpected bool) {
	var outputStrs []string
	offsetInExpectedOutput := 0

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
			if expectedOutput != nil {
				if offsetInExpectedOutput >= len(expectedOutput) || valToOutput != expectedOutput[offsetInExpectedOutput] {
					matchesExpected = false
					return
				}
				offsetInExpectedOutput++
			}
			output = append(output, valToOutput)
			outputStrs = append(outputStrs, strconv.Itoa(valToOutput))
		case 6:
			state.regB = state.regA / int(math.Pow(2, float64(comboOperand)))
		case 7:
			state.regC = state.regA / int(math.Pow(2, float64(comboOperand)))
		}
	}

	return output, strings.Join(outputStrs, ","), expectedOutput == nil || len(output) == len(expectedOutput)
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

type Operation struct {
	operator          Operator
	operand1          int
	operand2          int
	operand1Recursive *Operation
	operand2Recursive *Operation
}
type Operator int

const (
	LITERAL          = 0
	XOR     Operator = iota
	MODULO
	POWER
	DIVIDE
	INPUT_REG_A
	EQUAL
	NOT_EQUAL
)

func (operation *Operation) String() string {
	var operand1Str string
	if operation.operand1Recursive == nil {
		operand1Str = strconv.Itoa(operation.operand1)
	} else {
		operand1Str = operation.operand1Recursive.String()
	}
	var operand2Str string
	if operation.operand2Recursive == nil {
		operand2Str = strconv.Itoa(operation.operand2)
	} else {
		operand2Str = operation.operand2Recursive.String()
	}

	switch operation.operator {
	case LITERAL:
		return operand1Str
	case XOR:
		return fmt.Sprintf("(%v ^ %v)", operand1Str, operand2Str)
	case MODULO:
		return fmt.Sprintf("(%v %% %v)", operand1Str, operand2Str)
	case POWER:
		return fmt.Sprintf("(%v ** %v)", operand1Str, operand2Str)
	case DIVIDE:
		return fmt.Sprintf("(%v / %v)", operand1Str, operand2Str)
	case INPUT_REG_A:
		return "R"
	case EQUAL:
		return fmt.Sprintf("(%v == %v)", operand1Str, operand2Str)
	case NOT_EQUAL:
		return fmt.Sprintf("(%v != %v)", operand1Str, operand2Str)
	}
	return ""
}

func (operation *Operation) simplifyConstants() *Operation {
	if operation.operand1Recursive != nil {
		operation.operand1Recursive = operation.operand1Recursive.simplifyConstants()
	}
	if operation.operand2Recursive != nil {
		operation.operand2Recursive = operation.operand2Recursive.simplifyConstants()
	}

	if operation.operand1Recursive != nil && operation.operand1Recursive.operator == LITERAL {
		operation.operand1 = operation.operand1Recursive.operand1
		operation.operand1Recursive = nil
	}
	if operation.operand2Recursive != nil && operation.operand2Recursive.operator == LITERAL {
		operation.operand2 = operation.operand2Recursive.operand1
		operation.operand2Recursive = nil
	}

	if operation.operand1Recursive == nil && operation.operand2Recursive == nil {
		switch operation.operator {
		case XOR:
			return &Operation{operator: LITERAL, operand1: operation.operand1 ^ operation.operand2}
		case MODULO:
			return &Operation{operator: LITERAL, operand1: operation.operand1 % operation.operand2}
		case POWER:
			return &Operation{
				operator: LITERAL,
				operand1: int(math.Pow(float64(operation.operand1), float64(operation.operand2))),
			}
		case DIVIDE:
			return &Operation{operator: LITERAL, operand1: operation.operand1 / operation.operand2}
		case EQUAL:
			equal := 0
			if operation.operand1 == operation.operand2 {
				equal = 1
			}
			return &Operation{operator: LITERAL, operand1: equal}
		case NOT_EQUAL:
			notEqual := 1
			if operation.operand1 == operation.operand2 {
				notEqual = 0
			}
			return &Operation{operator: LITERAL, operand1: notEqual}
		}
	}

	return operation
}

func FindRegAValueWhichMakesQuine(initialState State) int {
	// Observations:
	// Changing regA by 1 changes first output
	// Changing regA by 8 changes second output
	// Changing regA by 64 changes third output
	// etc.
	// (in each case, there is a chance of no change; changing regA by 8^n is necessary, not sufficient)

	regA := 0
	for outputOffset := len(initialState.program) - 1; outputOffset >= 0; outputOffset-- {
		interval := int(math.Pow(8, float64(outputOffset)))
		for {
			state := initialState.Clone()
			state.regA = regA
			output, _, _ := ExecuteProgram(state, nil)
			if len(output) == len(initialState.program) && output[outputOffset] == initialState.program[outputOffset] {
				break
			}
			regA += interval
		}
	}
	return regA
}
