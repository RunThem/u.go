package u_go

// Keys retrieve all the existing keys of a map
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, len(m))

	idx := 0
	for k := range m {
		keys[idx] = k
		idx++
	}

	return keys
}

// Values retrieve all the existing values of a map
func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, len(m))

	idx := 0
	for _, v := range m {
		values[idx] = v
		idx++
	}

	return values
}
