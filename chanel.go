package main

import "fmt"

func main() {

	messages := make(chan string)
	marky := make(chan string)

	go func() { messages <- "hellllo you can receive this outsite" }()
	go func() { marky <- "marky" }()

	msg := <-messages
	msgs := <-marky

	fmt.Println(msg)
	fmt.Println(msgs)

}
