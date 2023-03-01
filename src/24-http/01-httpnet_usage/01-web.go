package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	//httpDemoOne()
	httpDemoTwo()
}

func httpDemoOne() {
	var (
		err error
	)

	http.HandleFunc("/", index) // index 为向 url发送请求时，调用的函数

	/**
	 * 1、ListenAndServe 方法有两个参数，其中第一个参数 addr 即监听地址，第二个参数表示服务端处理程序，通常为空。
	 * 2、第二个参数为空时，意味着服务端调用 http.DefaultServeMux 进行处理
	 * 3、服务端编写的业务逻辑处理程序 http.Handle() 或 http.HandleFunc() 默认注入 http.DefaultServeMux 中
	 */
	err = http.ListenAndServe("localhost:8008", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func httpDemoTwo() {
	var (
		err        error
		mux        *http.ServeMux
		httpServer *http.Server
	)

	// 配置路由
	mux = http.NewServeMux()
	mux.HandleFunc("/", index2)

	/**
	 * 自定义 http.Server
	 * 1、Addr表示监听地址和端口
	 * 2、Handler表示配置HandlerFunc
	 * 3、ReadTimeout表示读取超时时间
	 * 4、WriteTimeout表示写入超时时间
	 */
	httpServer = &http.Server{
		Addr:           "localhost:8008",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err = httpServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "黄斌黄")
	b := "黄斌黄"
	w.Write([]byte(b)) //该方法只接受byte类型
}

func index2(w http.ResponseWriter, r *http.Request) {
	content, _ := ioutil.ReadFile("./index.html")
	w.Write(content)
}
