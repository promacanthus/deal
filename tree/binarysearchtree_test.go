package tree

import (
	"reflect"
	"testing"
)

func TestBinarySearchTree_Delete(t *testing.T) {
	type fields struct {
		root  *Node
		count uint
	}
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			"1",
			fields{
				root:  NewNode(1),
				count: 1,
			},
			args{value: 1},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BinarySearchTree{
				root:  tt.fields.root,
				count: tt.fields.count,
			}
			if got := b.Delete(tt.args.value); got != tt.want {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchTree_InOrder(t *testing.T) {
	tree := NewBinarySearchTree(NewNode(1))
	tree.root.left = NewNode(0)
	tree.root.right = NewNode(2)

	type fields struct {
		root  *Node
		count uint
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		// TODO: Add test cases.
		{
			"1",
			fields{
				root:  tree.root,
				count: 3,
			},
			[]int{0, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BinarySearchTree{
				root:  tt.fields.root,
				count: tt.fields.count,
			}
			if got := b.InOrder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchTree_Max(t *testing.T) {
	tree := NewBinarySearchTree(NewNode(1))
	tree.root.left = NewNode(0)
	tree.root.right = NewNode(2)
	type fields struct {
		root  *Node
		count uint
	}
	tests := []struct {
		name   string
		fields fields
		want   *Node
	}{
		// TODO: Add test cases.
		{
			"1",
			fields{
				root:  tree.root,
				count: 3,
			},
			tree.root.right,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BinarySearchTree{
				root:  tt.fields.root,
				count: tt.fields.count,
			}
			if got := b.Max(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchTree_Min(t *testing.T) {
	tree := NewBinarySearchTree(NewNode(1))
	tree.root.left = NewNode(0)
	tree.root.right = NewNode(2)
	type fields struct {
		root  *Node
		count uint
	}
	tests := []struct {
		name   string
		fields fields
		want   *Node
	}{
		// TODO: Add test cases.
		{
			"1",
			fields{
				root:  tree.root,
				count: 3,
			},
			tree.root.left,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BinarySearchTree{
				root:  tt.fields.root,
				count: tt.fields.count,
			}
			if got := b.Min(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchTree_PostOrder(t *testing.T) {
	tree := NewBinarySearchTree(NewNode(1))
	tree.root.left = NewNode(0)
	tree.root.right = NewNode(2)

	type fields struct {
		root  *Node
		count uint
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		// TODO: Add test cases.
		{
			"1",
			fields{
				root:  tree.root,
				count: 3,
			},
			[]int{0, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BinarySearchTree{
				root:  tt.fields.root,
				count: tt.fields.count,
			}
			if got := b.PostOrder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PostOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchTree_PreOrder(t *testing.T) {
	tree := NewBinarySearchTree(NewNode(1))
	tree.root.left = NewNode(0)
	tree.root.right = NewNode(2)

	type fields struct {
		root  *Node
		count uint
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		// TODO: Add test cases.
		{
			"1",
			fields{
				root:  tree.root,
				count: 3,
			},
			[]int{1, 0, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BinarySearchTree{
				root:  tt.fields.root,
				count: tt.fields.count,
			}
			if got := b.PreOrder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PreOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchTree_Search(t *testing.T) {
	tree := NewBinarySearchTree(NewNode(1))
	tree.root.left = NewNode(0)
	tree.root.right = NewNode(2)
	type fields struct {
		root  *Node
		count uint
	}
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			"1",
			fields{
				root:  tree.root,
				count: 3,
			},
			args{value: 1},
			true,
		},
		{
			"1",
			fields{
				root:  tree.root,
				count: 3,
			},
			args{value: 3},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BinarySearchTree{
				root:  tt.fields.root,
				count: tt.fields.count,
			}
			if got := b.Search(tt.args.value); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchTree_Insert(t *testing.T) {
	type fields struct {
		root  *Node
		count uint
	}
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			"1",
			fields{
				root:  NewNode(1),
				count: 1,
			},
			args{value: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BinarySearchTree{
				root:  tt.fields.root,
				count: tt.fields.count,
			}
			b.Insert(tt.args.value)
			b.PreOrder()
		})
	}
}
