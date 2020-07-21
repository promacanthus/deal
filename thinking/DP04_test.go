package thinking

import "testing"

func Test_dp04(t *testing.T) {
	type args struct {
		i  int
		cw int
		cv int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"1",
			args{
				i:  0,
				cw: 0,
				cv: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dp04(tt.args.i, tt.args.cw, tt.args.cv)
		})
	}
}
