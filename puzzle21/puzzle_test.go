package puzzle21

import (
	"reflect"
	"slices"
	"testing"
)

func TestTotalComplexityOfCodes(t *testing.T) {
	got := TotalComplexityOfCodes([]string{"029A", "980A", "179A", "456A", "379A"})
	expected := 126384
	if expected != got {
		t.Errorf("expected %d, got %d", expected, got)
	}
}

func TestComplexityOfCode(t *testing.T) {
	examples := []struct {
		code       string
		complexity int
	}{
		{"029A", 68 * 29},
		{"980A", 60 * 980},
		{"179A", 68 * 179},
		{"456A", 64 * 456},
		{"379A", 64 * 379},
	}

	for _, example := range examples {
		got := complexityOfCode(example.code)
		if got != example.complexity {
			t.Errorf("%v: expected %d, got %d", example.code, example.complexity, got)
		}
	}
}

func TestShortestPathToInputCode(t *testing.T) {
	got := shortestPathToInputCode("029A")
	expected := "<vA<AA>>^AvAA<^A>A<v<A>>^AvA^A<vA>^A<v<A>^A>AAvA^A<v<A>A>^AAAvA<^A>A"
	if len(expected) != len(got) {
		t.Errorf("029A: expected length %d eg. %v, got %v (length %d)", len(expected), expected, got, len(got))
	}
}

func TestCandidateNumericKeypadSequenceToInputCode(t *testing.T) {
	t.Run("usual case", func(t *testing.T) {
		got := candidateNumericKeypadSequencesToInputCode([]rune("029A"))
		expected := []Instruction{Left, A, Up, A, Right, Up, Up, A, Down, Down, Down, A}
		for _, sequence := range got {
			if slices.Equal(expected, sequence) {
				return
			}
		}
		t.Errorf("expected sequence not found: expected %v, got %v", expected, got)
	})

	t.Run("never routes over position {0, 3}", func(t *testing.T) {
		got := candidateNumericKeypadSequencesToInputCode([]rune("10"))
		expected := [][]Instruction{{Up, Left, Left, A, Right, Down, A}}
		if !reflect.DeepEqual(expected, got) {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})
}

func TestCandidateDirectionalKeypadSequencesToInputInstructions(t *testing.T) {
	targetInstructions := []Instruction{Left, A, Up, A, Right, Up, Up, A, Down, Down, Down, A}
	got := candidateDirectionalKeypadSequencesToInputInstructions(targetInstructions)
	expected := []Instruction{Down, Left, Left, A, Right, Right, Up, A, Left, A, Right, A, Down, A, Left, Up, A, A, Right, A, Left, Down, A, A, A, Right, Up, A}
	for _, sequence := range got {
		if slices.Equal(expected, sequence) {
			return
		}
	}
	t.Errorf("expected sequence not found: expected %v, got %v", expected, got)
}

func TestAllCandidateKeypadSequences(t *testing.T) {
	got := allCandidateKeypadSequences([]vector{{0, 1}, {3, 2}, {-1, -1}})
	expected := [][]Instruction{
		{Down, A, Right, Right, Right, Down, Down, A, Left, Up, A},
		{Down, A, Right, Right, Right, Down, Down, A, Up, Left, A},
		{Down, A, Down, Down, Right, Right, Right, A, Left, Up, A},
		{Down, A, Down, Down, Right, Right, Right, A, Up, Left, A},
	}
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}
