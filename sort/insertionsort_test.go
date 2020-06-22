package sort

import (
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	type args struct {
		set []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"test insertion sort",
			args{[]int{4, 5, 1, 6, 2, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InsertionSort(tt.args.set)
			fmt.Println(tt.args.set)
		})
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertionSort([]int{4, 5, 1, 6, 2, 3, 4, 5, 1, 6, 2, 3, 4, 5, 1, 6, 2, 3})
	}
}
