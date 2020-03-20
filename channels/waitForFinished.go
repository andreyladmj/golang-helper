package channels

import (
	"fmt"
	"math/rand"
	"time"
)

func waitForFinished() {
	ch := make(chan struct{})

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		close(ch)
		fmt.Println("employee: send signal")
	}()

	_, wd := <-ch // wd - with data
	fmt.Println("manager: recv signal", wd)

	time.Sleep(time.Second)
	fmt.Println("---------------------------------------------")
}
