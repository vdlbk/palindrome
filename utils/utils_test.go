package utils

import (
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
			if got := IsBase10NumberAPalindromeInBaseK(tt.args.numberBase10, tt.args.K); got != tt.want {
				t.Errorf("IsBase10NumberAPalindromeInBaseK() = %v, want %v", got, tt.want)
			}
		})
	}
}
