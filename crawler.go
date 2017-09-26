// crawler.go
package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

//消费者线程函数
func consume(urls chan string) {
	defer wg.Done()
	f, err := os.OpenFile("./test", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for url := range urls {
		if _, err = f.WriteString(url + "\n"); err != nil {
			panic(err)
		}

	}

}

//生产者线程函数
func produce(c chan string) {
	defer wg.Done()
	for i := 0; i < 20; i++ {
		url := strconv.Itoa(i)
		c <- url
	}
	close(c)
}

func main() {
	c := make(chan string)
	//消费者数量
	n := 5
	//生产者数量
	m := 1
	wg.Add(n + m)

	for i := 0; i < n; i++ {
		go consume(c)
	}

	for i := 0; i < m; i++ {
		go produce(c)
	}

	//这里可以执行reduce操作，比如合并多个线程的结果

	wg.Wait()
}
