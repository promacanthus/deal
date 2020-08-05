package graph

import (
	"container/list"
	"fmt"
)

// 无向图
type Graph struct {
	Vertexes []*list.List
	count    uint
}

// v表示图中顶点数量
func NewGraph(v uint) *Graph {
	graph := &Graph{
		Vertexes: make([]*list.List, v),
		count:    v,
	}

	for i := uint(0); i < v; i++ {
		graph.Vertexes[i] = list.New()
	}

	return graph
}

func (g *Graph) AddEdge(s, t int) {
	g.Vertexes[s].PushBack(t)
	g.Vertexes[t].PushBack(s)
}

// 从起点s到终点t的路径
// s和t是按层序遍历的顶点位置
func (g *Graph) BFS(s, t int) {
	if s == t {
		return
	}

	visited := make([]bool, g.count)
	visited[s] = true

	queue := make([]int, g.count)
	queue = append(queue, s)

	prev := make([]int, g.count)
	for i := 0; i < int(g.count); i++ {
		prev[i] = -1
	}

	for len(queue) != 0 {
		w := queue[0]
		queue = queue[1:]
		vertex := g.Vertexes[w]
		for next := vertex.Front(); next != nil; next = next.Next() {
			q := next.Value.(int)
			if !visited[q] {
				prev[q] = w
				if q == t {
					prints(prev, s, t)
					return
				}
				visited[q] = true
				queue = append(queue, q)
			}
		}
	}
}

func prints(prev []int, s int, t int) {
	if prev[t] != -1 && t != s {
		prints(prev, s, prev[t])
	}
	fmt.Printf("%d  ", t)
}

// 从起点s到终点t的路径
func (g *Graph) DFS(s, t int) {
	if s == t {
		return
	}

	visited := make([]bool, g.count)

	prev := make([]int, g.count)
	for i := 0; i < int(g.count); i++ {
		prev[i] = -1
	}

	dfs(g, &visited, &prev, s, t)

	prints(prev, s, t)
}

func dfs(g *Graph, visited *[]bool, prev *[]int, s, t int) {

	(*visited)[s] = true
	if s == t {
		return
	}

	for next := g.Vertexes[s].Front(); next != nil; next = next.Next() {
		q := next.Value.(int)
		if !(*visited)[q] {
			(*prev)[next.Value.(int)] = s
			dfs(g, visited, prev, next.Value.(int), t)
		}
	}
}
