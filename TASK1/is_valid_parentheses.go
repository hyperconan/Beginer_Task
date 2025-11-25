package main

import "fmt"

func isValidParentheses(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	stack := []byte{}

	tail := -1
	for idx, _ := range s {
		alpha := s[idx] // byte _是rune-int32
		if alpha == '{' || alpha == '(' || alpha == '[' {
			stack = append(stack, alpha)
			tail++
		} else if alpha == '}' && tail > -1 && stack[tail] == '{' {
			stack = stack[:tail]
			tail--
		} else if alpha == ')' && tail > -1 && stack[tail] == '(' {
			stack = stack[:tail]
			tail--
		} else if alpha == ']' && tail > -1 && stack[tail] == '[' {
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
	res := isValidParentheses("(")
	fmt.Println(res)
}
