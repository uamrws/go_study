package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in   chan int
	done func()
}

func doWork(i int, w worker) {
	for n := range w.in {
		fmt.Printf("worker%d print %c\n", i, n)
		w.done()
	}
}
func createWorker(i int, wg *sync.WaitGroup) worker {
	w := worker{in: make(chan int, 1), done: wg.Done}
	go doWork(i, w)
	return w
}

func main() {
	workers := new([10]worker)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}
	//wg.Add(20)
	for i, w := range workers {
		wg.Add(1)
		w.in <- 'a' + i
	}
	for i, w := range workers {
		wg.Add(1)
		w.in <- 'A' + i
	}
	wg.Wait()
}
