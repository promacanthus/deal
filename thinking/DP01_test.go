package thinking

import (
	"testing"
)

func Test_dp01(t *testing.T) {
	type args struct {
		i  int
		cw int
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
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dp01(tt.args.i, tt.args.cw)
		})
	}
}
