package common

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
)

// getQuestNumber extracts the quest number from the package name or file path
func GetQuestNumber() int {
	// Get the caller's file path
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return -1
	}

	// Extract quest number from path like "/path/to/quest00/quest00.go"
	re := regexp.MustCompile(`quest(\d+)`)
	matches := re.FindStringSubmatch(filename)
	if len(matches) >= 2 {
		if num, err := strconv.Atoi(matches[1]); err == nil {
			return num
		}
	}

	return -1
}

// IO handling

func ReadFile(name string) (string, error) {
	content, err := os.ReadFile(name)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %w", name, err)
	}
	return string(content), nil
}

func GetInput(ctx Context, part int) (string, error) {
	// This function will attempt to read input files based on the context and part number
	// with fallbacks to non-numbered files if not found.
	var content string
	var err error

	var targetFileType string

	_, currentFile, _, ok := runtime.Caller(1)
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

	content, err = ReadFile(partFilePath)
	if err == nil {
		return content, nil
	}

	fallbackFile := targetFileType + ".txt"
	fallbackFilePath := filepath.Join(parentDir, fallbackFile)

	content, err = ReadFile(fallbackFilePath)
	if err != nil {
		return "", fmt.Errorf("failed to read input files for part %d: %w", part, err)
	}

	return content, nil
}