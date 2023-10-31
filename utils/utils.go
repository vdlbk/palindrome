package utils

import (
	"fmt"
	"math"
	"palindrome/alphabet"
	"strings"
)

func ToBase10(numberBaseK string, k float64, alphabet map[rune]float64) *float64 {
	length := len(numberBaseK)
	var result float64 = 0
	for idx, char := range numberBaseK {
		letter, exist := alphabet[char]
		if !exist {
			fmt.Printf("invalid input '%v' at index %d", char, idx)
			return nil
		}

		exp := float64(length - idx - 1)

		result += letter * math.Pow(k, exp)
	}
	return &result
}

func ToBaseK(numberBase10, k int) string {
	if numberBase10 == 0 {
		return "0"
	}

	var result string

	for numberBase10 > 0 {
		r := numberBase10 % k
		if r < 10 {
			result = string(rune(alphabet.Zero+r)) + result
		} else {
			result = string(rune(alphabet.LetterA+(r-10))) + result
		}
		numberBase10 = numberBase10 / k
	}

	return result
}

func IsPalindrome(text string) bool {
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

func IsBase10NumberAPalindromeInBaseK(numberBase10 int, currentAlphabet alphabet.Alphabet) bool {
	numberBaseK := ToBaseK(numberBase10, currentAlphabet.K)
	isPalindrome := IsPalindrome(numberBaseK)
	if isPalindrome {
		var words []string
		for _, runes := range currentAlphabet.Variants {
			words = append(words, alphabet.MapToAlphabet(numberBaseK, runes))
		}

		fmt.Printf("%s: | %s | %d | %s |\n", currentAlphabet.Name, numberBaseK, numberBase10, strings.Join(words, " | "))
	}
	return isPalindrome
}
