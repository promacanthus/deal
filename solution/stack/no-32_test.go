package stack

import "testing"

func Test_longestValidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"1",
			args{"(()"},
			2,
		},
		{
			"2",
			args{")()())"},
			4,
		},
		{
			"3",
			args{"()(()"},
			2,
		},
		{
			"4",
			args{"()(()"},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParenthesesDP(tt.args.s); got != tt.want {
				t.Errorf("longestValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
