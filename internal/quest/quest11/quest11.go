package quest11

import (
	"fmt"
	"slices"
	"strconv"
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

func parseInput(content string) map[int]int {
	configuration := map[int]int{}
	splitLines := strings.Split(content, "\n")
	for i, line := range splitLines {
		numDucks, _ := strconv.Atoi(line)
		configuration[i+1] = numDucks
	}
	return configuration
}

func getColOrder(duckConfig map[int]int) []int {
	colOrder := []int{}
	for col := range duckConfig {
		colOrder = append(colOrder, col)
	}
	slices.Sort(colOrder)
	return colOrder
}

func balanceBirds(duckConfig map[int]int, isPhase1 bool, colOrder []int) bool {
	for _, currCol := range colOrder {
		if currCol == 1 {
			continue
		}
		prevCol := currCol-1
		currCount := duckConfig[currCol]
		prevCount := duckConfig[prevCol]

		if isPhase1 {
			if prevCount > currCount {
				duckConfig[currCol] += 1
				duckConfig[prevCol] -= 1
			}
		} else {
			if prevCount < currCount {
				duckConfig[prevCol] += 1
				duckConfig[currCol] -= 1
			}
		}
	}

	// Check if more moves are possible
	phaseComplete := true
	lastCount := 0
	for _, currCol := range colOrder {
		currCount := duckConfig[currCol]
		if currCol == 1 {
			lastCount = currCount
			continue
		}
		
		if isPhase1 {
			phaseComplete = phaseComplete && currCount >= lastCount
			lastCount = currCount
		} else {
			phaseComplete = phaseComplete && currCount <= lastCount
			lastCount = currCount
		}
	}
	return phaseComplete
}

func getChecksum(duckConfig map[int]int) int {
	checksum := 0
	for col, num := range duckConfig {
		checksum += col*num
	}
	return checksum
}

// Logical answers to the parts

func part1(context common.Context) int {
	part := 1
	content, err := common.GetInput(context, part)
	if err != nil {
		fmt.Printf("No input for part %d: %s\n", part, err)
		return 0
	}

	// Logic below
	duckConfig := parseInput(content)
	colOrder := getColOrder(duckConfig)

	isPhase1 := true
	for range 10 {
		phaseComplete := balanceBirds(duckConfig, isPhase1, colOrder)
		if phaseComplete && isPhase1 {
			isPhase1 = false
		} else if phaseComplete {
			break  // Done
		}
	}

	return getChecksum(duckConfig)
}

func part2(context common.Context) int {
	part := 2
	content, err := common.GetInput(context, part)
	if err != nil {
		fmt.Printf("No input for part %d: %s\n", part, err)
		return 0
	}

	duckConfig := parseInput(content)
	colOrder := getColOrder(duckConfig)

	isPhase1 := true
	iters := 0
	for {
		iters += 1
		phaseComplete := balanceBirds(duckConfig, isPhase1, colOrder)
		if phaseComplete && isPhase1 {
			isPhase1 = false
		} else if phaseComplete {
			break  // Done
		}
		break //TEMP
	}
	return iters
}

func part3(context common.Context) int {
	part := 3
	content, err := common.GetInput(context, part)
	if err != nil {
		fmt.Printf("No input for part %d: %s\n", part, err)
		return 0
	}

	duckConfig := parseInput(content)
	totalNumDucks := 0
	totalNumCols := len(duckConfig)
	for _, c := range duckConfig {
		totalNumDucks += c
	}
	ducksPerCol := totalNumDucks / totalNumCols
	/*Note that the input is set up to be monotonically increasing, so we only
	need to find the difference between a column and the required number of 
	ducks per column, and only count columns where ducks are greater than this
	number

	This ONLY works because of the input - the same does not calculate this.

	I could go back and refactor this to make part 2 a lot faster (and come up
	with a more clever algo) but eh
	*/
	totalMoves := 0
	for _, c := range duckConfig {
		if c > ducksPerCol {
			totalMoves += (c - ducksPerCol)
		}
	}
	return totalMoves
}

// Finally, main runner function

func Run(ctx common.Context) error {
	base_start := time.Now()

	part1_answer := part1(ctx)
	part1_time := time.Since(base_start)
	fmt.Printf("Part 1 answer: %d (computed in %s)\n", part1_answer, part1_time)

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
