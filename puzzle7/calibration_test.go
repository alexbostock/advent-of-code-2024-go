package puzzle7

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseInput7(t *testing.T) {
	got := ParseInput7(strings.NewReader(`190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
`))

	expected := []CalibrationEquation{
		{190, []int{10, 19}},
		{3267, []int{81, 40, 27}},
		{83, []int{17, 5}},
		{156, []int{15, 6}},
		{7290, []int{6, 8, 6, 15}},
		{161011, []int{16, 10, 13}},
		{192, []int{17, 8, 14}},
		{21037, []int{9, 7, 18, 13}},
		{292, []int{11, 6, 16, 20}},
	}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestSumValidCalibrationEquations(t *testing.T) {
	equations := []CalibrationEquation{
		{190, []int{10, 19}},
		{3267, []int{81, 40, 27}},
		{83, []int{17, 5}},
		{156, []int{15, 6}},
		{7290, []int{6, 8, 6, 15}},
		{161011, []int{16, 10, 13}},
		{192, []int{17, 8, 14}},
		{292, []int{11, 6, 16, 20}},
	}

	got := SumValidCalibrationEquations(equations, false)
	if got != 3749 {
		t.Errorf("expected 3749, got %v", got)
	}
}

func TestCanSolveCalibrationEquation(t *testing.T) {
	examples := []struct {
		equation CalibrationEquation
		expected bool
	}{
		{CalibrationEquation{190, []int{10, 19}}, true},
		{CalibrationEquation{3267, []int{81, 40, 27}}, true},
		{CalibrationEquation{83, []int{17, 5}}, false},
		{CalibrationEquation{156, []int{15, 6}}, false},
		{CalibrationEquation{7290, []int{6, 8, 6, 15}}, false},
		{CalibrationEquation{161011, []int{16, 10, 13}}, false},
		{CalibrationEquation{192, []int{17, 8, 14}}, false},
		{CalibrationEquation{292, []int{11, 6, 16, 20}}, true},
	}

	for _, example := range examples {
		got := canSolveCalibrationEquation(example.equation, false)
		if got != example.expected {
			t.Errorf("%v: expected %v, got %v", example.equation, example.expected, got)
		}
	}
}

func TestCanSolveCalibrationEquationWithConcatenation(t *testing.T) {
	examples := []struct {
		equation CalibrationEquation
		expected bool
	}{
		{CalibrationEquation{190, []int{10, 19}}, true},
		{CalibrationEquation{3267, []int{81, 40, 27}}, true},
		{CalibrationEquation{83, []int{17, 5}}, false},
		{CalibrationEquation{156, []int{15, 6}}, true},

		{CalibrationEquation{6, []int{6}}, true},
		{CalibrationEquation{48, []int{6, 8}}, true},
		{CalibrationEquation{486, []int{6, 8, 6}}, true},
		{CalibrationEquation{7290, []int{6, 8, 6, 15}}, true},

		{CalibrationEquation{161011, []int{16, 10, 13}}, false},
		{CalibrationEquation{192, []int{17, 8, 14}}, true},
		{CalibrationEquation{292, []int{11, 6, 16, 20}}, true},
	}

	for _, example := range examples {
		got := canSolveCalibrationEquation(example.equation, true)
		if got != example.expected {
			t.Errorf("%v: expected %v, got %v", example.equation, example.expected, got)
		}
	}
}
