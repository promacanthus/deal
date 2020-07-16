package stringMatch

import "testing"

func TestBruteForce(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			"1",
			args{
				s: "abcdef",
				p: "cd",
			},
			true,
		},
		{
			"2",
			args{
				s: "akfduff",
				p: "rr",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BruteForce(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("BruteForce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRabinKarp(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			"1",
			args{
				s: "abcdef",
				p: "cd",
			},
			true,
		},
		{
			"2",
			args{
				s: "akfduff",
				p: "rr",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RabinKarp(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("RabinKarp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImprovedRabinKarp(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			"1",
			args{
				s: "abcdef",
				p: "cd",
			},
			true,
		},
		{
			"2",
			args{
				s: "akfduff",
				p: "rr",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ImprovedRabinKarp(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("ImprovedRabinKarp() = %v, want %v", got, tt.want)
			}
		})
	}
}
