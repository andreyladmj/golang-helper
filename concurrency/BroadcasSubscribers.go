package main

import (
	"fmt"
	"sync"
	"time"
)

type Button struct {
	Clicked *sync.Cond
}

func main() {
	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	subscribe := func(c *sync.Cond, wg *sync.WaitGroup, fn func()) {
		wg.Add(1)
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)
		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			defer wg.Done()
			c.Wait()
			fn()
		}()
		goroutineRunning.Wait()
	}

	var clickRegistered sync.WaitGroup
	//clickRegistered.Add(3)
	subscribe(button.Clicked, &clickRegistered, func() {
		fmt.Println("Maximizing window.")
		//clickRegistered.Done()
	})
	subscribe(button.Clicked, &clickRegistered, func() {
		fmt.Println("Displaying annoying dialog box!")
		//clickRegistered.Done()
	})
	subscribe(button.Clicked, &clickRegistered, func() {
		fmt.Println("Mouse clicked.")
		//clickRegistered.Done()
	})

	//time.Sleep(1*time.Second) // Or we should use WaitGroup to confirm goroutines was running
	button.Clicked.Broadcast()

	clickRegistered.Wait()
}
