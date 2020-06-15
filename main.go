package main

import (
	"fmt"
)

func largestRectangleArea(heights []int) int {
	max := func(n1, n2 int) int {
		if n1 >= n2 {
			return n1
		}

		return n2
	}

	min := func(n1, n2 int) int {
		if n1 <= n2 {
			return n1
		}

		return n2
	}

	area := 0

	for i := 0; i < len(heights); i++ {
		m := heights[i]
		for j := i; j < len(heights); j++ {
			m = min(m, heights[j])
			area = max(area, (j - i +1) *m)
		}
	}

	return area
}

func main() {
	fmt.Println(largestRectangleArea([]int{2,1,5,6,2,3}))
}
// 100, 1, 11, 1, 120, 111, 123, 1, -1, -100
// 120, 11, 120, 120, 123, 123, -1, 100, 100