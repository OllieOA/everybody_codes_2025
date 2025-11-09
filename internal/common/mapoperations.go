package common

func CloneMap[K comparable, V any](src map[K]V) map[K]V {
	if src == nil {
		panic("invalid input map")
	}

	dst := make(map[K]V, len(src))
	for k, v := range src {
		dst[k] = v
	}

	return dst
}