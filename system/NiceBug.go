package main

import (
	"fmt"
	"log"
)

type customError struct{}

func (c *customError) Error() string {
	return "Find the bug"
}

func fail() ([]byte, *customError) {
	return nil, nil
}

func main() {
	var err error
	_, err = fail()

	print(err)
	fmt.Println(err)

	if err != nil {
		log.Fatal("why did this fail?")
	}

	log.Println("No Error")
}
