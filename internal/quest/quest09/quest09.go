package quest09

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

type Pair struct {
	Duck1 int
	Duck2 int
}

func init() {
	// Automatically determine quest number from package name
	questNumber = common.GetQuestNumber()
	common.Register(questNumber, fmt.Sprintf("Quest %02d placeholder", questNumber), Run)
}

// Shared helper functions

func parseInput(content string) map[int]string {
	duckMap := make(map[int]string)
	for _, line := range strings.Split(content, "\n") {
		splitLine := strings.Split(line, ":")
		id, _ := strconv.Atoi(splitLine[0])
		sequence := splitLine[1]
		duckMap[id] = sequence
	}

	return duckMap
}

func getSimilarityScore(dna1 string, dna2 string) int {
	simScore := 0
	for i := 0; i < len(dna1); i++ {
		if dna1[i] == dna2[i] {
			simScore += 1
		}
	}
	return simScore
}

func checkChildRelated(child string, parent1 string, parent2 string) bool {
	for i, childDna := range child {
		childDnaStr := string(childDna)
		parent1Dna := string(parent1[i])
		parent2Dna := string(parent2[i])

		if childDnaStr != parent1Dna && childDnaStr != parent2Dna {
			return false
		}

	}
	return true
}

func identifyParents(child int, duckMap map[int]string, allPairs []Pair) Pair {
	for _, pair := range allPairs {
		if child == pair.Duck1 || child == pair.Duck2 {
			continue
		}

		if checkChildRelated(duckMap[child], duckMap[pair.Duck1], duckMap[pair.Duck2]) {
			return pair
		}
	}
	return Pair{Duck1: -1, Duck2: -1}
}

func getAllDucks(duckMap map[int]string) []int {
	ducks := make([]int, 0, len(duckMap))
	for d := range duckMap {
		ducks = append(ducks, d)
	}
	return ducks
}

func getAllParents(duckMap map[int]string, ducks []int) map[int]Pair {
	allCombos := common.GetAllUniquePairs(ducks, ducks)

	allPairs := []Pair{}
	for _, combo := range allCombos {
		if combo[1] <= combo[0] {
			continue
		}
		allPairs = append(allPairs, Pair{Duck1: combo[0], Duck2: combo[1]})
	}

	parentMap := map[int]Pair{}
	for _, duck := range ducks {
		parents := identifyParents(duck, duckMap, allPairs)

		if parents.Duck1 == -1 || parents.Duck2 == -1 {
			continue
		} else {
			parentMap[duck] = parents
		}
	}
	return parentMap
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
	duckMap := parseInput(content)
	simScore := 1
	simScore *= getSimilarityScore(duckMap[1], duckMap[3])
	simScore *= getSimilarityScore(duckMap[2], duckMap[3])

	return simScore
}

func part2(context common.Context) int {
	part := 2
	content, err := common.GetInput(context, part)
	if err != nil {
		fmt.Printf("No input for part %d: %s\n", part, err)
		return 0
	}

	// Logic below
	duckMap := parseInput(content)
	ducks := getAllDucks(duckMap)
	allParents := getAllParents(duckMap, ducks)

	simScore := 0
	for _, duck := range ducks {
		
		parents, ok := allParents[duck]
		if !ok {
			continue
		}

		simScore += (getSimilarityScore(duckMap[duck], duckMap[parents.Duck1]) * 
			getSimilarityScore(duckMap[duck], duckMap[parents.Duck2]))

	}
	return simScore
}

func part3(context common.Context) int {
	part := 3
	content, err := common.GetInput(context, part)
	if err != nil {
		fmt.Printf("No input for part %d: %s\n", part, err)
		return 0
	}

	// Logic below
	duckMap := parseInput(content)
	ducks := getAllDucks(duckMap)
	allParents := getAllParents(duckMap, ducks)
	// TODO - this currently only cares about findings parents from children - 
	// need to refactor this to find all siblings and children

	allFamilies := map[int][]int{}
	for _, baseDuck := range ducks {
		fmt.Println("LOOKING FOR PARENTS OF", baseDuck)
		alreadyFound := false
		for _, family := range allFamilies {
			if slices.Contains(family, baseDuck) {
				alreadyFound = true
				break
			}
		}
		if alreadyFound {
			continue
		}

		// Otherwise, we will try to build a family tree with an extremely simple
		// depth first search
		toExplore := []int{baseDuck}
		processed := []int{}

		iters := 0
		for {
			// fmt.Println("NEW LOOP, curr explore list is: ", toExplore)
			if len(toExplore) == 0 {
				break // found full family
			}
			nextBaseDuck := toExplore[len(toExplore)-1]
			if !slices.Contains(processed, nextBaseDuck) {
				processed = append(processed, nextBaseDuck)
			}
			toExplore = toExplore[:len(toExplore)-1]
			fmt.Println("Popped", nextBaseDuck, "from toExplore which is now", toExplore)

			nextPair, duckHasParents := allParents[nextBaseDuck]
			if !duckHasParents {
				continue
			}
			if nextPair.Duck1 == -1 {
				continue
			}
			toExplore = append(toExplore, []int{nextPair.Duck1, nextPair.Duck2}...)
			fmt.Println("Next up is", toExplore)
			iters += 1
		}
		fullFamily := make([]int, len(processed))
		copy(fullFamily, processed)
		allFamilies[baseDuck] = fullFamily
		fmt.Println(allFamilies)
		
	}

	_ = content
	return 0
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
