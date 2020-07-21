package thinking

import "testing"

func TestDAC(t *testing.T) {
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
			"1",
			args{nums: []int{1, 5, 6, 2, 3, 4}},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DAC(tt.args.nums); got != tt.want {
				t.Errorf("DAC() = %v, want %v", got, tt.want)
			}
		})
	}
}
