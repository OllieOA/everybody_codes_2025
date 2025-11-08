package common

import (
	"fmt"
	"strings"
)

func ParseDelimtedList[T any](content string, parseMethod func(string) (T, error), delimiter string) ([]T, error) {
	if delimiter == "" {
		delimiter = ","
	}
	
	parts := strings.Split(content, delimiter)
    out := make([]T, 0, len(parts))
    for _, p := range parts {
        p = strings.TrimSpace(p)
        v, err := parseMethod(p)
        if err != nil {
            return nil, fmt.Errorf("parsing %q: %w", p, err)
        }
        out = append(out, v)
    }
    return out, nil
}