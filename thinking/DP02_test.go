package thinking

import "testing"

func Test_dp02(t *testing.T) {
	type args struct {
		weight []int
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
				n:      5,
				w:      9},
			9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dp02(tt.args.weight, tt.args.n, tt.args.w); got != tt.want {
				t.Errorf("dp02() = %v, want %v", got, tt.want)
			}
		})
	}
}
