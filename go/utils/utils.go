package utils

func TernararyHelper[T any](callback func() bool, happy T, sad T) T {
	if callback() {
		return happy
	}
	return sad
}
