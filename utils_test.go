package utils

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		args args
		want bool
	}{	
		{args:args{"ABA"},want:true},
		{args:args{"ABBA"},want:true},
		{args:args{"ABABA"},want:true},
		{args:args{"ACACA"},want:true},
		{args:args{"ACDCA"},want:true},
		{args:args{"AC CA"},want:true},
		{args:args{"A a A"},want:true},
		{args:args{"A a A"},want:true},

		{args:args{"ABDCA"},want:false},
		{args:args{"AB"},want:false},
		{args:args{"ABC"},want:false},
		{args:args{""},want:false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := isPalindrome(tt.args.text); got != tt.want {
				t.Errorf("isPalindrome(%v) = %v, want %v", tt.args.text, got, tt.want)
			}
		})
	}
}
