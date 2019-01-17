package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {

	// 调用对应路由处理函数处理路由
	http.HandleFunc("/", handle)

	// 输出访问数量接口
	http.HandleFunc("/count", counter)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

//  普通路由处理函数
func handle(w http.ResponseWriter, r *http.Request) {

	mu.Lock()

	count++

	mu.Unlock()

	fmt.Fprintf(w, "the path is %s", r.URL.Path)

}

// 访问统计路由处理
func counter(w http.ResponseWriter, r *http.Request) {

	mu.Lock()

	fmt.Fprintf(w, "flow of the server: %d", count)

	mu.Unlock()

}
