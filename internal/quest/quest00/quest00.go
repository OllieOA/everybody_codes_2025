package quest00

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
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

func getInput(ctx questregistry.Context) (string, error) {
	var content string
	var err error

	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("failed to determine current file path")
	}
	parentDir := filepath.Dir(currentFile)

	if ctx.Sample {
		fmt.Printf("Running Quest %02d with sample input.\n", questNumber)
		content, err = readFile(filepath.Join(parentDir, "sample.txt"))
		if err != nil {
			return "", fmt.Errorf("failed to read sample file: %w", err)
		}

	} else {
		fmt.Printf("Running Quest %02d with real input.\n", questNumber)
		content, err = readFile(filepath.Join(parentDir, "input.txt"))
		if err != nil {
			return "", fmt.Errorf("failed to read input file: %w", err)
		}
	}
	return content, nil
}

func parseContent(content string) (string, error) {
	// Placeholder for parsing logic
	// Used in the case where there is a shared transformation needed
	return content, nil
}

// Logical answers to the parts

func part1(content string) int {
	_ = content // placeholder to avoid unused variable error
	return 0
}

func part2(content string) int {
	_ = content // placeholder to avoid unused variable error
	return 0
}

func part3(content string) int {
	_ = content // placeholder to avoid unused variable error
	return 0
}

func Run(ctx questregistry.Context) error {
	var content string
	var err error
	base_start := time.Now()

	content, err = getInput(ctx)
	if err != nil {
		return err
	}

	io_complete_time := time.Since(base_start)
	logic_start := time.Now()

	parsed_content, err := parseContent(content)
	if err != nil {
		return fmt.Errorf("failed to parse content: %w", err)
	}

	part1_answer := part1(parsed_content)
	part1_time := time.Since(logic_start)
	fmt.Printf("Part 1 answer: %d (computed in %s)\n", part1_answer, part1_time)

	part2_answer := part2(parsed_content)
	part2_time := time.Since(logic_start) - part1_time
	fmt.Printf("Part 2 answer: %d (computed in %s)\n", part2_answer, part2_time)

	part3_answer := part3(parsed_content)
	part3_time := time.Since(logic_start) - part1_time - part2_time
	fmt.Printf("Part 3 answer: %d (computed in %s)\n", part3_answer, part3_time)

	total_time := time.Since(base_start)
	fmt.Printf("I/O time: %s, Total time: %s\n", io_complete_time, total_time)

	return nil
}
