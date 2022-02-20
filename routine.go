package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func s(from string) {
	for i := 0; i < 10; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("direct")
	go s("Sarah")
	fmt.Println("done")

	f("Marky")
	go func(msg string) {
		fmt.Println(msg)
	}("going")
	time.Sleep(time.Hour)
	fmt.Println("done")
}
