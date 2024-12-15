package puzzle15

import (
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

func TestParseInput15Wide(t *testing.T) {
	got := ParseInput15Wide(strings.NewReader(`########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

<^^>>>vv<v>>v<<
`))

	expected := &WideWarehouse{
		width:  16,
		height: 8,
		robot:  Coords{4, 2},
		boxes: map[Coords]bool{
			{6, 1}:  true,
			{10, 1}: true,
			{8, 2}:  true,
			{8, 3}:  true,
			{8, 4}:  true,
			{8, 5}:  true,
		},
		walls: map[Coords]bool{
			{0, 0}: true, {1, 0}: true, {2, 0}: true, {3, 0}: true, {4, 0}: true, {5, 0}: true, {6, 0}: true, {7, 0}: true,
			{8, 0}: true, {9, 0}: true, {10, 0}: true, {11, 0}: true, {12, 0}: true, {13, 0}: true, {14, 0}: true, {15, 0}: true,
			{0, 1}: true, {1, 1}: true, {14, 1}: true, {15, 1}: true,
			{0, 2}: true, {1, 2}: true, {2, 2}: true, {3, 2}: true, {14, 2}: true, {15, 2}: true,
			{0, 3}: true, {1, 3}: true, {14, 3}: true, {15, 3}: true,
			{0, 4}: true, {1, 4}: true, {4, 4}: true, {5, 4}: true, {14, 4}: true, {15, 4}: true,
			{0, 5}: true, {1, 5}: true, {14, 5}: true, {15, 5}: true,
			{0, 6}: true, {1, 6}: true, {14, 6}: true, {15, 6}: true,
			{0, 7}: true, {1, 7}: true, {2, 7}: true, {3, 7}: true, {4, 7}: true, {5, 7}: true, {6, 7}: true, {7, 7}: true,
			{8, 7}: true, {9, 7}: true, {10, 7}: true, {11, 7}: true, {12, 7}: true, {13, 7}: true, {14, 7}: true, {15, 7}: true,
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

func TestExecuteInstructionsWide(t *testing.T) {
	warehouse := ParseInput15Wide(strings.NewReader(`##########
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
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^`))

	warehouse.ExecuteInstructions()

	warehouse.assertRobotPosition(t, Coords{4, 7})

	got := warehouse.SumGPSCoordsOfBoxes()
	expected := 9021

	if expected != got {
		t.Errorf("expected %v, got %v", expected, got)
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

	warehouse.assertRobotPosition(t, Coords{2, 2})
	warehouse.attemptToMove(warehouse.robot, Left)
	warehouse.assertRobotPosition(t, Coords{2, 2})
	warehouse.attemptToMove(warehouse.robot, Up)
	warehouse.assertRobotPosition(t, Coords{2, 1})
	warehouse.attemptToMove(warehouse.robot, Up)
	warehouse.assertRobotPosition(t, Coords{2, 1})
	warehouse.attemptToMove(warehouse.robot, Right)
	warehouse.assertRobotPosition(t, Coords{3, 1})
	warehouse.assertBoxAt(t, Coords{4, 1})
	warehouse.assertNoBoxAt(t, Coords{3, 1})
	warehouse.attemptToMove(warehouse.robot, Right)
	warehouse.assertRobotPosition(t, Coords{4, 1})
	warehouse.assertBoxAt(t, Coords{5, 1})
	warehouse.assertBoxAt(t, Coords{6, 1})
	warehouse.attemptToMove(warehouse.robot, Right)
	warehouse.assertBoxAt(t, Coords{5, 1})
	warehouse.assertBoxAt(t, Coords{6, 1})
	warehouse.assertNoBoxAt(t, Coords{7, 1})
	warehouse.assertRobotPosition(t, Coords{4, 1})
	warehouse.attemptToMove(warehouse.robot, Down)
	warehouse.assertRobotPosition(t, Coords{4, 2})
	warehouse.attemptToMove(warehouse.robot, Down)
	warehouse.assertRobotPosition(t, Coords{4, 2})
}

func TestAttemptToMoveWide(t *testing.T) {
	warehouse := ParseInput15Wide(strings.NewReader(`#######
#...#.#
#.....#
#..OO@#
#..O..#
#.....#
#######

<vv<<^^<<^^`))

	warehouse.assertRobotPosition(t, Coords{10, 3})
	warehouse.assertBoxAt(t, Coords{8, 3})
	warehouse.assertBoxAt(t, Coords{6, 3})
	warehouse.attemptToMove(Coords{10, 3}, Left, false)
	warehouse.assertRobotPosition(t, Coords{9, 3})
	warehouse.assertBoxAt(t, Coords{7, 3})
	warehouse.assertBoxAt(t, Coords{5, 3})

	warehouse.attemptToMove(warehouse.robot, Down, false)
	warehouse.attemptToMove(warehouse.robot, Down, false)
	warehouse.attemptToMove(warehouse.robot, Left, false)
	warehouse.attemptToMove(warehouse.robot, Left, false)
	warehouse.assertRobotPosition(t, Coords{7, 5})

	warehouse.attemptToMove(warehouse.robot, Up, false)
	warehouse.assertRobotPosition(t, Coords{7, 4})
	warehouse.assertBoxAt(t, Coords{5, 2})
	warehouse.assertBoxAt(t, Coords{7, 2})
	warehouse.assertBoxAt(t, Coords{6, 3})
	warehouse.assertNoBoxAt(t, Coords{7, 3})
	warehouse.assertNoBoxAt(t, Coords{5, 3})

	warehouse.attemptToMove(warehouse.robot, Up, false)
	warehouse.assertRobotPosition(t, Coords{7, 4})
	warehouse.assertBoxAt(t, Coords{5, 2})
	warehouse.assertBoxAt(t, Coords{7, 2})
	warehouse.assertBoxAt(t, Coords{6, 3})
	warehouse.assertNoBoxAt(t, Coords{6, 2})
	warehouse.assertNoBoxAt(t, Coords{7, 1})

	warehouse.attemptToMove(warehouse.robot, Left, false)
	warehouse.attemptToMove(warehouse.robot, Left, false)
	warehouse.attemptToMove(warehouse.robot, Up, false)
	warehouse.attemptToMove(warehouse.robot, Up, false)
	warehouse.assertRobotPosition(t, Coords{5, 2})
	warehouse.assertBoxAt(t, Coords{5, 1})
	warehouse.assertNoBoxAt(t, Coords{5, 2})
	warehouse.assertBoxAt(t, Coords{7, 2})
	warehouse.assertBoxAt(t, Coords{6, 3})
}

func (warehouse *Warehouse) assertRobotPosition(t *testing.T, pos Coords) {
	if warehouse.robot != pos {
		t.Fatalf("expected robot at %v, got %v", pos, warehouse.robot)
	}
}
func (warehouse *WideWarehouse) assertRobotPosition(t *testing.T, pos Coords) {
	if warehouse.robot != pos {
		t.Fatalf("expected robot at %v, got %v", pos, warehouse.robot)
	}
}

func (warehouse *Warehouse) assertBoxAt(t *testing.T, pos Coords) {
	if !warehouse.boxes[pos] {
		t.Fatalf("expected box at %v", pos)
	}
}
func (warehouse *WideWarehouse) assertBoxAt(t *testing.T, pos Coords) {
	if !warehouse.boxes[pos] {
		t.Fatalf("expected box at %v", pos)
	}
}

func (warehouse *Warehouse) assertNoBoxAt(t *testing.T, pos Coords) {
	if warehouse.boxes[pos] {
		t.Fatalf("expected no box at %v", pos)
	}
}
func (warehouse *WideWarehouse) assertNoBoxAt(t *testing.T, pos Coords) {
	if warehouse.boxes[pos] {
		t.Fatalf("expected no box at %v", pos)
	}
}

func TestSumGPSCoordsOfBoxesWide(t *testing.T) {
	warehouse := ParseInput15Wide(strings.NewReader(`#####
#..O@
`))
	warehouse.attemptToMove(warehouse.robot, Left, false)

	got := warehouse.SumGPSCoordsOfBoxes()
	expected := 105
	if expected != got {
		t.Errorf("expected %v, got %v", expected, got)
	}
}
