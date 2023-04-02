package main

import (
	"fmt"
	"palindrome/alphabet"
	"palindrome/utils"
	"time"
)

const (
	K                  = 28
	StartingNumber     = 0
	NumberOfPalindrome = 1_000_000_000
)

func main() {
	start := time.Now()

	//generatePalindromes(testPalindromicNumber, K, alphabet.SpanishAlphabets)
	generatePalindromes(testPalindromicNumber, 27, alphabet.FrenchAlphabets)

	fmt.Println("took", time.Since(start))
}

func testPalindromicNumber(number int, k int, alphabets [][]rune) bool {
	return utils.IsBase10NumberAPalindromeInBaseK(number, k, alphabets)
}

func generatePalindromes(test func(number int, k int, alphabets [][]rune) bool, k int, alphabets [][]rune) {
	for number := StartingNumber; number <= NumberOfPalindrome; number++ {
		temp := number
		reversedNumber := 0
		for temp > 0 {
			rem := temp % 10
			temp = temp / 10
			reversedNumber = (reversedNumber * 10) + rem
		}
		if number == reversedNumber {
			go test(number, k, alphabets)
		}
	}
}
