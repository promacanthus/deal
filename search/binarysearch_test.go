package search

import "testing"

func TestBinarySearch(t *testing.T) {
	type args struct {
		list  []int
		value int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test binary search",
			args{
				list:  []int{1, 2, 3, 4, 5, 6, 7, 8},
				value: 2,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.list, tt.args.value); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchRecursive(t *testing.T) {
	type args struct {
		list  []int
		value int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test binary search recursive",
			args{
				list:  []int{1, 2, 3, 4, 5, 6, 7, 8},
				value: 9,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BinarySearchRecursive(tt.args.list, tt.args.value)
			if got != tt.want {
				t.Errorf("BinarySearchRecursive() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchFirst(t *testing.T) {
	type args struct {
		list  []int
		value int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test binary search first",
			args{
				list:  []int{1, 1, 2, 2, 3, 4, 5, 6, 7, 8, 8},
				value: 2,
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearchFirst(tt.args.list, tt.args.value); got != tt.want {
				t.Errorf("BinarySearchFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchLast(t *testing.T) {
	type args struct {
		list  []int
		value int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test binary search first",
			args{
				list:  []int{1, 1, 2, 2, 3, 4, 5, 6, 7, 8, 8},
				value: 8,
			},
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearchLast(tt.args.list, tt.args.value); got != tt.want {
				t.Errorf("BinarySearchLast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchLE(t *testing.T) {
	type args struct {
		list  []int
		value int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test binary search less or equal",
			args{
				list:  []int{3, 5, 6, 8, 9, 10},
				value: 7,
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearchLE(tt.args.list, tt.args.value); got != tt.want {
				t.Errorf("BinarySearchLE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchGE(t *testing.T) {
	type args struct {
		list  []int
		value int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test binary search greater or equal",
			args{
				list:  []int{3, 5, 6, 8, 9, 10},
				value: 7,
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearchGE(tt.args.list, tt.args.value); got != tt.want {
				t.Errorf("BinarySearchGE() = %v, want %v", got, tt.want)
			}
		})
	}
}
