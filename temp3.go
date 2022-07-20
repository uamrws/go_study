package main

import (
	"fmt"
	"sort"
)

func minimumAbsDifference(arr []int) (ans [][]int) {

	sort.Ints(arr)
	n := len(arr) - 1
	best := arr[n] - arr[0]
	for i := 0; i < n; i++ {
		a, b := arr[i], arr[i+1]
		delta := b - a
		if delta == best {
			ans = append(ans, []int{a, b})
		} else if delta < best {
			ans = [][]int{{a, b}}
			best = delta
		}
	}
	return ans
}
func main() {
	fmt.Println(minimumAbsDifference([]int{3, 8, -10, 23, 19, -4, -14, 27}))
}
