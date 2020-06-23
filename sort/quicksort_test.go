package sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test Quick Sort",
			args{
				[]int{4, 6, 3, 1, 5, 2},
			},
			[]int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_partition(t *testing.T) {
	type args struct {
		list []int
		head int
		tail int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test partition",
			args{
				list: []int{4, 6, 3, 1, 5, 2},
				head: 0,
				tail: 5,
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.list, tt.args.head, tt.args.tail); got != tt.want {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_quickSort(t *testing.T) {
	type args struct {
		list []int
		head int
		tail int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test quick Sort",
			args{
				list: []int{4, 6, 3, 1, 5, 2},
				head: 0,
				tail: 5,
			},
			[]int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quickSort(tt.args.list, tt.args.head, tt.args.tail); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
