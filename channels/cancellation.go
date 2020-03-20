package channels

import (
	"fmt"
	"math/rand"
	"time"
)

func cancellation() {
	ch := make(chan string, 1)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "paper"
		fmt.Println("employee: sent signal")
	}()

	tc := time.After(100 * time.Millisecond)

	select {
	case p := <-ch:
		fmt.Println("manager: recv signal", p)
	case t := <-tc:
		fmt.Println("manager timeout", t)
	}

	time.Sleep(time.Second)
	fmt.Println("------------------------------------")
}
