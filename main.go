package main

import (
	"fmt"
	"palindrome/alphabet"
	"palindrome/utils"
	"strings"
	"sync"
	"time"
)

const (
	Base10 = 10
	EndingNumber   = 1_000_000
)

func main() {
	start := time.Now()

	generatePalindromes(EndingNumber, testPalindromicNumber, []alphabet.Alphabet{
		alphabet.SpanishAlphabet,
		alphabet.InternationalAlphabet,
	})

	fmt.Println("took", time.Since(start))
}

func testPalindromicNumber(number int, alphabet alphabet.Alphabet, waitGroup *sync.WaitGroup) bool {
	defer waitGroup.Done()
	return utils.IsBase10NumberAPalindromeInBaseK(number, alphabet)
}

func generatePalindromes(ending int, test func(int, alphabet.Alphabet, *sync.WaitGroup) bool, alphabets []alphabet.Alphabet) {
	for _, alphabet := range alphabets {
		variants := []string{}
		for _, variant := range alphabet.Variants {
			variants = append(variants, string(variant))
		}
		fmt.Printf("%s: | `%s` |\n", alphabet.Name, strings.Join(variants, "` | `"))
	}


	// odd and even
	waitGroup := sync.WaitGroup{}
	for j := 0; j < 2; j++ {
		i := 1
		for number := createPalindrome(i, j%2 == 0); number < ending; number = createPalindrome(i, j%2 == 0) {
			waitGroup.Add(len(alphabets))
			for _, alphabet := range alphabets {
				go test(number, alphabet, &waitGroup)
			}
			i++
		}
	}
	waitGroup.Wait()
}

func createPalindrome(input int, isOdd bool) int {
	n := input
	palindrome := input

	if isOdd {
		n /= Base10
	}

	for n > 0 {
		palindrome = (palindrome * Base10) + (n % Base10)
		n /= Base10
	}
	return palindrome
}