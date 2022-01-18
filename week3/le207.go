package week3

// https://leetcode-cn.com/problems/course-schedule/
// 207. 课程表
// 度数，入度，出度

func canFinish(numCourses int, prerequisites [][]int) bool {
	to := make([][]int, numCourses)
	inDegree := make([]int, numCourses)
	lessons := make([]int, 0, numCourses)
	for _, pre := range prerequisites {
		ai, bi := pre[0], pre[1]
		to[bi] = append(to[bi], ai)
		inDegree[ai]++
	}
	queue := make([]int, 0, numCourses)

	// 拓扑排序第一步：从零入度点作为起点
	for idx, de := range inDegree {
		if de == 0 {
			queue = append(queue, idx)
		}
	}

	for len(queue) > 0 {
		next := make([]int, 0, numCourses)
		for _, idx := range queue {
			lessons = append(lessons, idx)
			// 第二部：扩展一个点，周围的入度减1
			for _, y := range to[idx] {
				inDegree[y]-- // 入度减一
				// 第三部：入度为0，可以学习
				if inDegree[y] == 0 {
					next = append(next, y)
				}
			}
		}
		queue = append(queue[:0], next...)
	}
	return len(lessons) == numCourses
}
