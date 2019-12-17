package main

import (
	"context"
	"net"
	"net/http"
	"time"
)

func main() {
	// https://colobu.com/2016/07/01/the-complete-guide-to-golang-net-http-timeouts/
	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (conn net.Conn, e error) {
				return net.DialTimeout(network, addr, 2*time.Second)
			},
			ResponseHeaderTimeout: 5 * time.Second,
			TLSHandshakeTimeout:   2 * time.Second,
			IdleConnTimeout:       60 * time.Second,
		},
	}
	r, _ := client.Get("http://httpbin.org/delay/10")
}
