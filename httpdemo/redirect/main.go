package main

import (
	"errors"
	"fmt"
	"net/http"
)

func redirectLimitTimes() {
	// 限制重定向的次数
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) > 10 {
				return errors.New("redirect too times")
			}
			return nil
		},
	}

	request, _ := http.NewRequest(
		http.MethodGet,
		"http://httpbin.org/redirect/20",
		nil,
	)
	_, err := client.Do(request)
	if err != nil {
		panic(err)
	}
}

func redirectForbidden() {
	// 禁止重定向
	// 登录请求，防止重定向到首页
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	request, _ := http.NewRequest(
		http.MethodGet,
		"http://httpbin.org/cookies/set?name=poloxue",
		nil,
	)
	r, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer func() {_ = r.Body.Close()}()
	fmt.Println(r.Request.URL)
}

func main() {
	// 重定向
	// 返回一个状态码，3xx 301 302 303 307 308
	redirectForbidden()
}
