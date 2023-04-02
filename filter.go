package u_go

// Filter returns all the elements from the collection which satisfies the conditional logic of callback function
func Filter[T any](slice []T, fn func(T) bool) []T {
	res := make([]T, 0)

	for _, v := range slice {
		if fn(v) {
			res = append(res, v)
		}
	}

	return res
}

// FilterMap iterates over the elements of a collection and returns a new collection
// representing all the items which satisfies the criteria formulated in the callback function
func FilterMap[K comparable, V any](m map[K]V, fn func(V) bool) map[K]V {
	filter := map[K]V{}

	for k, v := range m {
		if fn(v) {
			filter[k] = v
		}

	}

	return filter
}
