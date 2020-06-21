package queue

import (
	"container/list"
	"reflect"
	"testing"
)

func TestListQueue_Dequeue(t *testing.T) {
	type fields struct {
		items    *list.List
		capacity uint
	}
	tests := []struct {
		name   string
		fields fields
		want   int
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &ListQueue{
				items:    tt.fields.items,
				capacity: tt.fields.capacity,
			}
			got, got1 := l.Dequeue()
			if got != tt.want {
				t.Errorf("Dequeue() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Dequeue() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestListQueue_Enqueue(t *testing.T) {
	type fields struct {
		items    *list.List
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &ListQueue{
				items:    tt.fields.items,
				capacity: tt.fields.capacity,
			}
			if got := l.Enqueue(tt.args.n); got != tt.want {
				t.Errorf("Enqueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewListQueue(t *testing.T) {
	type args struct {
		capacity uint
	}
	tests := []struct {
		name string
		args args
		want *ListQueue
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewListQueue(tt.args.capacity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewListQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}
