package puzzle3

import (
	"reflect"
	"testing"
)

func TestParseInput3(t *testing.T) {
	examples := []struct {
		input    string
		expected []*Command
	}{
		{"", nil},
		{"mul(1,1)", []*Command{{Mul, 1, 1}}},
		{"mul(131,45)", []*Command{{Mul, 131, 45}}},
		{"mul(1234,7)", nil},
		{" mul(2,3)", []*Command{{Mul, 2, 3}}},
		{" mul(2, 3)", nil},
		{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))", []*Command{{Mul, 2, 4}, {Mul, 5, 5}, {Mul, 11, 8}, {Mul, 8, 5}}},
	}

	for _, example := range examples {
		got := parseInput3Str(example.input)
		gotExpected := len(got) == 0 && len(example.expected) == 0 || reflect.DeepEqual(got, example.expected)
		if !gotExpected {
			t.Errorf("%v: expected %v, got %v", example.input, example.expected, got)
		}
	}
}

func TestSumMuls(t *testing.T) {
	got := SumMuls([]*Command{{Mul, 2, 4}, {Mul, 5, 5}, {Mul, 11, 8}, {Mul, 8, 5}}, false)
	expected := 161
	if got != expected {
		t.Errorf("got %v, expected %v", got, expected)
	}
}

func TestSumMulsWithDoAndDont(t *testing.T) {
	commands := parseInput3Str("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))")
	got := SumMuls(commands, true)
	expected := 48
	if got != expected {
		t.Errorf("got %v, expected %v", got, expected)
	}
}
