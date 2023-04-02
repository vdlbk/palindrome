package utils

func isPalindrome(text string) bool {
	left, right := 0, len(text)-1

	if right < left {
		return false
	}

	for left < right && left < len(text) {
		if text[left] != text[right] {
			return false
		}

		left++
		right--
	}

	return true
}