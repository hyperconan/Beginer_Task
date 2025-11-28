package main

import "fmt"

/*
题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
考察点 ：指针运算、切片操作。
*/

func multiply2(sli *[]int) {
	// 将元素内的元素都乘以2
	for i := range *sli {
		(*sli)[i] *= 2
	}
}

func main() {
	nums := []int{1, 2, 3}
	multiply2(&nums)
	fmt.Println(nums)
}
