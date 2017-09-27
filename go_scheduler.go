// go_scheduler.go
package main

/*
 * http://www.sarathlakshman.com/2016/06/15/pitfall-of-golang-scheduler
 */
import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var x int
	threads := runtime.GOMAXPROCS(0)
	fmt.Println(threads)

	for i := 0; i < threads; i++ {
		go func() {
			for {
				x++
				runtime.Gosched()
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("x=", x)
}
