// buffered_channel_block.go
package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3

	fmt.Println("done")
}
