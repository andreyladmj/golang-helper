package context

import (
	"fmt"
	"io"
	"os"
	"os/signal"
	"sync"
	"time"
)

type device struct {
	problem bool
}

func (d *device) Write(p []byte) (n int, err error) {
	for d.problem {
		time.Sleep(time.Second)
	}

	fmt.Println(string(p))
	return len(p), nil
}

func main() {
	const grs = 10

	var d device
	l := New(&d, grs)

	for i := 0; i < grs; i++ {
		go func(id int) {
			for {
				l.Println(fmt.Sprintf("%d: log data", id))
				time.Sleep(10 * time.Millisecond)
			}
		}(i)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	for {
		<-sigChan

		d.problem = !d.problem
	}
}

type Logger struct {
	ch chan string
	wg sync.WaitGroup
}

func New(w io.Writer, cap int) *Logger {
	l := Logger{
		ch: make(chan string, cap),
	}

	l.wg.Add(1)
	go func() {
		defer l.wg.Done()
		for v := range l.ch {
			fmt.Fprintln(w, v)
		}
	}()

	return &l
}

func (l *Logger) Close() {
	close(l.ch)
	l.wg.Wait()
}

func (l *Logger) Println(v string) {
	select {
	case l.ch <- v:
	default:
		fmt.Println("DROP")
	}
}
