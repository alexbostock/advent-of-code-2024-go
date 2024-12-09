package puzzle9

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseInput9(t *testing.T) {
	got := ParseInput9(strings.NewReader("2333133121414131402\n"))
	expected := Input{
		Blocks: []int{0, 0, -1, -1, -1, 1, 1, 1, -1, -1, -1, 2, -1, -1, -1, 3, 3, 3, -1, 4, 4, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1, 7, 7, 7, -1, 8, 8, 8, 8, 9, 9},
		Files: []File{
			{false, 0, 2},
			{true, 0, 3},
			{false, 1, 3},
			{true, 0, 3},
			{false, 2, 1},
			{true, 0, 3},
			{false, 3, 3},
			{true, 0, 1},
			{false, 4, 2},
			{true, 0, 1},
			{false, 5, 4},
			{true, 0, 1},
			{false, 6, 4},
			{true, 0, 1},
			{false, 7, 3},
			{true, 0, 1},
			{false, 8, 4},
			{false, 9, 2},
		},
	}
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestParseByteAsDigit(t *testing.T) {
	examples := []struct {
		char byte
		num  int
		err  bool
	}{
		{'0', 0, false},
		{'1', 1, false},
		{'9', 9, false},
		{'a', 0, true},
		{'.', 0, true},
	}

	for _, example := range examples {
		got, err := parseByteAsDigit(example.char)
		if example.err {
			if err == nil {
				t.Errorf("expected non-nil error parsing %v", example.char)
			}
		} else {
			if err != nil {
				t.Errorf("unexpected error parsing %v", example.char)
			}
			if got != example.num {
				t.Errorf("expected %v, got %v", example.num, got)
			}
		}
	}
}

func TestMoveBlocks(t *testing.T) {
	blocks := []int{0, 0, -1, -1, -1, 1, 1, 1, -1, -1, -1, 2, -1, -1, -1, 3, 3, 3, -1, 4, 4, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1, 7, 7, 7, -1, 8, 8, 8, 8, 9, 9}
	MoveBlocks(blocks)
	expected := []int{0, 0, 9, 9, 8, 1, 1, 1, 8, 8, 8, 2, 7, 7, 7, 3, 3, 3, 6, 4, 4, 6, 5, 5, 5, 5, 6, 6, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}

	if !reflect.DeepEqual(blocks, expected) {
		t.Errorf("expected %v, got %v", expected, blocks)
	}
}

func TestMoveFiles(t *testing.T) {
	files := []File{
		{false, 0, 2},
		{true, 0, 3},
		{false, 1, 3},
		{true, 0, 3},
		{false, 2, 1},
		{true, 0, 3},
		{false, 3, 3},
		{true, 0, 1},
		{false, 4, 2},
		{true, 0, 1},
		{false, 5, 4},
		{true, 0, 1},
		{false, 6, 4},
		{true, 0, 1},
		{false, 7, 3},
		{true, 0, 1},
		{false, 8, 4},
		{false, 9, 2},
	}
	got := MoveFiles(files)
	expected := []File{
		{false, 0, 2},
		{false, 9, 2},
		{false, 2, 1},
		{false, 1, 3},
		{false, 7, 3},
		{true, 0, 1},
		{false, 4, 2},
		{true, 0, 1},
		{false, 3, 3},
		{true, 0, 4},
		{false, 5, 4},
		{true, 0, 1},
		{false, 6, 4},
		{true, 0, 5},
		{false, 8, 4},
		{true, 0, 2},
	}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestComputeChecksum(t *testing.T) {
	blocks := []int{0, 0, 9, 9, 8, 1, 1, 1, 8, 8, 8, 2, 7, 7, 7, 3, 3, 3, 6, 4, 4, 6, 5, 5, 5, 5, 6, 6, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
	checksum := ComputeChecksum(blocks)
	expected := 1928

	if checksum != expected {
		t.Errorf("expected %v, got %v", expected, checksum)
	}
}

func TestComputeChecksumFiles(t *testing.T) {
	files := []File{
		{false, 0, 2},
		{false, 9, 2},
		{false, 2, 1},
		{false, 1, 3},
		{false, 7, 3},
		{true, 0, 1},
		{false, 4, 2},
		{true, 0, 1},
		{false, 3, 3},
		{true, 0, 4},
		{false, 5, 4},
		{true, 0, 1},
		{false, 6, 4},
		{true, 0, 5},
		{false, 8, 4},
		{true, 0, 2},
	}
	got := ComputeChecksumFiles(files)
	expected := 2858

	if expected != got {
		t.Errorf("expected %v, got %v", expected, got)
	}
}
