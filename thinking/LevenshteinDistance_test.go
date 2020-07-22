package thinking

import "testing"

func TestLD(t *testing.T) {
	type args struct {
		a string
		b string
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
				a: "mitcmu",
				b: "mtacnu",
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LDBT(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("LDBT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLDDP(t *testing.T) {
	type args struct {
		a string
		b string
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
				a: "mitcmu",
				b: "mtacnu",
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LDDP(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("LDDP() = %v, want %v", got, tt.want)
			}
		})
	}
}
