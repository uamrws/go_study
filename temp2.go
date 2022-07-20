package main

import (
	"fmt"
	"sync"
	"time"
)

var cond = sync.NewCond(&sync.Mutex{})

func read(i int, wg *sync.WaitGroup) {
	func() {
		cond.L.Lock()
		fmt.Printf("%d get the lock \n", i)
		defer cond.L.Unlock()
		cond.Wait()
		fmt.Printf("%d release the lock \n", i)
	}()
	fmt.Printf("read %d\n", i)
	time.Sleep(time.Second * 1)
	wg.Done()
}
func write(wg *sync.WaitGroup) {
	time.Sleep(time.Second * 4)
	fmt.Printf("write \n")
	cond.Broadcast()

	wg.Done()
}
func main() {

	wg := sync.WaitGroup{}
	wg.Add(11)
	go write(&wg)
	for i := 0; i < 10; i++ {
		go read(i, &wg)
	}
	wg.Wait()

}
