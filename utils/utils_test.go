package utils

import (
	// "fmt"
	"palindrome/alphabet"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		args args
		want bool
	}{
		{args: args{"ABA"}, want: true},
		{args: args{"ABBA"}, want: true},
		{args: args{"ABABA"}, want: true},
		{args: args{"ACACA"}, want: true},
		{args: args{"ACDCA"}, want: true},
		{args: args{"AC CA"}, want: true},
		{args: args{"A a A"}, want: true},
		{args: args{"A a A"}, want: true},

		{args: args{"88"}, want: true},
		{args: args{"QEQ"}, want: true},
		{args: args{"A3O3A"}, want: true},
		{args: args{"F4L4F"}, want: true},

		{args: args{"ABDCA"}, want: false},
		{args: args{"AB"}, want: false},
		{args: args{"ABC"}, want: false},
		{args: args{""}, want: false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := IsPalindrome(tt.args.text); got != tt.want {
				t.Errorf("isPalindrome(%v) = %v, want %v", tt.args.text, got, tt.want)
			}
		})
	}
}

func TestToBaseK(t *testing.T) {
	type args struct {
		numberBase10 int
		k            int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "9",
			args: args{
				9, 28,
			},
			want: "9",
		},
		{
			name: "HH",
			args: args{
				232, 28,
			},
			want: "88",
		},
		{
			name: "MGM",
			args: args{
				16961, 28,
			},
			want: "LHL",
		},
		{
			name: "YNY",
			args: args{
				20802, 28,
			},
			want: "QEQ",
		},
		{
			name: "DVKVD",
			args: args{
				2972792, 28,
			},
			want: "4NBN4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToBaseK(tt.args.numberBase10, tt.args.k); got != tt.want {
				t.Errorf("ToBaseK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsBase10NumberAPalindromeInBaseK(t *testing.T) {
	type args struct {
		numberBase10 int
		K            int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "0 is a palindrome in any base",
			args: args{
				numberBase10: 0,
				K:            28,
			},
			want: true,
		},
		{
			name: "single digits are palindromes in any base",
			args: args{
				numberBase10: 5,
				K:            19,
			},
			want: true,
		},
		{
			name: "20802 is a palindrome in base 28",
			args: args{
				numberBase10: 20802,
				K:            28,
			},
			want: true,
		},
		{
			name: "20802 is not a palindrome in base 27",
			args: args{
				numberBase10: 20802,
				K:            27,
			},
			want: false,
		},
		{
			name: "10499401 is a palindrome in base 28",
			args: args{
				numberBase10: 10499401,
				K:            28,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBase10NumberAPalindromeInBaseK(tt.args.numberBase10, alphabet.Alphabet{
				K: tt.args.K,
			}); got != tt.want {
				t.Errorf("IsBase10NumberAPalindromeInBaseK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsBase10NumberAPalindromeInBaseK_KnownNumbers(t *testing.T) {
	type args struct {
		numberBase10 int
		alphabet            alphabet.Alphabet
	}
	tests := []args{}

	internationalKnownNumbers := []int{
		0,1,2,3,4,5,6,7,8,9,11,22,252,616,757,838,919,10301,13031,15951,17871,65856,1197911,2287822,4385834,5475745,5549455,6278726,6639366,7368637,7573757,8663668,8737378,9392939,9466649,9827289,67166176,214171412,609808906,836040638,2132882312,2487997842,5364224635,8417447148,9058338509,2823881883282,2904768674092,2969036309692,3817559557183,3958336338593,4280234320824,4446399936444,4520609060254,4725972795274,4806859586084,5052248422505,5193025203915,5779540459775,6014294924106,6041548451406,570984989489075,824685424586428,872951606159278,963217686712369,987115383511789,1276149009416721,2533880440883352,3774273333724773,63240319091304236,295871648846178592,542001114445237365,752009074470900257,957001143341100759,
	}

	for _, number := range internationalKnownNumbers {
		tests = append(tests, args{alphabet:alphabet.InternationalAlphabet, numberBase10: number})
	}

	spanishKnownNumbers := []int{
		0,1,2,3,4,5,6,7,8,9,11,22,232,464,696,5775,6336,10401,15251,16961,20802,2972792,5695965,6231326,6601066,7003007,8788878,9324239,10499401,161040161,427777724,500595005,711757117,751585157,773939377,904171409,947141749,1538668351,5760110675,6175005716,7298118927,10549494501,11242524211,13240004231,13363136331,79124342197,89279097298,276220022672,396183381693,731212212137,2636029206362,3597871787953,4120874780214,4699722279964,5788955598875,6385093905836,8332192912338,8338331338338,9693965693969,54173144137145,319102262201913,525408828804525,548785939587845,629596242695926,850128939821058,918395696593819,5249970550799425,20210235053201202,314435021120534413,
	}

	for _, number := range spanishKnownNumbers {
		tests = append(tests, args{alphabet:alphabet.SpanishAlphabet, numberBase10: number})
	}

	for _, tt := range tests {
		// t.Run(fmt.Sprintf("%v", tt.numberBase10), func(t *testing.T) {
			IsBase10NumberAPalindromeInBaseK(tt.numberBase10, tt.alphabet)
		// })
	}
}