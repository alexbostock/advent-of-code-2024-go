package puzzle3

import (
	"io"
	"strconv"
	"strings"
)

type Command struct {
	instruction Instruction
	op1         int
	op2         int
}

type Instruction int

const (
	NOP  Instruction = 0
	Do               = 1
	Dont             = 2
	Mul              = 3
)

func ParseInput3(input io.Reader) []*Command {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, input)
	if err != nil {
		panic(err)
	}
	return parseInput3Str(buf.String())
}

func parseInput3Str(input string) []*Command {
	if len(input) == 0 {
		return nil
	}
	var answer []*Command
	for offset := 0; offset < len(input); {
		if input[offset] == 'm' {
			mul, newOffset := parseMul(input, offset)
			if mul != nil {
				answer = append(answer, mul)
			}
			if newOffset == offset {
				offset++
			} else {
				offset = newOffset
			}
		} else if input[offset] == 'd' {
			doOrDont, newOffset := parseDoOrDont(input, offset)
			if doOrDont == Do {
				answer = append(answer, &Command{Do, 0, 0})
			} else if doOrDont == Dont {
				answer = append(answer, &Command{Dont, 0, 0})
			}
			offset = newOffset
		} else {
			offset++
		}
	}
	return answer
}

func parseMul(input string, offset int) (mul *Command, newOffset int) {
	if len(input) < offset+6 {
		return nil, offset + 1
	}
	if input[offset:offset+4] != "mul(" {
		return nil, offset + 1
	}
	offset += 4
	op1, offsetAfterOp1 := parseInt(input, offset)
	if offsetAfterOp1 == offset {
		return nil, offset
	}
	offset = offsetAfterOp1
	if input[offset] != ',' {
		return nil, offset + 1
	}
	offset++
	op2, offsetAfterOp2 := parseInt(input, offset)
	if offsetAfterOp2 == offset {
		return nil, offset
	}
	offset = offsetAfterOp2
	if input[offset] != ')' {
		return nil, offset
	}
	return &Command{Mul, op1, op2}, offset + 1
}

func parseInt(input string, offset int) (num int, newOffset int) {
	newOffset = offset
	for i := 0; i < 3; i++ {
		if offset+i >= len(input) {
			return
		}
		digit, err := strconv.Atoi(string(input[offset+i]))
		if err != nil {
			// Not a digit
			return
		}
		num = 10*num + digit
		newOffset++
	}
	return
}

func parseDoOrDont(input string, offset int) (instruction Instruction, newOffset int) {
	if len(input) < offset+2 {
		return NOP, offset + 2
	}
	if input[offset:offset+2] != "do" {
		return NOP, offset + 2
	}
	offset += 2
	if len(input) >= offset+2 && input[offset:offset+2] == "()" {
		return Do, offset + 2
	}
	if len(input) >= offset+5 && input[offset:offset+5] == "n't()" {
		return Dont, offset + 2
	}
	return NOP, offset + 1
}

func SumMuls(commands []*Command, observeDosAndDonts bool) int {
	sum := 0
	mulEnabled := true
	for _, command := range commands {
		switch command.instruction {
		case Do:
			mulEnabled = true
		case Dont:
			mulEnabled = false
		case Mul:
			if !observeDosAndDonts || mulEnabled {
				sum += command.op1 * command.op2
			}
		}
	}
	return sum
}
