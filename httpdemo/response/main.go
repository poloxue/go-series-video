package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func responseBody(r *http.Response) {
	content, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("%s", content)
}

func status(r *http.Response) {
	fmt.Println(r.StatusCode) // 状态吗
	fmt.Println(r.Status) // 状态描述
}

func header(r *http.Response) {
	// content-type Content-Type 是都可以获取到内容
	fmt.Println(r.Header.Get("content-type"))
}

func encoding(r *http.Response) {
	// content-type 中会提供编码，比如 content-type="text/html;charset=utf-8"
	// html head meta 获取编码，
	// <meta http-equiv=Content-Type content="text/html;charset=utf-8"
	// 可以通过网页的头部猜测网页的编码信息。
}

func main() {
	r, err := http.Get("https://baidu.com")
	if err != nil {
		panic(err)
	}
	defer func() {_ = r.Body.Close()}()

	// responseBody(r)
	// status(r)
	// header(r)
	encoding(r)
}
