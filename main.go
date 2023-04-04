package main

import (
	"fmt"
	"palindrome/alphabet"
	"palindrome/utils"
	"time"
)

const (
	StartingNumber     = 6_000_000_000_00
	NumberOfPalindrome = 10_000_000_000_00
	MidPoint           = 1000000
)

func main() {
	start := time.Now()

	generatePalindromes(testPalindromicNumber, alphabet.SpanishAlphabet)
	//generatePalindromes(testPalindromicNumber, alphabet.InternationalAlphabet)

	fmt.Println("took", time.Since(start))
}

func testPalindromicNumber(number int, k int, alphabets [][]rune) bool {
	return utils.IsBase10NumberAPalindromeInBaseK(number, k, alphabets)
}

func generatePalindromes(test func(number int, k int, alphabets [][]rune) bool, alphabet alphabet.Alphabet) {
	for number := StartingNumber; number <= NumberOfPalindrome; number++ {
		temp := number
		reversedNumber := 0
		for temp > 0 {
			rem := temp % 10
			temp = temp / 10
			reversedNumber = (reversedNumber * 10) + rem
		}
		if number == reversedNumber {
			go test(number, alphabet.K, alphabet.Variants)
		}
	}
}
