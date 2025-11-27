package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	// 先排序，根据元素内的第一个元素从小到大排
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	left, right := intervals[0][0], intervals[0][1] // 第一个区间的左右边界
	for i := 1; i < len(intervals); i++ {
		rL, rR := intervals[i][0], intervals[i][1]
		if rL <= right && rR > right { // 交叉区间
			right = rR
		} else if rL > right { // 新区间的左边界 大于目前边界的最大值，则是新的区间
			res = append(res, []int{left, right}) // 记录目前区间
			left, right = rL, rR                  // 更新目前区间的左右边界
		}
	}
	res = append(res, []int{left, right}) // 记录最后一个区间
	return res
}

func main() {
	intervals := [][]int{{1, 4}, {4, 5}}
	fmt.Println(merge(intervals))
}
