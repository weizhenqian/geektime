package main

/*
接收客户端 request，并将 request 中带的 header 写入 response header
读取当前系统的环境变量中的 VERSION 配置，并写入 response header
Server端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server端的标准输出
当访问 localhost/healthz 时，应返回 200
*/

import (
	"fmt"
	"net/http"
	"os"

	"github.com/thinkeridea/go-extend/exnet"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	header := r.Header
	fmt.Fprintln(w, "Header的内容为:", header)
}

func GetVersion(w http.ResponseWriter, r *http.Request) {
	version := os.Getenv("VERSION")
	fmt.Fprintln(w, "Version的内容为:", version)
}

func SaveClientInfo(w http.ResponseWriter, r *http.Request) {
	info := exnet.ClientIP(r)
	fmt.Fprintln(w, "Client的IP信息为:%V", info)
}

func HealthCode(w http.ResponseWriter, r *http.Request) {
	code := 200
	fmt.Fprintln(w, code)
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/version", GetVersion)
	http.HandleFunc("/info", SaveClientInfo)
	http.HandleFunc("/healthz", HealthCode)
	http.ListenAndServe(":8888", nil)

}
