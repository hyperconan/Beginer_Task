package main

import "fmt"

func twoSum(nums []int, target int) []int {
	targetMap := make(map[int]int, len(nums))
	for idx, num := range nums {
		if _, ok := targetMap[num]; ok {
			return []int{targetMap[num], idx}
		}
		targetMap[target-num] = idx
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}

	target := 9
	fmt.Println(twoSum(nums, target))
}
