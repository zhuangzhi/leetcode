package leetcode

import (
	"strconv"
)

// https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/
// 150. 逆波兰表达式求值
func evalRPN(tokens []string) int {
	stack := intstack(make([]int, 0, len(tokens)))
	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			l2 := stack.top()
			stack.pop()
			l1 := stack.top()
			stack.pop()
			v := calcStr(l1, l2, token)
			stack.push(v)
		default:
			v, _ := strconv.Atoi(token)
			stack.push(v)
		}
	}
	return stack.top()
}

func calcStr(left, right int, op string) int {
	switch op {
	case "+":
		return left + right
	case "-":
		return left - right
	case "*":
		return left * right
	case "/":
		return left / right
	}
	return 0
}
