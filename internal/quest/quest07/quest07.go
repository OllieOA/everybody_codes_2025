package quest07

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/OllieOA/everybody_codes_2025/internal/common"
)

// Boilerplate

var questNumber int

func init() {
	// Automatically determine quest number from package name
	questNumber = common.GetQuestNumber()
	common.Register(questNumber, fmt.Sprintf("Quest %02d placeholder", questNumber), Run)
}

// Shared helper functions

func parseInput(content string) ([]string, map[string][]string) {
	newLineSplit := strings.Split(content, "\n\n")
	namesString := newLineSplit[0]
	rulesString := newLineSplit[1]
	
	names := strings.Split(namesString, ",")

	rulesSplit := strings.Split(rulesString, "\n")
	rules := make(map[string][]string, len(rulesSplit))

	for _, ruleStr := range rulesSplit {
		ruleParts := strings.Split(ruleStr, " > ")
		sourceLetter := ruleParts[0]
		targetLetters := strings.Split(ruleParts[1], ",")
		rules[sourceLetter] = targetLetters
	}

	return names, rules
}

func checkName(name string, rules map[string][]string) bool {
	for i := 1; i < len(name); i++ {
		sourceLetter := string(name[i-1])
		targetLetter := string(name[i])

		if !slices.Contains(rules[sourceLetter], targetLetter) {
			return false
		}
	}
	return true
}

// Logical answers to the parts

func part1(context common.Context) string {
	part := 1
	content, err := common.GetInput(context, part)
	if err != nil {
		fmt.Printf("No input for part %d: %s\n", part, err)
		return ""
	}

	names, rules := parseInput(content)
	// Logic below

	for _, name := range names {
		if checkName(name, rules) {
			return name
		}
	}

	return ""
}

func part2(context common.Context) int {
	part := 2
	content, err := common.GetInput(context, part)
	if err != nil {
		fmt.Printf("No input for part %d: %s\n", part, err)
		return 0
	}

	names, rules := parseInput(content)
	// Logic below
	validIndices := 0

	for i, name := range names {
		if checkName(name, rules) {
			validIndices += i + 1
		}
	}

	return validIndices
}

func getCleanedPrefixes(prefixes []string, rules map[string][]string) []string {
	/* This will check if any prefixes are contained within other prefixes, then
	find all other prefixes within the ruleset. It will also remove any prefix
	that currently breaks the rules
	*/
	validPrefixes := []string{}
	for _, prefix := range prefixes {
		if checkName(prefix, rules) {
			validPrefixes = append(validPrefixes, prefix)
		}
	}

	cleanPrefixes := make([]string, len(validPrefixes))
	copy(cleanPrefixes, validPrefixes)
	
	for {
		resolvesRequired := []string{}
		for i, testPrefix := range cleanPrefixes {
			for j, otherPrefix := range cleanPrefixes {
				if i == j {continue}
				if strings.Contains(otherPrefix, testPrefix) {
					resolvesRequired = append(resolvesRequired, testPrefix)
				}
			}
		}

		if len(resolvesRequired) == 0 {
			return cleanPrefixes
		}

		newPrefixes := []string{}
		for _, prefix := range cleanPrefixes {
			if !slices.Contains(resolvesRequired, prefix) {
				newPrefixes = append(newPrefixes, prefix)
			}
		}

		// Otherwise, let's try and resolve what is in resolvesRequired
		for _, subPrefix := range resolvesRequired {
			resolvedSubPrefixes := []string{}
			for _, letter := range rules[string(subPrefix[len(subPrefix)-1])] {
				resolvedSubPrefixes = append(resolvedSubPrefixes, subPrefix + letter)
			}
			newPrefixes = append(newPrefixes, resolvedSubPrefixes...)
		}

		cleanPrefixes = []string{}
		for _, prefix := range newPrefixes {
			if !slices.Contains(cleanPrefixes, prefix) {
				cleanPrefixes = append(cleanPrefixes, prefix)
			}
		}
	}
}

func part3(context common.Context) int {
	part := 3
	content, err := common.GetInput(context, part)
	if err != nil {
		fmt.Printf("No input for part %d: %s\n", part, err)
		return 0
	}

	prefixes, rules := parseInput(content)
	minLen := 7
	maxLen := 11
	// Logic below

	validNames := 0

	cleanPrefixes := getCleanedPrefixes(prefixes, rules)

	for _, prefix := range cleanPrefixes {
		currLen := len(prefix)
		currLetters := []string{string(prefix[currLen-1])}

		for i := currLen; i <= maxLen-1; i++ {
			nextLetters := []string{}
			for _, letter := range currLetters {
				if _, ok := rules[letter]; !ok {
					continue
				}
				nextLetters = append(nextLetters, rules[letter]...)
				currLen += 1
			}

			if i >= minLen-1 {
				validNames += len(nextLetters)
			}

			currLetters = make([]string, len(nextLetters))
			copy(currLetters, nextLetters)
		}
	}
	return validNames
}

// Finally, main runner function

func Run(ctx common.Context) error {
	base_start := time.Now()

	part1_answer := part1(ctx)
	part1_time := time.Since(base_start)
	fmt.Printf("Part 1 answer: %s (computed in %s)\n", part1_answer, part1_time)

	part2_answer := part2(ctx)
	part2_time := time.Since(base_start) - part1_time
	fmt.Printf("Part 2 answer: %d (computed in %s)\n", part2_answer, part2_time)

	part3_answer := part3(ctx)
	part3_time := time.Since(base_start) - part1_time - part2_time
	fmt.Printf("Part 3 answer: %d (computed in %s)\n", part3_answer, part3_time)

	total_time := time.Since(base_start)
	fmt.Printf("Total time: %s\n", total_time)

	return nil
}
