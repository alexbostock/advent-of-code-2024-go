package puzzle14

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Robot struct {
	pos Vector
	vel Vector
}
type Vector struct {
	x int
	y int
}

func ParseInput14(data io.Reader) []Robot {
	var robots []Robot
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		pos := strings.Split(parts[0], "=")[1]
		vel := strings.Split(parts[1], "=")[1]
		robots = append(robots, Robot{parseVector(pos), parseVector(vel)})
	}
	return robots
}

func parseVector(str string) Vector {
	components := strings.Split(str, ",")
	x, err := strconv.Atoi(components[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(components[1])
	if err != nil {
		panic(err)
	}
	return Vector{x, y}
}

func StateAfterSeconds(robots []Robot, seconds int, width, height int) []Robot {
	updated := make([]Robot, len(robots), len(robots))
	for index, robot := range robots {
		updated[index] = Robot{
			pos: Vector{
				((robot.pos.x+seconds*robot.vel.x)%width + width) % width,
				((robot.pos.y+seconds*robot.vel.y)%height + height) % height,
			},
			vel: robot.vel,
		}
	}
	return updated
}

func SafetyFactor(robots []Robot, width, height int) int {
	quadrant1Count := 0
	quadrant2Count := 0
	quadrant3Count := 0
	quadrant4Count := 0

	xThreshold := (width - 1) / 2
	yThreshold := (height - 1) / 2

	for _, robot := range robots {
		if robot.pos.x < xThreshold && robot.pos.y < yThreshold {
			quadrant1Count++
		} else if robot.pos.x > xThreshold && robot.pos.y < yThreshold {
			quadrant2Count++
		} else if robot.pos.x < xThreshold && robot.pos.y > yThreshold {
			quadrant3Count++
		} else if robot.pos.x > xThreshold && robot.pos.y > yThreshold {
			quadrant4Count++
		}
	}

	return quadrant1Count * quadrant2Count * quadrant3Count * quadrant4Count
}

func PrintEachState(robots []Robot, seconds int, width, height int) {
	for i := 0; i < seconds; i++ {
		layout := PrintRobots(robots, width, height)

		if looksLikePossibleTree(layout) {
			fmt.Println(layout)
			fmt.Println(i)
			return
		}

		robots = StateAfterSeconds(robots, 1, width, height)
	}
}

func PrintRobots(robots []Robot, width, height int) string {
	var stringBuilder strings.Builder

	pixels := make(map[Vector]bool)
	for _, robot := range robots {
		pixels[robot.pos] = true
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if pixels[Vector{x, y}] {
				stringBuilder.WriteString("O")
			} else {
				stringBuilder.WriteString(" ")
			}
		}
		stringBuilder.WriteString("\n")
	}
	stringBuilder.WriteString("\n")

	return stringBuilder.String()
}

func looksLikePossibleTree(layout string) bool {
	return strings.Contains(layout, "OOOOOOOOOOOO")
}
