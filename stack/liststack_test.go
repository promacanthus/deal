package stack

import (
	"container/list"
	"reflect"
	"testing"
)

func TestListStack_Pop(t *testing.T) {
	type fields struct {
		data     *list.List
		length   uint
		capacity uint
	}
	tests := []struct {
		name   string
		fields fields
		want   int
		want1  bool
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ListStack{
				items:    tt.fields.data,
				length:   tt.fields.length,
				capacity: tt.fields.capacity,
			}
			got, got1 := s.Pop()
			if got != tt.want {
				t.Errorf("Pop() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Pop() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestListStack_Push(t *testing.T) {
	type fields struct {
		data     *list.List
		length   uint
		capacity uint
	}
	type args struct {
		n int
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
			s := &ListStack{
				items:    tt.fields.data,
				length:   tt.fields.length,
				capacity: tt.fields.capacity,
			}
			if got := s.Push(tt.args.n); got != tt.want {
				t.Errorf("Push() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewListStack(t *testing.T) {
	type args struct {
		n uint
	}
	tests := []struct {
		name string
		args args
		want *ListStack
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewListStack(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewListStack() = %v, want %v", got, tt.want)
			}
		})
	}
}
