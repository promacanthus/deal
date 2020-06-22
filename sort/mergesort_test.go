package sort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test Merge Sort",
			args{list: []int{6, 5, 4, 3, 2, 1}},
			[]int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_merge(t *testing.T) {
	type args struct {
		list []int
		head int
		mid  int
		tail int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test merge",
			args{
				list: []int{4, 5, 6, 1, 2, 3},
				head: 0,
				mid:  2,
				tail: 5,
			},
			[]int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.list, tt.args.head, tt.args.mid, tt.args.tail); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeSort(t *testing.T) {
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
			"test merge sort",
			args{
				list: []int{6, 1, 3, 4, 2, 5},
				head: 0,
				tail: 5,
			},
			[]int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSort(tt.args.list, tt.args.head, tt.args.tail); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
