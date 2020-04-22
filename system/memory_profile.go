package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

type Simple struct {
	flag  bool
	i1    int64
	flag2 bool
	i2    int64
	flag3 bool
	i3    int64
}

var cpuprofile = "memory_profile_cpu.txt"
var memprofile = "memory_profile_mem.txt"

func main() {
	f, err := os.Create(cpuprofile)
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer f.Close()
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	c := Simple{}
	c2 := Simple{}
	c3 := Simple{}
	c4 := Simple{}
	c5 := Simple{}

	f, err = os.Create(memprofile)
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	defer f.Close()
	runtime.GC() // get up-to-date statistics
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}

	fmt.Println(c, c2, c3, c4, c5)
}
