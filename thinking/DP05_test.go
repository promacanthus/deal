package thinking

import "testing"

func Test_dp05(t *testing.T) {
	type args struct {
		weight []int
		value  []int
		n      int
		w      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"1",
			args{
				weight: []int{2, 2, 4, 6, 3},
				value:  []int{3, 4, 8, 9, 6},
				n:      5,
				w:      9,
			},
			13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dp05(tt.args.weight, tt.args.value, tt.args.n, tt.args.w); got != tt.want {
				t.Errorf("dp05() = %v, want %v", got, tt.want)
			}
		})
	}
}
