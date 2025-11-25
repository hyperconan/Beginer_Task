package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	x_str := strconv.Itoa(x)
	L := 0
	R := len(x_str) - 1
	tag := true
	for L < R && L < len(x_str) && R >= 0 {
		if x_str[L] != x_str[R] {
			tag = false
			break
		}
		L++
		R--
	}
	return tag
}

func isPalidromeByRevert(x int) bool { // 省去额外的空间申请

	if x < 0 || (x%10 == 0 && x != 0) { // 负数和末尾为0的数不是回文数,只有0是回文数
		return false
	}

	revert := 0
	for x > revert {
		revert = revert*10 + x%10
		x /= 10
	}
	isEvenPalidrome := x == revert   // 调转的一半数等于剩下的一半数
	isOddPalidrome := x == revert/10 // 调转的一半数的前两位等于剩下的一半数
	return isEvenPalidrome || isOddPalidrome
}

func main() {
	res := isPalidromeByRevert(101)
	fmt.Println(res)
}
