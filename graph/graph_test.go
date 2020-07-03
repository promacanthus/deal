package graph

import (
	"container/list"
	"testing"
)

func TestGraph_BFS(t *testing.T) {

	graph := NewGraph(8)
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 3)
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 5)
	graph.AddEdge(3, 4)
	graph.AddEdge(4, 5)
	graph.AddEdge(4, 6)
	graph.AddEdge(5, 7)
	graph.AddEdge(6, 7)

	type fields struct {
		Vertexes []*list.List
		count    uint
	}
	type args struct {
		s int
		t int
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
				Vertexes: graph.Vertexes,
				count:    8,
			},
			args{
				s: 0,
				t: 7,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				Vertexes: tt.fields.Vertexes,
				count:    tt.fields.count,
			}
			g.BFS(tt.args.s, tt.args.t)
		})
	}
}

func TestGraph_DFS(t *testing.T) {

	graph := NewGraph(8)
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 3)
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 5)
	graph.AddEdge(3, 4)
	graph.AddEdge(4, 5)
	graph.AddEdge(4, 6)
	graph.AddEdge(5, 7)
	graph.AddEdge(6, 7)

	type fields struct {
		Vertexes []*list.List
		count    uint
	}
	type args struct {
		s int
		t int
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
				Vertexes: graph.Vertexes,
				count:    8,
			},
			args{
				s: 0,
				t: 7,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				Vertexes: tt.fields.Vertexes,
				count:    tt.fields.count,
			}
			g.DFS(tt.args.s, tt.args.t)
		})
	}
}
