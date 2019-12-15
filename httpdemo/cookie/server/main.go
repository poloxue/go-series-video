package main

import (
	"net/http"
	"time"
)

func login(w http.ResponseWriter, req *http.Request) {
	username := req.PostFormValue("username")
	password := req.PostFormValue("password")

	if username == "poloxue" && password == "poloxue123" {
		http.SetCookie(w, &http.Cookie{
			Name:    "isLogin",
			Value:   "1",
			Expires: time.Now().Add(3 * time.Hour),
		})
		_, _ = w.Write([]byte("登录成功\n"))
	} else {
		_, _ = w.Write([]byte("登录失败"))
	}

	return
}

func center(w http.ResponseWriter, r *http.Request) {
	isLogin, err := r.Cookie("isLogin")
	if err == http.ErrNoCookie {
		_, _ = w.Write([]byte("无法访问"))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if isLogin.Value != "1" {
		_, _ = w.Write([]byte("无法访问"))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, _ = w.Write([]byte("个人主页\n"))
}

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/center", center)
	_ = http.ListenAndServe(":8080", nil)
}
