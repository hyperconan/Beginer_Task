package main

import "fmt"

func isValidParentheses(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	stack := []byte{}
	pairMap := map[byte]byte{
		'}': '{',
		')': '(',
		']': '[',
	}

	tail := -1
	for idx, _ := range s {
		alpha := s[idx]                                 // byte _是rune-int32
		if tail > -1 && pairMap[alpha] == stack[tail] { // 右找左
			stack = stack[:tail]
			tail--
		} else {
			stack = append(stack, alpha)
			tail++
		}
	}

	return len(stack) == 0 //空的就说明符号合法
}

func main() {
	res := isValidParentheses("()[][]({(())})")
	fmt.Println(res)
}
