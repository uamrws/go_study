package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var rev int
	for ; x != 0; x /= 10 {
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}
		m := x % 10
		rev = rev*10 + m
	}
	//if rev < math.MinInt32 || rev > math.MaxInt32 {
	//	return 0
	//}
	return rev
}
func main() {
	fmt.Println(reverse(-9463847412))
	fmt.Println(-123 / 10)
}
