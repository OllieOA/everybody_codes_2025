package quest03

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

func getCrateSizes(content string) []int {
	crateSizesStr := strings.Split(content, ",")
	crateSizes := []int{}
	for _, size := range crateSizesStr {
		crateSize, _ := strconv.Atoi(size)
		crateSizes = append(crateSizes, crateSize)
	}
	return crateSizes
}

func getPackedCrateFromVal(val int, crates []int) []int {
	var packedCrate []int
	for size := val; size > 0; size -= 1 {
		if slices.Contains(crates, size) {
			packedCrate = append(packedCrate, size)
		}
	}

	return packedCrate
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
	// crateSizes := getCrateSizes(content)
	crateSizes, _ := common.ParseDelimitedList(content, strconv.Atoi, ",")

	maxSize := slices.Max(crateSizes)

	packedCrate := getPackedCrateFromVal(maxSize, crateSizes)

	packedCrateVal := 0
	for _, crate := range packedCrate {
		packedCrateVal += crate
	}

	return packedCrateVal
}

func part2(context common.Context) int {
	part := 2
	content, err := common.GetInput(context, part)
	if err != nil {
		fmt.Printf("No input for part %d: %s\n", part, err)
		return 0
	}

	// Logic below
	crateSizes, _ := common.ParseDelimitedList(content, strconv.Atoi, ",")
	slices.Sort(crateSizes)

	crateSizesUnique, _ := common.FindUniqueValuesInIntegerArray(crateSizes)

	for _, v := range crateSizesUnique {
		packedCrate := getPackedCrateFromVal(v, crateSizes)
		if len(packedCrate) == 20 {
			packedCrateVal := 0
				for _, crate := range packedCrate {
					packedCrateVal += crate
				}
			return packedCrateVal
		}
	}

	return -1
}

func part3(context common.Context) int {
	part := 3
	content, err := common.GetInput(context, part)
	if err != nil {
		fmt.Printf("No input for part %d: %s\n", part, err)
		return 0
	}

	// Logic below
	crateSizes, _ := common.ParseDelimitedList(content, strconv.Atoi, ",")
	slices.Sort(crateSizes)
	slices.Reverse(crateSizes)

	var packedCrates [][]int

	for i := 0; len(crateSizes) > 0; i++ {
		newCrate := getPackedCrateFromVal(crateSizes[0], crateSizes)
		packedCrates = append(packedCrates, newCrate)

		// Remove sizes in newCrate from crateSizes
		updatedCrateSizes := slices.Clone(crateSizes)
		for _, newCrateVal := range newCrate {
			updatedCrateSizes, _ = common.RemoveValFromIntegerArray(updatedCrateSizes, newCrateVal)
		}
		crateSizes = slices.Clone(updatedCrateSizes)
	}

	return len(packedCrates)

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
