package week3

// edges = [[0,1][0,2]] -> 2
// 把图以任意一点为头变成树
func treeDiameter(edges [][]int) int {
	to := map[int][]int{} // to points, can be replaced by [][]int
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		to[x] = append(to[x], y)
		to[y] = append(to[y], x)
	}
	p := findFarthest(0, to)
	return findFarthest(p[0], to)[1]
}

func findFarthest(start int, to map[int][]int) [2]int { // 点，距离

	queue := make([]int, 0, 128)
	queue = append(queue, start)
	depth := map[int]int{} // with [n]int
	depth[start] = 0
	for len(queue) > 0 {
		children := make([]int, 0, 32)
		for _, x := range queue {
			for _, y := range to[x] {
				if _, ok := depth[y]; ok {
					continue
				}
				depth[y] = depth[x] + 1 //和父亲的距离是一
			}
			children = append(children, to[x]...)
		}
		queue = append(queue[:0], children...)
	}
	ans := start
	for p, dep := range depth {
		if dep > depth[ans] {
			ans = p
		}
	}
	return [2]int{ans, depth[ans]}
}
