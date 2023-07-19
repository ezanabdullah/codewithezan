package main

import (
	"fmt"
	"sync"
)

func main() {
	//Channel buffering
	ch2 := make(chan int, 5)
	go func() {
		ch2 <- 1
		ch2 <- 2
		ch2 <- 3
		ch2 <- 4
		ch2 <- 5
	}()
	i := 0
	for val := range ch2 {
		fmt.Println("val is", val)
		if i == 4 {
			break
		} else {
			i++
		}
	}
	//partial Synchronization
	var ch chan int = make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go SendNumber(ch, &wg)
	go RecieveNumber(ch, &wg)
	wg.Wait()
}
func SendNumber(ch chan int, wg *sync.WaitGroup) {
	for i := 1; i <= 10; i++ {
		ch <- i

	}
	close(ch)
}
func RecieveNumber(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range ch {
		fmt.Println("val's is", val)
	}
}
