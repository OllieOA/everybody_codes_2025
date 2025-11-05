package quest00

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
func ParseInstructions(content string) ([]string, []string) {
	strArray := strings.Split(content, "\n\n")
	namesString := strArray[0]
	instructionsString := strArray[1]

	nameList := strings.Split(namesString, ",")
	instructions := strings.Split(instructionsString, ",")
	return nameList, instructions
}

// Logical answers to the parts

func part1(context common.Context) string {
	part := 1
	content, err := common.GetInput(context, part)
	if err != nil {
		fmt.Printf("No input for part %d: %s\n", part, err)
		return ""
	}

	// Logic below
	// Simple left/right movement without wrap-around
	nameList, instructions := ParseInstructions(content)

	cursorPosition := 0
	for _, instr := range instructions {
		instrDir := string(instr[0])
		instrSteps, _ := strconv.Atoi(instr[1:])

		var newCursorPosition int
		switch instrDir {
		case "R":
			newCursorPosition = cursorPosition + instrSteps
		case "L":
			newCursorPosition = cursorPosition - instrSteps
		}

		cursorPosition = min(newCursorPosition, len(nameList)-1)
		cursorPosition = max(cursorPosition, 0)
	}

	finalName := nameList[cursorPosition]
	return finalName
}

func part2(context common.Context) string {
	part := 2
	content, err := common.GetInput(context, part)
	if err != nil {
		fmt.Printf("No input for part %d: %s\n", part, err)
		return ""
	}

	// Logic below
	// Same as above but with wrap-around supported
	nameList, instructions := ParseInstructions(content)
	cursorPosition := 0
	for _, instr := range instructions {
		instrDir := string(instr[0])
		instrSteps, _ := strconv.Atoi(instr[1:])
		
		var newCursorPosition int
		switch instrDir {
		case "R":
			newCursorPosition = (cursorPosition + instrSteps) % len(nameList)
		case "L":
			newCursorPosition = (cursorPosition - instrSteps) % len(nameList)
			if newCursorPosition < 0 {
				newCursorPosition += len(nameList)
			}
		}

		cursorPosition = newCursorPosition
	}

	finalName := nameList[cursorPosition]
	return finalName
}

func part3(context common.Context) string {
	part := 3
	content, err := common.GetInput(context, part)
	if err != nil {
		fmt.Printf("No input for part %d: %s\n", part, err)
		return ""
	}

	// Logic below
	// Same as above but with wrap-around supported
	nameList, instructions := ParseInstructions(content)

	// Initialize the target indices with the current order. We will swap these around instead of 
	// names themselves for efficiency.
	targetIndices := make([]int, len(nameList))
	for i := range nameList {
		targetIndices[i] = i
	}

	for _, instr := range instructions {
		instrDir := string(instr[0])
		instrVal, _ := strconv.Atoi(instr[1:])

		cursorPosition := 0
		switch instrDir {
		case "R":
			cursorPosition = (cursorPosition + instrVal) % len(nameList)
		case "L":
			cursorPosition = (cursorPosition - instrVal) % len(nameList)
			if cursorPosition < 0 {
				cursorPosition += len(nameList)
			}
		}

		// Swap the target indices
		currTop := targetIndices[0]
		targetIndices[0] = targetIndices[cursorPosition]
		targetIndices[cursorPosition] = currTop
	}

	return nameList[targetIndices[0]]
}

// Finally, main runner function

func Run(ctx common.Context) error {
	base_start := time.Now()

	part1_answer := part1(ctx)
	part1_time := time.Since(base_start)
	fmt.Printf("Part 1 answer: %s (computed in %s)\n", part1_answer, part1_time)

	part2_answer := part2(ctx)
	part2_time := time.Since(base_start) - part1_time
	fmt.Printf("Part 2 answer: %s (computed in %s)\n", part2_answer, part2_time)

	part3_answer := part3(ctx)
	part3_time := time.Since(base_start) - part1_time - part2_time
	fmt.Printf("Part 3 answer: %s (computed in %s)\n", part3_answer, part3_time)

	total_time := time.Since(base_start)
	fmt.Printf("Total time: %s\n", total_time)

	return nil
}
