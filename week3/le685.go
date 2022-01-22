package week3

import (
	. "github.com/zhuangzhi/leetcode/util"
)

type unionFind struct {
	ancestor []int
}

func newUnionFind(n int) unionFind {
	ancestor := make([]int, n)
	for i := 0; i < n; i++ {
		ancestor[i] = i // 每个节点链接到自己
	}
	return unionFind{ancestor}
}

func (uf unionFind) find(x int) int {
	if uf.ancestor[x] != x {
		//如果存在祖先, 压缩祖先，直到root，因为可能是union过后没有merge过的
		uf.ancestor[x] = uf.find(uf.ancestor[x])
	}
	return uf.ancestor[x]
}

func (uf unionFind) union(from, to int) {
	uf.ancestor[uf.find(from)] = uf.find(to) // 合并祖先，把from祖先的祖先设成to的祖先
}

func find(edges [][]int) []int {
	n := MaxEdgeNumber(edges)
	uf := newUnionFind(n + 1)
	parent := make([]int, n+1)
	for i := range parent {
		parent[i] = i
	}
	var conflictEdge, cycleEdge []int
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		if parent[to] != to {
			conflictEdge = edge
		} else {
			parent[to] = from

			// 如果from和to有共同祖先，则from->to会导致环
			if uf.find(from) == uf.find(to) {
				//
				cycleEdge = edge
			} else {
				// 如果没有共同祖先，合并二者的祖先
				uf.union(from, to)
			}
		}
	}
	// 若不存在一个节点有两个父节点的情况，则附加的边一定导致环路出现
	if conflictEdge == nil {
		return cycleEdge
	}
	// conflictEdge[1] 有两个父节点，{其中之一}与其构成附加的边
	// 由于我们是按照 edges 的顺序连接的，若在访问到 conflictEdge 之前已经形成了环路，则附加的边在环上
	// 否则附加的边就是 conflictEdge
	if cycleEdge != nil {
		return []int{parent[conflictEdge[1]], conflictEdge[1]}
	}

	return conflictEdge
}

type graph struct {
	Edges    [][]int
	MaxPoint int
	In, Out  [][]int
	IdxMap   map[int]int
	father   []int
}

func (g *graph) Build() {
	n := 0
	idxMap := map[int]int{}
	for _, edge := range g.Edges {
		x, y := edge[0], edge[1]
		n = MaxInt(n, MaxInt(x, y))
	}
	g.MaxPoint = n
	for idx, edge := range g.Edges {
		x, y := edge[0], edge[1]
		idxMap[g.EdgeToInt(x, y)] = idx
	}
	g.IdxMap = idxMap

	to := make([][]int, n+1)
	in := make([][]int, n+1)
	g.father = make([]int, n+1)

	for _, edge := range g.Edges {
		x, y := edge[0], edge[1]
		to[x] = append(to[x], y) //出度
		in[y] = append(in[y], x) //入度
	}
	g.In, g.Out = in, to
}

func (g graph) EdgeToInt(x, y int) int {
	return x*(g.MaxPoint+1) + y
}

func (g graph) IntToEdge(val int) (x, y int) {
	n := g.MaxPoint + 1
	y = val % n
	x = val / n
	return
}

func (g graph) HasBackCycle(x, y int) bool {
	// There is a x->y edge find if x back to y.
	// dfs in reach y
	var dfs func(i int) bool
	dfs = func(i int) bool {
		for _, f := range g.In[i] {
			if f == x {
				return false
			}
			if f == y {
				return true
			}
			if dfs(f) {
				return true
			}
		}
		return false
	}
	return dfs(x)
}

func findRedundantDirectedConnection(edges [][]int) []int {
	g := graph{Edges: edges}
	g.Build()
	root := -1
	result := map[int]bool{}
	// 入度为0的点是根结点
	for y, v := range g.In {
		if len(v) == 0 && y != 0 {
			root = y
		}
		if len(v) > 1 {
			// 入度大于1
			alt := -1
			for _, x := range v {
				if g.HasBackCycle(x, y) {
					alt = x
					result[g.EdgeToInt(x, y)] = true
					break
				}
			}
			if alt < 0 {
				for _, x := range v {
					result[g.EdgeToInt(x, y)] = true
				}
			}
		}
	}
	visited := make([]bool, g.MaxPoint+1)
	var dfs func(x int) bool
	dfs = func(x int) bool {
		visited[x] = true
		for _, y := range g.Out[x] {
			if visited[y] {
				result[g.EdgeToInt(x, y)] = true
				return true
			}
			if dfs(y) {
				return true
			}
		}
		visited[x] = false
		return false
	}
	if root < 0 {
		root = 1
	}
	dfs(root)
	max := -1
	r := []int{}

	for v, _ := range result {
		idx := g.IdxMap[v]
		if idx > max {
			max = idx
			x, y := g.IntToEdge(v)
			r = []int{x, y}
		}
	}
	return r
}
