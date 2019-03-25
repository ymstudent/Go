//一个迷你的回声服务器和计数服务器
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

//使用fetchall并发请求时，加锁和不加锁，count的值是不一样
//fetchall并发请求： http://localhost:8000 http://localhost:8000 http://localhost:8000/count http://localhost:8000/count
//不加锁count为4，加锁为2
//原因：竞态BUG（参考9.1节）
//TODO 待学习
var mu sync.Mutex

var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
