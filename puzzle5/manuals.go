package puzzle5

import (
	"bufio"
	"io"
	"slices"
	"strconv"
	"strings"
)

type Puzzle5Input struct {
	rules   []pageOrderingRule
	updates [][]int
}
type pageOrderingRule struct {
	before int
	after  int
}

func ParseInput5(data io.Reader) Puzzle5Input {
	var rules []pageOrderingRule
	var updates [][]int

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if strings.Contains(line, "|") {
			tokens := strings.Split(line, "|")
			before, err := strconv.Atoi(tokens[0])
			if err != nil {
				panic(err)
			}
			after, err := strconv.Atoi(tokens[1])
			if err != nil {
				panic(err)
			}
			rules = append(rules, struct {
				before int
				after  int
			}{before, after})
		} else {
			tokens := strings.Split(line, ",")
			var update []int
			for _, token := range tokens {
				num, err := strconv.Atoi(token)
				if err != nil {
					panic(err)
				}
				update = append(update, num)
			}
			updates = append(updates, update)
		}
	}

	return Puzzle5Input{rules, updates}
}

func SumMiddlePagesOfCorrectlyOrderedUpdates(input Puzzle5Input) int {
	rules := loadPageOrderingRules(input.rules)
	sum := 0
	for _, update := range input.updates {
		if isValidUpdate(update, rules) {
			sum += middlePageInUpdate(update)
		}
	}
	return sum
}

type pageOrderingRules struct {
	rules        map[pageOrderingRule]bool
	byBeforePage map[int][]int
	byAfterPage  map[int][]int
}

func loadPageOrderingRules(rules []pageOrderingRule) pageOrderingRules {
	loadedRules := pageOrderingRules{
		rules:        make(map[pageOrderingRule]bool),
		byBeforePage: make(map[int][]int),
		byAfterPage:  make(map[int][]int),
	}
	for index, rule := range rules {
		loadedRules.rules[rule] = true
		loadedRules.byBeforePage[rule.before] = append(loadedRules.byBeforePage[rule.before], index)
		loadedRules.byAfterPage[rule.after] = append(loadedRules.byAfterPage[rule.after], index)
	}
	return loadedRules
}

func isValidUpdate(update []int, rules pageOrderingRules) bool {
	rulesClosed := make(map[int]bool)
	for _, page := range update {
		for _, ruleID := range rules.byAfterPage[page] {
			rulesClosed[ruleID] = true
		}
		for _, ruleID := range rules.byBeforePage[page] {
			if rulesClosed[ruleID] {
				return false
			}
		}
	}
	return true
}

func middlePageInUpdate(update []int) int {
	return update[(len(update)+1)/2-1]
}

func SumMiddlePagesOfFixedUpdates(input Puzzle5Input) int {
	rules := loadPageOrderingRules(input.rules)
	sum := 0
	for _, update := range input.updates {
		if !isValidUpdate(update, rules) {
			fixed := fixUpdate(update, rules)
			sum += middlePageInUpdate(fixed)
		}
	}
	return sum
}

func fixUpdate(update []int, rules pageOrderingRules) []int {
	fixed := slices.Clone(update)
	slices.SortStableFunc(fixed, func(page1, page2 int) int {
		if rules.rules[pageOrderingRule{page1, page2}] {
			return -1
		} else if rules.rules[pageOrderingRule{page2, page1}] {
			return 1
		}
		return 0
	})
	return fixed
}
