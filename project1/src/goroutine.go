package main

/* go线程和协程的区别
https://juejin.im/post/5d9a9c12e51d45781420fb7e
*/
import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func TestGorutine() {
	//指定最大 P 为 1，从而管理协程最多的线程为 1 个
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		fmt.Println(1)
		fmt.Println(2)
		fmt.Println(3)
		wg.Done()
	}()

	go func() {
		fmt.Println(65)
		fmt.Println(66)
		time.Sleep(time.Second)
		fmt.Println(67)
		wg.Done()
	}()

	wg.Wait()
}

func main() {
	TestGorutine()
}
