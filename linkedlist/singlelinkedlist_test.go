package linkedlist

import (
	"reflect"
	"testing"
)

func TestNewSingleLinkedList(t *testing.T) {
	tests := []struct {
		name string
		want *SingleLinkedList
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSingleLinkedList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSingleLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSingleLinkedList_ExistCycle(t *testing.T) {
	type fields struct {
		head   *SingleNode
		length uint
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SingleLinkedList{
				head:   tt.fields.head,
				length: tt.fields.length,
			}
			if got := s.ExistCycle(); got != tt.want {
				t.Errorf("ExistCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSingleLinkedList_FindByIndex(t *testing.T) {
	type fields struct {
		head   *SingleNode
		length uint
	}
	type args struct {
		index uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SingleNode
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SingleLinkedList{
				head:   tt.fields.head,
				length: tt.fields.length,
			}
			if got := s.FindByIndex(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSingleLinkedList_First(t *testing.T) {
	type fields struct {
		head   *SingleNode
		length uint
	}
	type args struct {
		v int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SingleLinkedList{
				head:   tt.fields.head,
				length: tt.fields.length,
			}
			if got := s.First(tt.args.v); got != tt.want {
				t.Errorf("First() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSingleLinkedList_InsertAfter(t *testing.T) {
	type fields struct {
		head   *SingleNode
		length uint
	}
	type args struct {
		p *SingleNode
		v int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SingleLinkedList{
				head:   tt.fields.head,
				length: tt.fields.length,
			}
			if got := s.InsertAfter(tt.args.p, tt.args.v); got != tt.want {
				t.Errorf("InsertAfter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSingleLinkedList_InsertBefore(t *testing.T) {
	type fields struct {
		head   *SingleNode
		length uint
	}
	type args struct {
		p *SingleNode
		v int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SingleLinkedList{
				head:   tt.fields.head,
				length: tt.fields.length,
			}
			if got := s.InsertBefore(tt.args.p, tt.args.v); got != tt.want {
				t.Errorf("InsertBefore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSingleLinkedList_Last(t *testing.T) {
	type fields struct {
		head   *SingleNode
		length uint
	}
	type args struct {
		v int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SingleLinkedList{
				head:   tt.fields.head,
				length: tt.fields.length,
			}
			if got := s.Last(tt.args.v); got != tt.want {
				t.Errorf("Last() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSingleLinkedList_Remove(t *testing.T) {
	type fields struct {
		head   *SingleNode
		length uint
	}
	type args struct {
		p *SingleNode
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SingleLinkedList{
				head:   tt.fields.head,
				length: tt.fields.length,
			}
			if got := s.Remove(tt.args.p); got != tt.want {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSingleLinkedList_Reverse(t *testing.T) {

	type fields struct {
		head   *SingleNode
		length uint
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			"test_reverse",
			fields{
				head:   nil,
				length: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SingleLinkedList{
				head:   tt.fields.head,
				length: tt.fields.length,
			}
			s.Print()
			s.Reverse()
			s.Print()
		})
	}
}
