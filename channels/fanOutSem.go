package channels

import (
	"fmt"
	"math/rand"
	"time"
)

func fanOutSem() {
	emps := 20
	ch := make(chan string, emps)

	const cap = 5
	sem := make(chan bool, cap)

	for e := 0; e < emps; e++ {
		go func(emp int) {
			sem <- true
			{
				time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
				ch <- "paper"
				fmt.Println("employee: sent signal", emp)
			}
			<-sem
		}(e)
	}

	for emps > 0 {
		p := <-ch
		fmt.Println(p)
		fmt.Println("manager:: recv signal", emps)
		emps--
	}
}
