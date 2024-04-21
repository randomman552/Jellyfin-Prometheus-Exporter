package collectors

// GroupByProperty groups a slice of structs by a specific property.
func GroupByProperty[T any, K comparable](items []T, getProperty func(T) K) map[K][]T {
	grouped := make(map[K][]T)

	for _, item := range items {
		key := getProperty(item)
		grouped[key] = append(grouped[key], item)
	}

	return grouped
}

// Filter the given array with the given predicate function
func Filter[T any](items []T, filterFunc func(T) bool) []T {
	result := []T{}

	for _, item := range items {
		if filterFunc(item) {
			result = append(result, item)
		}
	}

	return result
}

// Map the array of values using the given map function
func Map[T any, K any](items []T, mapFunc func(T) K) []K {
	result := []K{}

	for _, item := range items {
		result = append(result, mapFunc(item))
	}

	return result
}

// Map the array of values and flatten the result
func FlatMap[T any, K any](items []T, mapFunc func(T) []K) []K {
	result := []K{}

	for _, item := range items {
		result = append(result, mapFunc(item)...)
	}

	return result
}
