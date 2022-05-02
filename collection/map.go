package collection

func MapValues[K comparable, V any](inputMap map[K]V) []V {
	values := make([]V, 0, len(inputMap))
	for _, v := range inputMap {
		values = append(values, v)
	}
	return values
}

func MapKeys[K comparable, V any](inputMap map[K]V) []K {
	keys := make([]K, 0, len(inputMap))
	for k := range inputMap {
		keys = append(keys, k)
	}
	return keys
}
