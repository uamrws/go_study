package main

import (
	"fmt"
	"time"
)

func generate() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i > -1; i++ {
			out <- i
		}
	}()
	return out
}

func main() {
	var c1, c2 = generate(), generate()
	tm := time.After(10 * time.Millisecond)
	for {
		select {
		case n := <-c1:
			fmt.Println("Received from c1", n)
		case n := <-c2:
			fmt.Println("Received from c2", n)
		case <-tm:
			fmt.Printf("Bye")
			return
		}
	}
}
