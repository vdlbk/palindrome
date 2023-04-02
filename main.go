package main

import (
	"fmt"
	"math"
	"math/rand"
	"palindrome/utils"
)

const (
	inputB28     string = "88"
	alphabetSize   float64     = 28
	alphabetSizeInt   int     = int(alphabetSize)
	zero float64 = float64('0')
	zeroInt int = '0'
	letterA float64 = float64('A')
	letterAInt int = 'A'
)

func main() {
	fmt.Println(inputB28)
	alphabet := getAlaphabet(alphabetSize)
	fmt.Println(alphabet)
	fmt.Println("last", generatePalindromes())

	base10 := toBase10(inputB28, alphabet)
	
	if base10 == nil {
		fmt.Println("failed to convert b28 to b10")
		return
	}
	fmt.Printf("%.0f\n", *base10)
	base28 := toBase28(1441)
	fmt.Printf("%s\n", base28, utils.isPalindrome(base28))
}

func toBase10(numberBase28 string, alphabet map[rune]float64) *float64 {
	length := len(numberBase28)
	var result  float64 = 0
	for idx, char := range inputB28 {
		letter, exist := alphabet[char]
		if !exist {
			fmt.Printf("invalid input '%v' at index %d", char, idx)
			return nil
		}

		exp := float64(length - idx - 1)

		result += letter * math.Pow(alphabetSize, exp)
	}
	return &result
}

func toBase28(numberBase10 int) string {
	if numberBase10 == 0 {
		return "0"
	}

	var result string

	for numberBase10 > 0 {
		r := numberBase10 % alphabetSizeInt
		if r < 10 {
			result = string(zeroInt+r) + result
		} else {
			result = string(letterAInt+(r-10)) + result
		}
		numberBase10 = numberBase10 / alphabetSizeInt
	}

	return result
}

func getAlaphabet(numberOfLetters float64) map[rune]float64 {
	var i float64 = 0
	alphabet := map[rune]float64{}
	for i = 0; i < numberOfLetters; i++ {
		if i < 10 {
			alphabet[rune(zero+i)] = i
		} else {
			alphabet[rune(letterA+(i-10))] = i
		}
	}
	return alphabet
}

func generatePalindromes() int {
	last := 0
	for num := 0; num <= 1_000_000; num++ {
		temp := num
		reverse_num := 0
		for temp > 0 {
			rem := temp % 10
			temp = temp / 10
			reverse_num = (reverse_num * 10) + rem
		}
		if num == reverse_num {
			last = num
			if rand.Float64() < 0.0009 {
				// fmt.Println(last)
			}
		}
	}
	return last
}



/*
+1 		0 1 2 3 4 5 6 7 8 9
+2

+11		11 22 33 44 55 66 77 88 99
+2

+10		101	111	121	131	141	151	161	171	181	191
+11
+10		202	212	222	232	242	252	262	272	282	292
+11
+10		303	313	323	333	343	353	363	373	383	393
...
+11
+10		909	919	929	939	949	959	969	979	989	999
+2

+110	1001 1111 1221 1331 1441 1551 1661 1771 1881 1991
+11
+110	2002 2112 2222 2332 2442 2552 2662 2772 2882 2992
...
+11
+110	9009 9119 9229 9339 9449 9559 9669 9779 9889 9999
+2
+100	10001 10101 10201 10301 10401 10501 10601 10701 10801 10901
		11011 11111	11211


*/
