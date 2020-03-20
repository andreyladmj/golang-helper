package channels

import (
	"fmt"
	"math/rand"
	"time"
)

func waitForResult() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "paper"
		fmt.Println("employee: sent signal")
	}()

	p := <-ch
	fmt.Println("manager: recv signal", p)

	time.Sleep(time.Second)
	fmt.Println("----------------------------------------------------")
}
