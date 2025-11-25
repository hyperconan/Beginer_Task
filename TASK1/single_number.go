package main

import "fmt"

func singleNumber(nums []int) int {
	countMap := make(map[int]int)
	for _, num := range nums {
		countMap[num]++
	}
	for num, count := range countMap {
		if count == 1 {
			return num
		}
	}
	return -3e5
}

func singleNumberByXor(nums []int) int {
	// 只想取出仅出现一次的数字用异或可以仅遍历一次获得
	// 任何数字异或0都等于本身，而任何数字异或本身都等于0，那么一个数字除了某一个数字只出现一次，其他数字都出现了两次,就可以用异或将重复两次的数字进行异或得到0
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}

func main() {
	res := singleNumber([]int{4, 1, 2, 2, 2})
	fmt.Println(res)
}
