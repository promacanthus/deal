package DP

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			"0",
			args{
				s: "ab",
				p: ".*c",
			},
			false,
		},
		{
			"1",
			args{
				s: "aa",
				p: "a",
			},
			false,
		},
		{
			"2",
			args{
				s: "aa",
				p: "a*",
			},
			true,
		},
		{
			"3",
			args{
				s: "ab",
				p: ".*",
			},
			true,
		},
		{
			"4",
			args{
				s: "aab",
				p: "c*a*b",
			},
			true,
		},
		{
			"5",
			args{
				s: "mississippi",
				p: "mis*is*p*.",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
