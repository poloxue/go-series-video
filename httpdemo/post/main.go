package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func postForm() {
	// form data 形式 query string，类似于 name=poloxue&age=18
	data := make(url.Values)
	data.Add("name", "poloxue")
	data.Add("age", "18")
	payload := data.Encode()

	r, _ := http.Post(
		"http://httpbin.org/post",
		"application/x-www-form-urlencoded",
		strings.NewReader(payload),
	)
	defer func() {_ = r.Body.Close()}()

	content, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("%s", content)
}

func postJson() {
	u := struct {
		Name string `json:"name"`
		Age int `json:"age"`
	}{
		Name: "poloxue",
		Age: 18,
	}
	payload, _ := json.Marshal(u)
	r, _ := http.Post(
		"http://httpbin.org/post",
		"application/json",
		bytes.NewReader(payload),
	)
	defer func() {_ = r.Body.Close()}()

	content, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("%s", content)
}

func main() {
	// post 请求的本质，它是 request body 提交，相对于 get 请求（urlencoded 提交查询参数, 提交内容有大小限制，好像 2kb）
	// post 不同的形式也就是 body 的格式不同
	// post form 表单，body 就是 urlencoded 的形式，比如 name=poloxue&age=18
	// post json，提交的 json 格式
	// post 文件，其实也是组织 body 数据
	postJson()
}
