package common

import (
	"strconv"
)

func MakeIntMapJSONSerialisable[V any](inputMap map[int]V) map[string]V {
	outputMap := make(map[string]V, len(inputMap))
	for k, v := range inputMap {
		outputMap[strconv.Itoa(k)] = v
	}
	return outputMap
}