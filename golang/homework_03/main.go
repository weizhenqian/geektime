package main

/*
1.接收客户端 request，并将 request 中带的 header 写入 response header
2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3.Server端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server端的标准输出
4.当访问 localhost/healthz 时，应返回 200
*/

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/thinkeridea/go-extend/exnet"
)

//1.接收request的header
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	header := r.Header
	//调用curl命令，curl -H 'Host:157.166.226.25' -H 'Accept-Language:es' -H 'Cookie:ID=1234' 127.0.0.1:8888
	//返回值：map[Accept:[*/*] Accept-Language:[es] Cookie:[ID=1234] User-Agent:[curl/7.77.0]]
	//由于返回值是map，所以对map进行遍历，将kv存储至response header.
	for k, v := range header {
		//v的值是header.Header，无法之间转成string的形式，故通过fmt，转化成其本身的值，再转成string.
		//此次应该有更好的办法，但是对golang的类型转换不太理解。
		v, _ := fmt.Printf("%v", v)
		//个人感觉此方法比较low
		w.Header().Add(k, strconv.Itoa(v))
	}
	fmt.Fprintln(w, "Header的内容为:", header)
}

//读取当前系统的环境变量中的 VERSION 配置，并写入 response header
//启动方式：export VERSION=xxxx && go run main.go
func GetVersion(w http.ResponseWriter, r *http.Request) {
	//os模块用于系统相关操作
	version := os.Getenv("VERSION")
	w.Header().Add("version", version)
	fmt.Fprintln(w, "Version的内容为:", version)
}

//3.Server端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server端的标准输出
//IP获取的地址为127.0.0.1，code尝试了httpsnoop，但是调用失败了。
func SaveClientInfo(w http.ResponseWriter, r *http.Request) {
	info := exnet.ClientIP(r)
	fmt.Fprintln(w, "Client的IP信息为:%V", info)
}

//不太理解StatusText的用法，修改为404并无反应
func HealthCode(w http.ResponseWriter, r *http.Request) {
	//StatusText(code int)
	code := 200
	w.WriteHeader(code)
	http.StatusText(code)
	fmt.Fprintln(w, code)
}

func main() {
	/*
		// myH is your app's http handler, perhaps a http.ServeMux or similar.
		var myH http.Handler
		// wrappedH wraps myH in order to log every request.
		wrappedH := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			m := httpsnoop.CaptureMetrics(myH, w, r)
			log.Printf(
				"%s %s (code=%d dt=%s written=%d)",
				r.Method,
				r.URL,
				m.Code,
				m.Duration,
				m.Written,
			)
		})
	*/
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/version", GetVersion)
	http.HandleFunc("/healthz", HealthCode)
	http.ListenAndServe(":8888", nil)
}
