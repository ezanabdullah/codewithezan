package main

import (
	"fmt"
	"sync"
)

func main() {
	mutex := sync.Mutex{}
	count := 0
	wg := sync.WaitGroup{}
	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go increment(&count, &wg, &mutex)
	}
	wg.Wait()
	fmt.Println("count is ", count)
}
func increment(balance *int, wg *sync.WaitGroup, locker sync.Locker) {
	defer wg.Done()
	locker.Lock()
	*balance++
	locker.Unlock()
}
