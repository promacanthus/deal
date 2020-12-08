package arrays

import "testing"

func Test_firstMissingPositive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"-1",
			args{[]int{1}},
			1,
		},
		{
			"0",
			args{[]int{4, 3, 4, 1, 1, 4, 1, 4}},
			2,
		},
		{
			"1",
			args{[]int{1, 2, 0}},
			3,
		},
		{
			"2",
			args{[]int{3, 4, -1, 1}},
			2,
		},
		{
			"3",
			args{[]int{7, 8, 9, 11, 12}},
			1,
		},
		{
			"4",
			args{[]int{}},
			1,
		},
		{
			"5",
			args{[]int{2}},
			1,
		},
		{
			"6",
			args{[]int{1}},
			2,
		},
		{
			"7",
			args{[]int{0}},
			1,
		},
		{
			"8",
			args{[]int{0, 1}},
			2,
		},
		{
			"9",
			args{[]int{1, 1}},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive(tt.args.nums); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
