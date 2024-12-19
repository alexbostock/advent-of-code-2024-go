package puzzle19

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

type Input struct {
	towels  []string
	designs []string
}

func ParseInput19(input io.Reader) Input {
	parsed := Input{}
	scanner := bufio.NewScanner(input)

	scanner.Scan()
	parsed.towels = strings.Split(scanner.Text(), ", ")
	scanner.Scan()

	for scanner.Scan() {
		parsed.designs = append(parsed.designs, scanner.Text())
	}

	return parsed
}

func CountPossibleDesigns(input Input) int {
	pattern := regexp.MustCompile(
		fmt.Sprintf("^(%s)*$", strings.Join(input.towels, "|")),
	)
	count := 0
	for _, design := range input.designs {
		if pattern.MatchString(design) {
			count++
		}
	}
	return count
}

func SumPossibleWaysOfMakingDesigns(input Input) int {
	cache := make(map[string]int)

	sum := 0
	for _, design := range input.designs {
		sum += countPossibleWaysOfMakingDesign(input.towels, design, cache)
	}
	return sum
}

func countPossibleWaysOfMakingDesign(towels []string, design string, cache map[string]int) int {
	if design == "" {
		return 1
	}
	cached, isCached := cache[design]
	if isCached {
		return cached
	}

	count := 0
	for _, towel := range towels {
		if len(design) >= len(towel) && design[0:len(towel)] == towel {
			count += countPossibleWaysOfMakingDesign(towels, design[len(towel):], cache)
		}
	}

	cache[design] = count
	return count
}
