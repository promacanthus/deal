package thinking

import "testing"

func Test_dp03(t *testing.T) {
	type args struct {
		items []int
		n     int
		w     int
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
				items: []int{2, 2, 4, 6, 3},
				n:     5,
				w:     9,
			},
			9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dp03(tt.args.items, tt.args.n, tt.args.w); got != tt.want {
				t.Errorf("dp03() = %v, want %v", got, tt.want)
			}
		})
	}
}
