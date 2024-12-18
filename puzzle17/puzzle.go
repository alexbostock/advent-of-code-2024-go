package puzzle17

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"runtime"
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

func ExecuteProgram(state State, expectedOutput []int) (outputStr string, matchesExpected bool) {
	var output []string
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
					return "", false
				}
				offsetInExpectedOutput++
			}
			output = append(output, strconv.Itoa(valToOutput))
		case 6:
			state.regB = state.regA / int(math.Pow(2, float64(comboOperand)))
		case 7:
			state.regC = state.regA / int(math.Pow(2, float64(comboOperand)))
		}
	}

	return strings.Join(output, ","), expectedOutput == nil || len(output) == len(expectedOutput)
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

// Assumptions about the input:
// - final instruction is a branch
// - penultimate instruction is an output
// so we can decide expected number of iterations based on length of expected output
func FindRegAValueWhichMakesQuine(initialState State) int {
	var constraints []*Operation
	regA := &Operation{operator: INPUT_REG_A}
	regB := &Operation{operator: LITERAL, operand1: initialState.regB}
	regC := &Operation{operator: LITERAL, operand1: initialState.regC}

	outputOffset := 0

	for pc := 0; pc < len(initialState.program); pc += 2 {
		opcode := initialState.program[pc]
		operand := &Operation{operator: LITERAL, operand1: initialState.program[pc+1]}
		comboOperand := operand
		switch operand.operand1 {
		case 4:
			comboOperand = regA
		case 5:
			comboOperand = regB
		case 6:
			comboOperand = regC
		}

		switch opcode {
		case 0:
			regA = &Operation{
				operator:          DIVIDE,
				operand1Recursive: regA,
				operand2Recursive: &Operation{operator: POWER, operand1: 2, operand2Recursive: comboOperand},
			}
		case 1:
			regB = &Operation{operator: XOR, operand1Recursive: regB, operand2: operand.operand1}
		case 2:
			regB = &Operation{operator: MODULO, operand1Recursive: comboOperand, operand2: 8}
		case 3:
			if outputOffset == len(initialState.program) {
				constraints = append(constraints, &Operation{operator: EQUAL, operand1Recursive: regA, operand2: 0})
			} else {
				constraints = append(constraints, &Operation{operator: NOT_EQUAL, operand1Recursive: regA, operand2: 0})
				pc = operand.operand1 - 2
			}
		case 4:
			regB = &Operation{operator: XOR, operand1Recursive: regB, operand2Recursive: regC}
		case 5:
			constraints = append(constraints, &Operation{
				operator:          EQUAL,
				operand1Recursive: &Operation{operator: MODULO, operand1Recursive: comboOperand, operand2: 8},
				operand2:          initialState.program[outputOffset],
			})
			outputOffset++
		case 6:
			regB = &Operation{
				operator:          DIVIDE,
				operand1Recursive: regA,
				operand2Recursive: &Operation{operator: POWER, operand1: 2, operand2Recursive: comboOperand},
			}
		case 7:
			regC = &Operation{
				operator:          DIVIDE,
				operand1Recursive: regA,
				operand2Recursive: &Operation{operator: POWER, operand1: 2, operand2Recursive: comboOperand},
			}
		}
	}

	for _, constraint := range constraints {
		fmt.Println(constraint.simplifyConstants())
		fmt.Println()
	}
	// return 0

	for regA := 0; regA < 8; regA++ {
		state := initialState.Clone()
		state.regA = regA
		output, _ := ExecuteProgram(state, nil)
		fmt.Println(regA, output)
	}

	// Starting point derived from the above step (does not work for test example)
	lowerBound := int(math.Pow(8, 16))

	// We also know, based on first output, that regA % 8 == 5 (== 0 for the test example)
	for lowerBound%8 != 5 {
		lowerBound--
	}

	numWorkers := runtime.NumCPU()
	resultChan := make(chan int)
	for i := 0; i < numWorkers; i++ {
		go searchQuineWorker(resultChan, initialState, lowerBound+8*i, 8*numWorkers)
	}

	return <-resultChan
}

func searchQuineWorker(resultChan chan int, initialState State, startSearchAt int, interval int) {
	for regA := startSearchAt; true; regA += interval {
		state := initialState.Clone()
		state.regA = regA
		_, isQuine := ExecuteProgram(state, initialState.program)
		if isQuine {
			resultChan <- regA
			return
		}
	}
}
