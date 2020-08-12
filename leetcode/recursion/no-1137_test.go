package recursion

import (
	"testing"
)

func Test_tribonacci(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"1",
			args{4},
			4,
		},
		{
			"2",
			args{25},
			1389537,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tribonacci(tt.args.n); got != tt.want {
				t.Errorf("tribonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tribonacciDP(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"1",
			args{4},
			4,
		},
		{
			"2",
			args{25},
			1389537,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tribonacciDP(tt.args.n); got != tt.want {
				t.Errorf("tribonacciIteration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_tribonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tribonacci(33)
	}
}
