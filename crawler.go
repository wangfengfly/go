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

const ADMIN_USER_ID = "QGVbe7gWfyLfb046Xntz0Q__"

//消费者线程函数
func consume(scids chan string) {
	defer wg.Done()
	for scid := range scids {
		url := "http://api.miaopai.com/m/v2_channel.json?fillType=259&scid=" + scid
		res := getBody(url)
		fmt.Println(res)
	}

}

//生产者线程函数
func produce(uids chan string, scids chan string) {
	defer wg.Done()
	for uid := range uids {
		url := "http://www.miaopai.com/gu/u?fen_type=channel&suid=" + uid
		res := getBody(url)
		match := regexp.MustCompile(`data-scid=\\"([\S]+)\\"`)
		for _, scid := range match.FindAllString(res, -1) {
			temps := strings.Split(scid, "\"")
			scid = strings.Trim(temps[1], "\\")
			fmt.Println(scid)
			scids <- scid
		}
	}
	close(scids)
}

func getBody(url string) string {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return ""
	} else {
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		res := string(body)
		return res
	}
}

func main() {
	scids := make(chan string)
	//消费者数量
	n := 10
	//生产者数量
	m := 1
	wg.Add(n + m)

	uids := make(chan string)
	for i := 0; i < m; i++ {
		go produce(uids, scids)
	}

	for i := 0; i < n; i++ {
		go consume(scids)
	}

	res := getBody("http://www.miaopai.com/gu/follow?suid=" + ADMIN_USER_ID)
	if res == "" {
		return
	}
	match := regexp.MustCompile(`suid=\\"([\S]+)\\"`)
	for _, uid := range match.FindAllString(res, -1) {
		temps := strings.Split(uid, "\"")
		uid = strings.Trim(temps[1], "\\")
		fmt.Println(uid)
		uids <- uid
	}

	close(uids)

	wg.Wait()

	fmt.Println("done")
}
