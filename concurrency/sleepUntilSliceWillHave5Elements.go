package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})
	//wg := sync.WaitGroup{}
	lines := make([]interface{}, 0, 10)
	startedGoRoutines := 0

	addLine := func(start, end int) {
		defer func() { startedGoRoutines--; c.Signal() }()
		for start <= end {
			time.Sleep(1 * time.Second)
			c.L.Lock()
			lines = append(lines, start)
			start++
			c.L.Unlock()
			c.Signal()
		}
	}

	startedGoRoutines = 3
	go addLine(1, 7)
	go addLine(14, 19)
	go addLine(27, 39)

	for startedGoRoutines > 0 {
		c.L.Lock()
		for len(lines) < 5 && startedGoRoutines > 0 {
			fmt.Println("Wait", startedGoRoutines, len(lines))
			c.Wait()
		}

		fmt.Println(lines)
		lines = lines[len(lines):]
		c.L.Unlock()
	}
}
