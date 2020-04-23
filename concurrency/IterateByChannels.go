package main

import (
	"fmt"
	"sync"
)

func main() {
	intStream := make(chan int)
	go func() {
		//defer close(intStream)
		for i := 1; i <= 5; i++ {
			intStream <- i
		}
	}()
	go func() {
		//defer close(intStream)
		for i := 10; i <= 15; i++ {
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Printf("%v ", integer)
	}
}

func BlockedGoRountines() {
	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Printf("%v has begun\n", i)
		}(i)
	}

	fmt.Println("Unblocking goroutines...")
	close(begin)
	wg.Wait()
}
