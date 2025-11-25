package main

import (
	"fmt"
	"math"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	L, R := 0, 1
	minLen := math.MaxInt
	for _, str := range strs {
		minLen = int(math.Min(float64(len(str)), float64(minLen)))
	}
	pr := ""
	stop_tag := false
	for R <= minLen && stop_tag != true {
		tmp_pr := (strs[0])[L:R]
		for idx := 1; idx < len(strs); idx++ {
			if tmp_pr == (strs[idx])[L:R] {
				continue
			} else {
				stop_tag = true
				break
			}
		}
		if stop_tag == false {
			pr = tmp_pr
			R++
		}

	}
	return pr

}

func main() {
	res := longestCommonPrefix([]string{"d", "d"})
	fmt.Println(res)
}
