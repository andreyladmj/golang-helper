package main

import (
	"fmt"
	"unsafe"
)

func getAddr(fn interface{}) uintptr {
	// emptyInterface is the header for an interface{} value.
	type emptyInterface struct {
		typ   uintptr
		value *uintptr
	}
	e := (*emptyInterface)(unsafe.Pointer(&fn))
	return *e.value
}
func main() {
	a := "a1123"
	b := []byte(a)
	size1 := unsafe.Sizeof(a)
	size2 := unsafe.Sizeof(b)
	fmt.Println(size1, len(a), a)
	fmt.Println(size2, len(b), b)
	fmt.Println(string(b))

	fmt.Println("First Element Size", unsafe.Sizeof(a[0]), unsafe.Sizeof(a[0:3]))
	fmt.Println("First Element Size", unsafe.Sizeof(b[0]), unsafe.Sizeof(b[0:3]))

	o := 2
	fmt.Println(unsafe.Sizeof(a[0]), unsafe.Sizeof(o), unsafe.Sizeof(&o), unsafe.Sizeof(2.4))

	memoryAddress := uintptr(unsafe.Pointer(&a))

	pointer := (*[]byte)(unsafe.Pointer(memoryAddress))
	fmt.Print("One more: ", *pointer, " ")

	pointer2 := (*string)(unsafe.Pointer(memoryAddress))
	fmt.Print("One more: ", *pointer2, " ")
	//
	//pointer := a[0]
	//fmt.Println(pointer)
	//firstSliceElementAddr := getAddr(pointer) + 8
	//fmt.Println("firstSliceElementAddr", firstSliceElementAddr)
	//
	//p := (*int)(unsafe.Pointer(firstSliceElementAddr))
	//fmt.Println("p", p)
	//
	//fmt.Print(*p, " ")

	//code := **(**uintptr)(unsafe.Pointer(&pointer))
	//fmt.Println(code)

	//fmt.Println(unsafe.Pointer(&pointer))
	//
	//memoryAddress := uintptr(unsafe.Pointer(a[0])) + unsafe.Sizeof(array[0])

	//a := make([]int, 5)
	//a[0]=1
	//a[2]=1
	//a[3]=1
	//a[4]=1
	//b := a[1:4]
	//b[1] = 1000
	//fmt.Println(&a, a)
	//fmt.Println(&b, b)
	//a = append(a, 5)
	//a = append(a, 5)
	//a = append(a, 5)
	//fmt.Println(&a, a)
	//fmt.Println(&b, b)
}
