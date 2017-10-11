// channel_sync.go
package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	done <- true
	done <- true
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

}

func main() {
	done := make(chan bool, 1)
	go worker(done)
	<-done
	<-done
	time.Sleep(2 * time.Second)
}
