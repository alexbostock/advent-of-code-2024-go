package puzzle8

import (
	"bufio"
	"io"
)

type Puzzle8Input struct {
	height   int
	width    int
	antennas map[rune][]Coords
}
type Coords struct {
	i int
	j int
}

func ParseInput8(data io.Reader) Puzzle8Input {
	height := 0
	width := 0
	antennas := make(map[rune][]Coords)

	scanner := bufio.NewScanner(data)
	currentLineNum := -1
	for scanner.Scan() {
		currentLineNum++
		height = currentLineNum + 1
		line := scanner.Text()
		width = len(line)
		for indexInLine, char := range []rune(line) {
			if char == '.' {
				continue
			}
			antennas[char] = append(antennas[char], Coords{currentLineNum, indexInLine})
		}
	}

	return Puzzle8Input{height, width, antennas}
}

func CountDistinctAntinodes(area Puzzle8Input) int {
	antinodes := make(map[Coords]bool)
	for _, nodes := range area.antennas {
		for _, node1 := range nodes {
			for _, node2 := range nodes {
				if node1 == node2 {
					continue
				}

				diffI := node1.i - node2.i
				diffJ := node1.j - node2.j

				antinode1 := Coords{node1.i + diffI, node1.j + diffJ}
				antinode2 := Coords{node2.i - diffI, node2.j - diffJ}

				if antinode1.i >= 0 && antinode1.i < area.height && antinode1.j >= 0 && antinode1.j < area.width {
					antinodes[antinode1] = true
				}
				if antinode2.i >= 0 && antinode2.i < area.height && antinode2.j >= 0 && antinode2.j < area.width {
					antinodes[antinode2] = true
				}
			}
		}
	}
	return len(antinodes)
}

func CountDistinctAntinodesWithResonantHarmonics(area Puzzle8Input) int {
	antinodes := make(map[Coords]bool)
	for _, nodes := range area.antennas {
		for _, antenna1 := range nodes {
			for _, antenna2 := range nodes {
				if antenna1 == antenna2 {
					continue
				}
				gradient := float64(antenna1.i-antenna2.i) / float64(antenna1.j-antenna2.j)

				for i := 0; i < area.height; i++ {
					for j := 0; j < area.width; j++ {
						pos := Coords{i, j}
						if antinodes[pos] {
							continue
						}
						gradientToAntenna1 := float64(antenna1.i-i) / float64(antenna1.j-j)
						if gradientToAntenna1 == gradient {
							antinodes[pos] = true
						}
					}
				}
			}
		}
	}
	return len(antinodes)
}
