package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func downloadFile(url, filename string) {
	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer func() {_ = r.Body.Close()}()

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer func() {_ = f.Close()}()

	n, err := io.Copy(f, r.Body)
	fmt.Println(n, err)
}

type Reader struct {
	io.Reader
	Total int64
	Current int64
}

func (r *Reader) Read(p []byte) (n int, err error){
	n, err = r.Reader.Read(p)

	r.Current += int64(n)
	fmt.Printf("\r进度 %.2f%%", float64(r.Current * 10000/ r.Total)/100)

	return
}

func DownloadFileProgress(url, filename string) {
	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer func() {_ = r.Body.Close()}()

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer func() {_ = f.Close()}()

	reader := &Reader{
		Reader: r.Body,
		Total: r.ContentLength,
	}

	_, _ = io.Copy(f, reader)

}

func main() {
	// 自动文件下载，比如自动下载图片、压缩包
	url := "https://user-gold-cdn.xitu.io/2019/6/30/16ba8cb6465a6418?w=826&h=782&f=png&s=279620"
	filename := "poloxue.png"
	DownloadFileProgress(url, filename)
}
