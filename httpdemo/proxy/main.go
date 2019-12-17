package main

import (
	"io"
	"net/http"
	"net/url"
	"os"
)

func main() {
	// proxyUrl, _ := url.Parse("socks5://127.0.0.1:1080")
	proxyUrl, _ := url.Parse("http://127.0.0.1:8087")
	t := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}
	// 一般主要两种，http 代理 和 shadowsocks 的代码, socks5
	client := http.Client{Transport: t}
	r, _ := client.Get("https://google.com")
	defer func() { _ = r.Body.Close() }()

	_, _ = io.Copy(os.Stdout, r.Body)
	// session
	// session.Get
	// session.Post
	// session.Put
	// session.Delete
}
