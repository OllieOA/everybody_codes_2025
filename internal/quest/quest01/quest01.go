package quest00

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/OllieOA/everybody_codes_2025/internal/quest_registry"
)

// Boilerplate

var questNumber int

func init() {
	// Automatically determine quest number from package name
	questNumber = getQuestNumber()
	questregistry.Register(questNumber, fmt.Sprintf("Quest %02d placeholder", questNumber), Run)
}

// getQuestNumber extracts the quest number from the package name or file path
func getQuestNumber() int {
	// Get the caller's file path
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return 0
	}

	// Extract quest number from path like "/path/to/quest00/quest00.go"
	re := regexp.MustCompile(`quest(\d+)`)
	matches := re.FindStringSubmatch(filename)
	if len(matches) >= 2 {
		if num, err := strconv.Atoi(matches[1]); err == nil {
			return num
		}
	}

	return 0
}

// IO handling

func readFile(name string) (string, error) {
	content, err := os.ReadFile(name)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %w", name, err)
	}
	return string(content), nil
}

func getInput(ctx questregistry.Context, part int) (string, error) {
	// This function will attempt to read input files based on the context and part number
	// with fallbacks to non-numbered files if not found.
	var content string
	var err error

	var targetFileType string

	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("failed to determine current file path")
	}
	parentDir := filepath.Dir(currentFile)

	if ctx.Sample {
		targetFileType = "sample"
	} else {
		targetFileType = "input"
	}

	partFile := targetFileType + strconv.Itoa(part) + ".txt"
	partFilePath := filepath.Join(parentDir, partFile)

	content, err = readFile(partFilePath)
	if err == nil {
		return content, nil
	}

	fallbackFile := targetFileType + ".txt"
	fallbackFilePath := filepath.Join(parentDir, fallbackFile)

	content, err = readFile(fallbackFilePath)
	if err != nil {
		return "", fmt.Errorf("failed to read input files for part %d: %w", part, err)
	}

	return content, nil
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

func part1(context questregistry.Context) string {
	part := 1
	content, err := getInput(context, part)
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

func part2(context questregistry.Context) string {
	part := 2
	content, err := getInput(context, part)
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

func part3(context questregistry.Context) string {
	part := 3
	content, err := getInput(context, part)
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

func Run(ctx questregistry.Context) error {
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
