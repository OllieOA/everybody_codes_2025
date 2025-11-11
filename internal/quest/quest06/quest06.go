package quest06

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

func getPairings(positions map[string][]int, novice string) int {
	knight := strings.ToUpper(novice)

	pairCount := 0
	for _, novicePos := range positions[novice] {
		for _, knightPos := range positions[knight] {
			if knightPos < novicePos {
				pairCount += 1
			}
		}
	}

	return pairCount
}

func getPositions(content string) map[string][]int {
	positions := map[string][]int{}

	for i, charRune := range content {
		char := string(charRune)  // I don't understand how to make this simpler
		if _, ok := positions[char]; !ok {
			positions[char] = []int{}
		}
		positions[char] = append(positions[char], i)
	}

	return positions
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
	positions := getPositions(content)
	numPairs := getPairings(positions, "a")
	return numPairs
}

func part2(context common.Context) int {
	part := 2
	content, err := common.GetInput(context, part)
	if err != nil {
		fmt.Printf("No input for part %d: %s\n", part, err)
		return 0
	}

	// Logic below
	positions := getPositions(content)
	configurationCount := 0
	novices := []string{"a", "b", "c"}
	for _, novice := range novices {
		configurationCount += getPairings(positions, novice)
	}

	return configurationCount
}

func findAllMentors(positions map[string][]int, novice string, noviceMin int, noviceMax int, knightRange int) int {
	knight := strings.ToUpper(novice)

	totalPairs := 0
	
	for novI := noviceMin; novI < noviceMax; novI++ {
		if !slices.Contains(positions[novice], novI) {
			continue
		}
		totalForNov := 0
		minKnight := novI - knightRange
		maxKnight := novI + knightRange
		for knI := minKnight; knI <= maxKnight; knI++ {
			if slices.Contains(positions[knight], knI) {
				totalForNov += 1
				totalPairs += 1
			}
		}
		// fmt.Printf("novice %s at %d can pair with %d knights\n", novice, novI, totalForNov)
	}
	return totalPairs
}

func part3(context common.Context) int {
	part := 3
	content, err := common.GetInput(context, part)
	if err != nil {
		fmt.Printf("No input for part %d: %s\n", part, err)
		return 0
	}

	// Logic below
	limit := 1000
	repeats := 1000
	if context.Sample {
		limit = 10
		repeats = 4
	}

	_= repeats

	blockLength := len(content)
	doubleLengthInput := content + content  // Use this to calculate lead and tail
	tripleLengthInput := content + content + content  // Use this to calculate central "blocks"

	positionsEdge := getPositions(doubleLengthInput)
	positionsCore := getPositions(tripleLengthInput)

	novices := []string{"a", "b", "c"}

	leadPairings := 0
	tailPairings := 0
	corePairings := 0

	for _, novice := range novices {
		// fmt.Println("===== NOVICE",novice,"=====")
		
		// fmt.Println("=== LEADING EDGE")
		leadPairings += findAllMentors(positionsEdge, novice, 0, blockLength, limit)
		// fmt.Println("=== TAILING EDGE")
		tailPairings += findAllMentors(positionsEdge, novice, blockLength, 2*blockLength, limit)
		// fmt.Println("=== CORE")
		corePairings += findAllMentors(positionsCore, novice, blockLength, 2*blockLength, limit)

	}

	totalPairings := leadPairings + tailPairings + ((repeats-2) * corePairings)
	// fmt.Printf("lead (%d) + tail (%d) + ((%d-2)*core(%d)) = %d total pairings\n", leadPairings, tailPairings, repeats, corePairings, totalPairings)
	//1667666002; length and first char correct
	return totalPairings
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
