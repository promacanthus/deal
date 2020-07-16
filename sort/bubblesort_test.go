package sort

import (
	"testing"
)

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort([]int{4, 5, 1, 6, 2, 3, 4, 5, 1, 6, 2, 3, 4, 5, 1, 6, 2, 3})
	}
}

func TestBubbleSort1(t *testing.T) {
	a := []int{4, 5, 1, 6, 2, 3}
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test bubble sort",
			args: args{a},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BubbleSort(tt.args.list)
		})
	}
}
