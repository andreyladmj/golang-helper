package main

import "fmt"

type Pipeline struct {
	fns  []func(interface{}) interface{}
	done chan interface{}
}

func (p *Pipeline) Add(f func(interface{}) interface{}) {
	p.fns = append(p.fns, f)
}
func (p *Pipeline) Close() {
	close(p.done)
}

func (p *Pipeline) Run(slice []interface{}) <-chan interface{} {
	p.done = make(chan interface{})
	//defer close(p.done)
	stream := p.generator(slice...)
	for _, fn := range p.fns {
		stream = p.apply(stream, fn)
	}
	return stream
}

func (p *Pipeline) apply(stream <-chan interface{}, fn func(interface{}) interface{}) <-chan interface{} {
	currStream := make(chan interface{})
	go func() {
		defer close(currStream)
		for i := range stream {
			select {
			case <-p.done:
				return
			case currStream <- fn(i):
			}
		}
	}()
	return currStream
}

func (p *Pipeline) generator(data ...interface{}) <-chan interface{} {
	stream := make(chan interface{})
	go func() {
		defer close(stream)
		for _, i := range data {
			select {
			case <-p.done:
				return
			case stream <- i:
			}
		}
	}()
	return stream
}

func main() {
	pipeline := Pipeline{}

	pipeline.Add(func(i interface{}) interface{} {
		v := i.(int)
		return v * 2
	})

	pipeline.Add(func(i interface{}) interface{} {
		v := i.(int)
		return v + 1
	})

	pipeline.Add(func(i interface{}) interface{} {
		v := i.(int)
		return v * 2
	})

	pipelineStream := pipeline.Run([]interface{}{1, 2, 3, 4})

	for v := range pipelineStream {
		fmt.Println(v)
		//if v == 10 {
		//	pipeline.Close()
		//}

	}

	defer pipeline.Close()
	//6
	//10
	//14
	//18
}
