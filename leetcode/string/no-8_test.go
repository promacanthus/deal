package string

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		// {
		// 	"1",
		// 	args{" ++1"},
		// 	0,
		// },
		// {
		// 	"2",
		// 	args{"4193 with words"},
		// 	4193,
		// },
		{
			"3",
			args{"   -42"},
			-42,
		},
		{
			"4",
			args{"words and 987"},
			0,
		},
		{
			"5",
			args{"-91283472332"},
			-2147483648,
		},
		{
			"6",
			args{"+1"},
			1,
		},
		{
			"7",
			args{"+"},
			0,
		},
		{
			"8",
			args{"+-2"},
			0,
		},
		{
			"9",
			args{"-+2"},
			0,
		},
		{
			"10",
			args{"2147483648"},
			2147483647,
		},
		{
			"11",
			args{"42"},
			42,
		},
		{
			"12",
			args{"21474836++"},
			21474836,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi2(tt.args.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
