package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("./input/1.txt")
	if err != nil {
		panic(err)
	}
	left, right, err := ParseInput(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(DiffLocations(left, right))
	fmt.Println(CalculateSimilarityScore(left, right))
}

// Puzzle 1
func ParseInput(data io.Reader) (left, right []int, err error) {
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		if len(tokens) < 2 {
			return nil, nil, errors.New("Malformatted line " + line)
		}

		leftVal, err := strconv.Atoi(tokens[0])
		if err != nil {
			return nil, nil, err
		}
		rightVal, err := strconv.Atoi(tokens[len(tokens)-1])
		if err != nil {
			return nil, nil, err
		}

		left = append(left, leftVal)
		right = append(right, rightVal)
	}

	return
}

func DiffLocations(left, right []int) int {
	slices.Sort(left)
	slices.Sort(right)
	distance := 0
	for index, leftVal := range left {
		rightVal := right[index]
		if leftVal > rightVal {
			distance += leftVal - rightVal
		} else {
			distance += rightVal - leftVal
		}
	}
	return distance
}

func CalculateSimilarityScore(left, right []int) int {
	slices.Sort(left)
	slices.Sort(right)
	score := 0
	rightCounts := make(map[int]int)
	for _, val := range right {
		rightCounts[val]++
	}

	for _, val := range left {
		score += val * rightCounts[val]
	}
	return score
}
