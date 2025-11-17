package common

import (
	"fmt"
	"slices"
)

func FindMaxIntegerArray(a []int) (int, error) {
	if len(a) == 0 {
		return 0, fmt.Errorf("empty array received")
	}

	maxVal := a[0]
	for i := range len(a) {
		if i == 0 {
			continue
		}
		if a[i] > maxVal {
			maxVal = a[i]
		}
	}

	return maxVal, nil
}

func FindUniqueValuesInIntegerArray(a []int) ([]int, error) {
	if len(a) == 0 {
		return []int{}, fmt.Errorf("empty array received")
	}

	var uniqueVals []int

	for _, s := range a {
		if slices.Contains(uniqueVals, s) {
			continue
		}
		uniqueVals = append(uniqueVals, s)
	}

	return uniqueVals, nil
}

func RemoveValFromIntegerArray(a []int, target int) ([]int, error) {
	if len(a) == 0 {
		return []int{}, fmt.Errorf("empty array received")
	}
	
	if !slices.Contains(a, target) {
		return a, nil
	}

    // Find first occurrence and remove it
    for i, v := range a {
        if v == target {
            result := make([]int, 0, len(a)-1)
            result = append(result, a[:i]...)
            result = append(result, a[i+1:]...)
            return result, nil
        }
    }
	return a, nil
}

func GetAllUniquePairs(a, b []int) [][2]int {
	out := [][2]int{}
	for _, i := range a {
		for _, j := range b {
			out = append(out, [2]int{i, j})
		}
	}
	return out
}