package main

import "fmt"

func main() {
	//a := 1
	//a++
	//fmt.Println(a)
	a := [...]int{1, 2}
	a2 := []int{1, 2}
	q := &a2

	fmt.Printf("%v, %v, %p\n", a, len(a), a)
	fmt.Printf("%v, %v, %p = (%p)\n\n", a2, len(a2), a2, q)

	a2 = append(a2, a2...)
	a2 = append(a2, a2...)

	fmt.Printf("%v, %v, %p\n", a, len(a), a)
	fmt.Printf("%v, %v, %p = (%p)\n\n\n", a2, len(a2), a2, q)

	a6 := []int{-10, 1, 2, 3, 4, 5}
	a4 := []int{-1, -2, -3, -4}
	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)

	copy(a6, a4)
	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)
	fmt.Println()

	s1 := "THE"
	fmt.Printf("%p\n", &s1)
	modifyString(s1)
	fmt.Println(s1)
	s1 += " HELLO"
	fmt.Printf("%p\n", &s1)
}

func modifyString(s1 string) {
	fmt.Printf("%p\n", &s1)
	s1 = s1 + " HELLO"
}
