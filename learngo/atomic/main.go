package main

import (
	"fmt"
	"sync"
)

type atomicInt struct {
	value int
	lock  sync.Mutex
}

func (a *atomicInt) increment() {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.value++
}
func (a *atomicInt) get() int {
	return a.value
}
func worker(a *atomicInt, wg *sync.WaitGroup) {
	for i := 0; i < 100000; i++ {
		a.increment()
	}
	wg.Done()
}
func main() {
	var a atomicInt
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go worker(&a, wg)
	go worker(&a, wg)
	wg.Wait()

	fmt.Println(a.get())
}
