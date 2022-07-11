package main

import (
	"fmt"
	fibo "go_study/learngo/testdemo/fibonacci"
)

func main() {
	fmt.Println(fibo.Fibonacci(50))
	fmt.Println(fibo.Fibonacci2(50))
}
