package main

import "fmt"

func isValidParentheses(s string) bool {
	stack := []byte(s)

	for _, alpha := range s {
		if alpha == "{" || alpha == "(" || alpha == "[":{
			stack = append(stack, alpha)
		}
	}

	return len(stack) == 0 //空的就说明符号合法
}

func main() {
	res := isValidParentheses("()")
	fmt.Println(res)
}
