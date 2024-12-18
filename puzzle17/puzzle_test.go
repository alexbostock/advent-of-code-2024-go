package puzzle17

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseInput17(t *testing.T) {
	got := ParseInput17(strings.NewReader(`Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0
`))

	expected := State{
		pc:   0,
		regA: 729, regB: 0, regC: 0,
		program: []int{0, 1, 5, 4, 3, 0},
	}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestExecuteProgram(t *testing.T) {
	state := State{
		pc:   0,
		regA: 729, regB: 0, regC: 0,
		program: []int{0, 1, 5, 4, 3, 0},
	}
	_, output, _ := ExecuteProgram(state, nil)
	expected := "4,6,3,5,6,3,5,2,1,0"

	if output != expected {
		t.Errorf("expected %v, got %v", expected, output)
	}
}

func TestFindRegAValueWhichMakesQuine(t *testing.T) {
	state := State{
		pc:   0,
		regA: 2024, regB: 0, regC: 0,
		program: []int{0, 3, 5, 4, 3, 0},
	}
	output := FindRegAValueWhichMakesQuine(state)
	expected := 117440
	if output != expected {
		t.Errorf("expected %v, got %v", expected, output)
	}
}
