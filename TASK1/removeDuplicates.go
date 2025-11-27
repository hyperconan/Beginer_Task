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
	//return normal(nums)
	return slowFast(nums)
}

func slowFast(nums []int) int {
	//慢指针用于替换元素，快指针用于查找元素
	slow, fast := 0, 1
	length := len(nums)
	for ; fast < length; fast++ {
		refNum := nums[slow]
		if nums[fast] != refNum {
			slow++
			nums[slow] = nums[fast]
		}
	}
	nums = nums[:slow+1]
	return slow + 1
}

func main() {
	res := removeDuplicates([]int{1, 2, 2, 3, 3, 4})
	fmt.Println(res)
}
