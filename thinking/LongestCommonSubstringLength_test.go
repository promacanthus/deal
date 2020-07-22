package thinking

import "testing"

func TestLCSLBT(t *testing.T) {
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
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LCSLBT(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("LCSLBT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLCSLDP(t *testing.T) {
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
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LCSLDP(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("LCSLDP() = %v, want %v", got, tt.want)
			}
		})
	}
}
