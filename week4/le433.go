package week4

// https://leetcode-cn.com/problems/minimum-genetic-mutation/
// 433. 最小基因变化
// 用bfs求层数
func minMutation(start string, end string, bank []string) int {

	depth := map[string]int{}
	hashBank := map[string]bool{}
	for _, b := range bank {
		hashBank[b] = true
	}
	if _, ok := hashBank[end]; ok {
		return -1
	}

	gen4 := []rune{'A', 'C', 'G', 'T'}

	queue := []string{start}
	for len(queue) > 0 {
		tmp := []string{}
		for _, s := range queue {
			for idx, c := range s {
				for _, gen := range gen4 {
					if gen != c {
						//
						ns := []rune(s)
						ns[idx] = gen
						nss := string(ns)
						// 如果没有在hashBank上就继续
						if _, ok := hashBank[nss]; !ok {
							continue
						}
						//如果depth里有就继续，每个点秩序访问一次，第一次就是最少层数
						if _, ok := depth[nss]; ok {
							continue
						}
						depth[nss] = depth[s] + 1
						tmp = append(tmp, nss)
						if nss == end {
							return depth[nss]
						}
					}
				}
			}
		}
		queue = tmp
	}
	return -1
}
