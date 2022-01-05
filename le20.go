package leetcode

// https://leetcode-cn.com/problems/valid-parentheses/
// 20. 有效的括号

func isValid(s string) bool {
	ring := NewRing(len(s))
	for _, c := range s {
		switch c {
		case '(', '[', '{':
			ring.InsertLast(c)
		case ')':
			v := ring.GetRear()
			if v == nil || v.(rune) != '(' {
				return false
			}
			ring.DeleteLast()
		case ']':
			v := ring.GetRear()
			if v == nil || v.(rune) != '[' {
				return false
			}
			ring.DeleteLast()
		case '}':
			v := ring.GetRear()
			if v == nil || v.(rune) != '{' {
				return false
			}
			ring.DeleteLast()
		}
	}
	return ring.IsEmpty()
}
