// channel.go
package main

import (
	"fmt"
)

func main() {
	messages1 := make(chan string)
	messages2 := make(chan string)
	messages3 := make(chan string, 1)

	go func() {
		messages2 <- "messages2"
		//messages <- "hello"
	}()
	go func() {
		msg := <-messages1
		fmt.Println(msg)

	}()

	messages3 <- "world"
	<-messages3

	fmt.Println("done")
}
