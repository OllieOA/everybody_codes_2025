package quest04

import (
	"fmt"
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

func getGearRatio(gearTeeth []int) float64 {
	gearRatio := 1.0
	for i, currGear := range gearTeeth {
		if i == 0 {
			continue
		}
		prevGear := gearTeeth[i-1]
		gearRatio *= float64(prevGear) / float64(currGear)
	}
	return gearRatio
}

func getGearRatioWithStacked(gearTeeth []string) float64 {
	gearRatio := 1.0
	for i, currGear := range gearTeeth {
		if i == 0 {
			continue
		}
		
		var currGearTurnValue int
		var prevGearTurnValue int

		if strings.Contains(currGear, "|") {
			currGearStacked := strings.Split(currGear, "|")
			currGearTurnValue, _ = strconv.Atoi(currGearStacked[0])
		} else {
			currGearTurnValue, _ = strconv.Atoi(currGear)
		}
		
		prevGear := gearTeeth[i-1]
		if strings.Contains(prevGear, "|") {
			prevGearStacked := strings.Split(prevGear, "|")
			prevGearTurnValue, _ = strconv.Atoi(prevGearStacked[1])
		} else {
			prevGearTurnValue, _ = strconv.Atoi(prevGear)
		}

		gearRatio *= float64(prevGearTurnValue) / float64(currGearTurnValue)
	}
	return gearRatio
}

// Logical answers to the parts

func part1(context common.Context) int {
	part := 1
	content, err := common.GetInput(context, part)
	if err != nil {
		fmt.Printf("No input for part %d: %s\n", part, err)
		return 0
	}

	firstTurns := 2025.0

	// Logic below
	gearTeeth, _ := common.ParseDelimtedList(content, strconv.Atoi, "\n")
	gearRatio := getGearRatio(gearTeeth)

	return int(firstTurns * gearRatio)
}

func part2(context common.Context) int {
	part := 2
	content, err := common.GetInput(context, part)
	if err != nil {
		fmt.Printf("No input for part %d: %s\n", part, err)
		return 0
	}

	// Logic below
	gearTeeth, _ := common.ParseDelimtedList(content, strconv.Atoi, "\n")
	gearRatio := getGearRatio(gearTeeth)

	lastTurns := 10000000000000.0

	return int(lastTurns / gearRatio) + 1
}

func part3(context common.Context) int {
	part := 3
	content, err := common.GetInput(context, part)
	if err != nil {
		fmt.Printf("No input for part %d: %s\n", part, err)
		return 0
	}

	// Logic below
	gearTeeth := strings.Split(content, "\n")
	gearRatio := getGearRatioWithStacked(gearTeeth)

	firstTurns := 100.0

	return int(firstTurns * gearRatio)
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
