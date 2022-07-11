package main

import (
	"fmt"
)

func generate(n int) <-chan string {
	out := make(chan string)
	go func() {
		//defer close(out)
		for i := 0; i < n; i++ {
			out <- fmt.Sprintf("%d message: %d", n-10, i)
		}
	}()
	return out
}

func main() {

	out := make(chan string)
	//runtime.Goexit()
	for i := 0; i < 10; i++ {
		go func(i int) {
			for j := 0; j < 10; j++ {
				out <- fmt.Sprintf("%d send message: %d", i, j)
			}
		}(i)
	}

	//time.Sleep(time.Second * 10)
	select {}
}
