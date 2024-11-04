package utils

func TernararyHelper[T any](callback func() bool, happy T, sad T) T {
	if callback() {
		return happy
	}
	return sad
}

func IsDigit(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	return false
}

func IsAlpha(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '_'
}

func IsAlphaNumeric(r rune) bool {
	return IsAlpha(r) || IsDigit(r)
}

func StrEndsWith(str string, pattern string) bool {
	if len(pattern) > len(str) {
		return false
	}
	for i, j := len(str), len(pattern); i >= 0 && j >= 0; i-- {
		if pattern[j] != str[i] {
			return false
		}
		j--
	}
	return true
}
