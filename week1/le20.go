package leetcode

import (
	. "github.com/zhuangzhi/leetcode/util"
)

// https://leetcode-cn.com/problems/valid-parentheses/
// 20. 有效的括号

func isValid(s string) bool {
	stack := NewCharStack(len(s))
	// ring := NewRing(len(s))
	for _, c := range s {
		switch c {
		case '(', '[', '{':
			stack.Push(c)
		case ')':
			if stack.Empty() || stack.Top() != '(' {
				return false
			}
			stack.Pop()
		case ']':
			if stack.Empty() || stack.Top() != '[' {
				return false
			}
			stack.Pop()
		case '}':
			if stack.Empty() || stack.Top() != '{' {
				return false
			}
			stack.Pop()
		}
	}
	return stack.Empty()
}

func isValidII(s string) bool {
	stack := NewCharStack(len(s))
	// ring := NewRing(len(s))
	match := [255]rune{}
	match[')'] = '('
	match[']'] = '['
	match['}'] = '{'

	for _, c := range s {
		switch c {
		case '(', '[', '{':
			stack.Push(c)
		case ')', ']', '}':
			if stack.Empty() || stack.Top() != match[c] {
				return false
			}
			stack.Pop()
		}
	}
	return stack.Empty()
}
