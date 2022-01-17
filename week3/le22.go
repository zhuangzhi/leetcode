package week3

// 第一段是不可分割的整体(A)B  (k-1) n-k
// 加法原理：k-1 + n-K, S = (A)B 乘法原理 A*B
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}
	if n == 1 {
		return []string{"()"}
	}
	s := []string{}
	perm := make([]byte, 0, n*2)
	for i := 0; i < n; i++ {
		s1 := generateParenthesis(i)
		s2 := generateParenthesis(n - i - 1)
		for _, x1 := range s1 {
			perm = append(perm, '(')
			perm = append(perm, x1...)
			perm = append(perm, ')')
			l := len(perm)
			for _, x2 := range s2 {
				perm = append(perm, x2...)
				s = append(s, string(perm))
				perm = perm[:l]
			}
			perm = perm[:0]
		}
	}
	return s
}
