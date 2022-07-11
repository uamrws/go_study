package main

import "fmt"

func foo() *int {
	t := 3
	return &t
}

func main() {
	defer fmt.Println(1)
	defer fmt.Println(2)
}
