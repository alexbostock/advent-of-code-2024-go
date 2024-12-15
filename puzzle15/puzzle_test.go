package puzzle15

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestParseInput15(t *testing.T) {
	got := ParseInput15(strings.NewReader(`########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

<^^>>>vv<v>>v<<
`))

	expected := &Warehouse{
		width:  8,
		height: 8,
		robot:  Coords{2, 2},
		boxes: map[Coords]bool{
			{3, 1}: true,
			{5, 1}: true,
			{4, 2}: true,
			{4, 3}: true,
			{4, 4}: true,
			{4, 5}: true,
		},
		walls: map[Coords]bool{
			{0, 0}: true, {1, 0}: true, {2, 0}: true, {3, 0}: true, {4, 0}: true, {5, 0}: true, {6, 0}: true, {7, 0}: true,
			{0, 1}: true, {7, 1}: true,
			{0, 2}: true, {1, 2}: true, {7, 2}: true,
			{0, 3}: true, {7, 3}: true,
			{0, 4}: true, {2, 4}: true, {7, 4}: true,
			{0, 5}: true, {7, 5}: true,
			{0, 6}: true, {7, 6}: true,
			{0, 7}: true, {1, 7}: true, {2, 7}: true, {3, 7}: true, {4, 7}: true, {5, 7}: true, {6, 7}: true, {7, 7}: true,
		},
		instructions: []Instruction{Left, Up, Up, Right, Right, Right, Down, Down, Left, Down, Right, Right, Down, Left, Left},
	}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestExecuteInstructions(t *testing.T) {
	examples := []struct {
		input               string
		expectedRobotEndPos Coords
		expectedSum         int
	}{
		{
			`########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

<^^>>>vv<v>>v<<`,
			Coords{4, 4},
			2028,
		},
		{
			`##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^`,
			Coords{3, 4},
			10092,
		},
	}

	for _, example := range examples {
		warehouse := ParseInput15(strings.NewReader(example.input))
		warehouse.ExecuteInstructions()
		if warehouse.robot != example.expectedRobotEndPos {
			t.Errorf("expected robot at %v, got robot at %v", example.expectedRobotEndPos, warehouse.robot)
		}
		got := warehouse.SumGPSCoordsOfBoxes()
		if example.expectedSum != got {
			t.Errorf("expected %v, got %v", example.expectedSum, got)
		}
	}
}

func TestAttemptToMove(t *testing.T) {
	warehouse := ParseInput15(strings.NewReader(`########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

<^^>>>vv<v>>v<<`))

	fmt.Println(warehouse)
	assertRobotPosition(t, warehouse, Coords{2, 2})
	warehouse.attemptToMove(warehouse.robot, Left)
	assertRobotPosition(t, warehouse, Coords{2, 2})
	warehouse.attemptToMove(warehouse.robot, Up)
	assertRobotPosition(t, warehouse, Coords{2, 1})
	warehouse.attemptToMove(warehouse.robot, Up)
	assertRobotPosition(t, warehouse, Coords{2, 1})
	warehouse.attemptToMove(warehouse.robot, Right)
	assertRobotPosition(t, warehouse, Coords{3, 1})
	assertBoxAt(t, warehouse, Coords{4, 1})
	assertNoBoxAt(t, warehouse, Coords{3, 1})
	warehouse.attemptToMove(warehouse.robot, Right)
	assertRobotPosition(t, warehouse, Coords{4, 1})
	assertBoxAt(t, warehouse, Coords{5, 1})
	assertBoxAt(t, warehouse, Coords{6, 1})
	warehouse.attemptToMove(warehouse.robot, Right)
	assertBoxAt(t, warehouse, Coords{5, 1})
	assertBoxAt(t, warehouse, Coords{6, 1})
	assertNoBoxAt(t, warehouse, Coords{7, 1})
	assertRobotPosition(t, warehouse, Coords{4, 1})
	warehouse.attemptToMove(warehouse.robot, Down)
	assertRobotPosition(t, warehouse, Coords{4, 2})
	warehouse.attemptToMove(warehouse.robot, Down)
	assertRobotPosition(t, warehouse, Coords{4, 2})
}

func assertRobotPosition(t *testing.T, warehouse *Warehouse, pos Coords) {
	if warehouse.robot != pos {
		t.Fatalf("expected robot at %v, got %v", pos, warehouse.robot)
	}
}

func assertBoxAt(t *testing.T, warehouse *Warehouse, pos Coords) {
	if !warehouse.boxes[pos] {
		t.Fatalf("expected box at %v", pos)
	}
}

func assertNoBoxAt(t *testing.T, warehouse *Warehouse, pos Coords) {
	if warehouse.boxes[pos] {
		t.Fatalf("expected no box at %v", pos)
	}
}
