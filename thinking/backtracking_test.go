package thinking

import "testing"

func Test_cal8queens(t *testing.T) {
	type args struct {
		row int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"1",
			args{row: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cal8queens(tt.args.row)
		})
	}
}

func Test_packages(t *testing.T) {
	type args struct {
		i     int
		cw    int
		items []int
		n     int
		w     int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"1",
			args{
				i:     0,
				cw:    0,
				items: []int{2, 2, 4, 6, 3},
				n:     5,
				w:     9,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			packages(tt.args.i, tt.args.cw, tt.args.items, tt.args.n, tt.args.w)
		})
	}
}
