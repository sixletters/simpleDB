package util

func InsertAt[T any](slice []T, value T, index int) []T {
	if len(slice) == index { // nil or empty slice or after last element
		return append(slice, value)
	}
	slice = append(slice[:index+1], slice[index:]...) // index < len(a)
	slice[index] = value
	return slice
}
