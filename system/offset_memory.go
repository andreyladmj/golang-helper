package main

import (
	"fmt"
	"unsafe"
)

type SimpleStruct struct {
	flag  bool
	i1    int64
	flag2 bool
	i2    int64
	flag3 bool
	i3    int64
}

type SimpleStructOptimized struct {
	i1    int64
	i2    int64
	i3    int64
	flag  bool
	flag2 bool
	flag3 bool
}

func main() {
	notOptimizedStruct := SimpleStruct{}
	fmt.Println("Struct Size of notOptimizedStruct", unsafe.Sizeof(notOptimizedStruct))

	fmt.Println("offset flag", unsafe.Offsetof(notOptimizedStruct.flag))
	fmt.Println("offset i1", unsafe.Offsetof(notOptimizedStruct.i1))
	fmt.Println("offset flag2", unsafe.Offsetof(notOptimizedStruct.flag2))
	fmt.Println("offset i2", unsafe.Offsetof(notOptimizedStruct.i2))
	fmt.Println("offset flag3", unsafe.Offsetof(notOptimizedStruct.flag3))
	fmt.Println("offset i3", unsafe.Offsetof(notOptimizedStruct.i3))

	fmt.Println("\n\n")
	optimizedStruct := SimpleStructOptimized{}
	fmt.Println("Struct Size of optimizedStruct", unsafe.Sizeof(optimizedStruct))

	fmt.Println("offset i1", unsafe.Offsetof(optimizedStruct.i1))
	fmt.Println("offset i2", unsafe.Offsetof(optimizedStruct.i2))
	fmt.Println("offset i3", unsafe.Offsetof(optimizedStruct.i3))
	fmt.Println("offset flag", unsafe.Offsetof(optimizedStruct.flag))
	fmt.Println("offset flag2", unsafe.Offsetof(optimizedStruct.flag2))
	fmt.Println("offset flag3", unsafe.Offsetof(optimizedStruct.flag3))
}
