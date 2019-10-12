// crawler.go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

var wg sync.WaitGroup

const ADMIN_USER_ID = "y8ST4JmXEzqgM0asXGcypA__"

//消费者线程函数
func consume(scids chan string) {
	defer wg.Done()
	for scid := range scids {
		url := "http://api.miaopai.com/m/v2_channel.json?fillType=259&scid=" + scid
		res,_ := getBody(url)
		fmt.Println(res)
	}

}

//生产者线程函数
func produce(uids chan string, scids chan string, finish chan bool) {
	defer wg.Done()
	for uid := range uids {
		url := "http://www.miaopai.com/gu/u?fen_type=channel&suid=" + uid
		res,_ := getBody(url)
		match := regexp.MustCompile(`data-scid=\\"([\S]+)\\"`)
		for _, scid := range match.FindAllString(res, -1) {
			temps := strings.Split(scid, "\"")
			scid = strings.Trim(temps[1], "\\")
			fmt.Println(scid)
			scids <- scid
		}
	}
	finish <- true
}

func getBody(url string) (string,bool) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return "",false
	} else {
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
			return "",false
		}
		res := string(body)
		return res,true
	}
}

func main() {
	scids := make(chan string)
	finish := make(chan bool)

	//消费者数量
	n := 10
	//生产者数量
	m := 3
	wg.Add(n + m)
	//关注的用户通道
	uids := make(chan string)
	for i := 0; i < m; i++ {
		go produce(uids, scids, finish)
	}

	for i := 0; i < n; i++ {
		go consume(scids)
	}

	queue := make([]string, 1)
	queue = append(queue, ADMIN_USER_ID)
	total := 1
	for len(queue) > 0 && total < 5 {
		res,success := getBody("http://www.miaopai.com/gu/follow?suid=" + queue[0])
		queue = queue[1:]
		if success == false {
			fmt.Println("res=", res)
			continue
		}

		total++
		match := regexp.MustCompile(`suid=\\"([\S]+)\\"`)
		for _, uid := range match.FindAllString(res, -1) {
			temps := strings.Split(uid, "\"")
			uid = strings.Trim(temps[1], "\\")
			fmt.Println(uid)
			uids <- uid
			queue = append(queue, uid)
		}
	}
	close(uids)

	//等待生产者线程结束
	produce_num := 0
	for produce_num < m {
		<-finish
		produce_num++
	}
	close(scids)

	wg.Wait()

	fmt.Println("done")
}
