package alphabet

import "testing"

func TestMapToAlphabet(t *testing.T) {
	type args struct {
		s             string
		listOfLetters []rune
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				s:             "F41E14F",
				listOfLetters: []rune(SpanishAlphabetAlphabeticOrder),
			},
			want: "ÑDANADÑ",
		},
		{
			name: "",
			args: args{
				s:             "QEQ",
				listOfLetters: []rune(SpanishAlphabetAlphabeticOrder),
			},
			want: "YNY",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapToAlphabet(tt.args.s, tt.args.listOfLetters); got != tt.want {
				t.Errorf("MapToAlphabet() = %v, want %v", got, tt.want)
			}
		})
	}
}
