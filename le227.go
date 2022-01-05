package leetcode

// https://leetcode-cn.com/problems/basic-calculator-ii/
// 227. 基本计算器 II

func calculateII(s string) int {
	s = s + " "
	tokens := strstack(make([]string, 0, (len(s)*2)/3))
	ops := charstack(make([]rune, 0, len(s)/2))
	number := make([]byte, 0, 64)
	for _, char := range s {
		// If is a number continue build number
		if char >= '0' && char <= '9' {
			number = append(number, byte(char))
			continue
		} else if len(number) != 0 {
			// if not and number and number not empty, append number
			tokens.push(string(number))
			number = number[:0]
		}
		// skip space.
		if char == ' ' {
			continue
		}

		// process opts
		currentRank := getRank(char)

		// MUST has '='
		for !ops.empty() && getRank(ops.top()) >= currentRank {
			tokens.push(string(ops.top()))
			ops.pop()
		}
		ops.push(char)
	}

	for len(ops) != 0 {
		tokens.push(string(ops.top()))
		ops.pop()
	}

	return evalRPN(tokens)
}

func getRank(r rune) int {
	switch r {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	}
	return 0 // error should panic
}
