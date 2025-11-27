package main

import "fmt"

func normal(nums []int) int {
	for i := 0; i < len(nums); i++ {
		navNum := nums[i]
		j := i + 1
		for ; j < len(nums) && navNum == nums[j]; j++ {

		}
		if j < len(nums) {
			nums = append(nums[:i], nums[j-1:]...)
		} else {
			nums = nums[:i+1]
		}
	}

	return len(nums)
}

func removeDuplicates(nums []int) int {
	return normal(nums)
}

func slowFast(nums []int) int {
	slow, fast := 0, 0
	length := len(nums)
	for ; fast < length; fast++ {
		
	}
}

func main() {
	res := removeDuplicates([]int{2, 2})
	fmt.Println(res)
}
