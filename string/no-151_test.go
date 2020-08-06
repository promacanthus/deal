package string

import "testing"

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		// {
		// 	"1",
		// 	args{"the sky is blue"},
		// 	"blue is sky the",
		// },
		// {
		// 	"2",
		// 	args{"  hello world!  "},
		// 	"world! hello",
		// },
		// {
		// 	"3",
		// 	args{"a good   example"},
		// 	"example good a",
		// },
		{
			"4",
			args{" "},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWordsBuiltIn(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
