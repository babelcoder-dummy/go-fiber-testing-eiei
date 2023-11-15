package demo

func Filter[T comparable](items []T, fn func(item T) bool) []T {
	result := make([]T, 0)

	for _, item := range items {
		if fn(item) {
			result = append(result, item)
		}
	}

	return result
}
