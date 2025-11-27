package main

import "fmt"

func forloop(digits []int) []int {
	length := len(digits)
	for i := length - 1; i >= 0; i-- {
		modnum := digits[i]
		if modnum < 9 { // 定位到第一个不是9的元素+1,其后置0，8-9 -> 9-0 78->79
			digits[i]++
			for j := i + 1; j < length; j++ {
				digits[j] = 0
			}
			return digits
		}
	}

	digits = make([]int, length+1) // 数组内都是9的情况就新生成一个全零数组将第一位设置成1
	digits[0] = 1
	return digits
}

func op(digits *[]int, idx int) {
	modnum := (*digits)[idx]
	if modnum == 9 && idx == 0 { // 进位处理
		(*digits)[idx] = 0
		*digits = append([]int{1}, (*digits)[0:]...)
		return
	}

	if modnum == 9 && idx > 0 {
		(*digits)[idx] = 0
		op(digits, idx-1)
		return
	}

	if modnum < 9 {
		(*digits)[idx]++
	}
}

func plusOne(digits []int) []int {
	op(&digits, len(digits)-1)
	//digits = forloop(digits)
	return digits
}

func main() {
	//res := plusOne([]int{9, 9, 9})
	res := plusOne([]int{1, 0, 9})
	fmt.Println(res)
}
