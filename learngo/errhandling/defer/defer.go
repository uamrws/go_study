package main

import (
	"fmt"
)

func main() {
	a := make([]int, 0, 3)
	_ = append(a, 1)
	defer fmt.Println(a[:3])
	_ = append(a, 2)
	defer fmt.Println(a[:3])
	_ = append(a, 3, 4)
	defer fmt.Println(a[:3])
}
