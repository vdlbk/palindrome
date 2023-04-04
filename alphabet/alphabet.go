package alphabet

const (
	Zero         int = '0'
	zeroFloat        = float64(Zero)
	LetterA      int = 'A'
	letterAFloat     = float64(LetterA)

	Charset                        = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	SpanishAlphabetAlphabeticOrder = "_ABCDEFGHIJKLMNÑOPQRSTUVWXYZ"
	SpanishAlphabetVowelEvenOrder  = "_BACDFEGHJIKLMNÑOPQRSTUVWXYZ"
	SpanishAlphabetMostCommonOrder = "_EAOSRNIDLCTUMPBGVYQHFZJÑXKW"
	SpanishAlphabetVowelFirstOrder = "_EAOIUNSDLCTRMPBGVYQHFZJÑXKW"
	SpanishAlphabetOjoOrder        = "_EAFIUNSDLCTRMPBGJYQHOZVÑXKW"

	InternationalAlphabetAlphabeticOrder = "_ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	InternationalAlphabetMostCommonOrder = "_EAISNRTOLUDCMPGBVHFQYXJKWZ"
)

var (
	SpanishAlphabet = Alphabet{
		K: len(SpanishAlphabetAlphabeticOrder),
		Variants: [][]rune{
			[]rune(SpanishAlphabetAlphabeticOrder),
			[]rune(SpanishAlphabetVowelEvenOrder),
			[]rune(SpanishAlphabetMostCommonOrder),
			[]rune(SpanishAlphabetVowelFirstOrder),
			[]rune(SpanishAlphabetOjoOrder),
		},
	}

	InternationalAlphabet = Alphabet{
		K: len(InternationalAlphabetAlphabeticOrder),
		Variants: [][]rune{
			[]rune(InternationalAlphabetAlphabeticOrder),
			[]rune(InternationalAlphabetMostCommonOrder),
		},
	}
)

type Alphabet struct {
	K        int
	Variants [][]rune
}

func GenerateAlphabet(numberOfLetters float64) map[rune]float64 {
	var i float64 = 0
	alphabet := map[rune]float64{}
	for i = 0; i < numberOfLetters; i++ {
		if i < 10 {
			alphabet[rune(zeroFloat+i)] = i
		} else {
			alphabet[rune(letterAFloat+(i-10))] = i
		}
	}
	return alphabet
}

func GetAlphabet(listOfLetters []rune) (map[rune]int, map[int]rune) {
	alphabet := map[rune]int{}
	alphabetInverted := map[int]rune{}
	for idx, char := range listOfLetters {
		alphabet[char] = idx
		alphabetInverted[idx] = char
	}
	return alphabet, alphabetInverted
}

func MapToAlphabet(numberBaseK string, listOfLetters []rune) string {
	result := ""
	chars, _ := GetAlphabet([]rune(Charset))
	_, m := GetAlphabet(listOfLetters)
	for _, char := range numberBaseK {
		result += string(m[chars[char]])
	}
	return result
}
