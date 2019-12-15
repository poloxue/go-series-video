package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"

	cookiejar2 "github.com/juju/persistent-cookiejar"
)

func rrCookie() {
	// 模拟完成一个登录
	// 请求一个页面，传递基本的登录信息，将响应的 cookie 设置到下一次之上重新请求
	// 请求 http://httpbin.org/cookies/set?name=poloxue&password=123456
	// 返回 set-cookie:
	// 再一次请求呢携带上 cookie，
	// 首页 http://httpbin.org/cookies 就会通过 body 打印出已经设置 cookie
	// http://httpbin.org/cookies/set? => response
	// request => http://httpbin.org/cookies
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	firstRequest, _ := http.NewRequest(
		http.MethodGet,
		"http://httpbin.org/cookies/set?name=poloxue&password=123456",
		nil,
	)
	firstResponse, _ := client.Do(firstRequest)
	defer func() { _ = firstResponse.Body.Close() }()

	secondRequest, _ := http.NewRequest(
		http.MethodGet,
		"http://httpbin.org/cookies",
		nil,
	)

	for _, cookie := range firstResponse.Cookies() {
		secondRequest.AddCookie(cookie)
	}

	secondResponse, _ := client.Do(secondRequest)
	defer func() { _ = secondResponse.Body.Close() }()

	content, _ := ioutil.ReadAll(secondResponse.Body)
	fmt.Printf("%s\n", content)
}

func jarCookie() {
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
	}
	r, _ := client.Get("http://httpbin.org/cookies/set?username=poloxue&password=123456")
	defer func() { _ = r.Body.Close() }()

	_, _ = io.Copy(os.Stdout, r.Body)
}

func login(jar http.CookieJar) {

	client := &http.Client{
		Jar: jar,
	}
	r, _ := client.PostForm(
		"http://localhost:8080/login",
		url.Values{"username": {"poloxue"}, "password": {"poloxue123"}},
	)
	defer func() { _ = r.Body.Close() }()
	fmt.Println(r.Cookies())

	_, _ = io.Copy(os.Stdout, r.Body)
}

func center(jar http.CookieJar) {
	client := &http.Client{
		Jar: jar,
	}
	r, _ := client.Get("http://localhost:8080/center")
	defer func() { _ = r.Body.Close() }()

	_, _ = io.Copy(os.Stdout, r.Body)
}

func main() {
	// rrCookie()
	// jarCookie()
	// cookie 的分类有两种 一种是会话期 cookie 一种是持久性 cookie
	// jar, _ := cookiejar.New(nil)
	jar, _ := cookiejar2.New(nil)
	// login(jar)
	center(jar)
	_ = jar.Save()
}
