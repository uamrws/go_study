package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func fibonacci() fibGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type fibGen func() int

func (f fibGen) Read(p []byte) (n int, err error) {
	next := f()
	s := fmt.Sprintf("%d\n", next)
	if next > 100 {
		return 0, io.EOF
	}
	// ToDo: incorrect if p is to small
	return strings.NewReader(s).Read(p)
}

func main() {

	fib := fibonacci()
	scanner := bufio.NewScanner(fib)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		println(i)
	}
}
