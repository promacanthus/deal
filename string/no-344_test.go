package string

import "testing"

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"1",
			args{[]byte{'h', 'e', 'l', 'l', 'o'}},
		},
		{
			"2",
			args{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}},
		},
		{
			"3",
			args{[]byte{'a'}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.args.s)
		})
	}
}
