package leetcode

import (
	"strconv"

	. "github.com/zhuangzhi/leetcode/util"
)

// https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/
// 150. 逆波兰表达式求值
func evalRPN(tokens []string) int {
	stack := NewIntStack(len(tokens))
	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			l2 := stack.Top()
			stack.Pop()
			l1 := stack.Top()
			stack.Pop()
			v := calcStr(l1, l2, token)
			stack.Push(v)
		default:
			v, _ := strconv.Atoi(token)
			stack.Push(v)
		}
	}
	return stack.Top()
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
