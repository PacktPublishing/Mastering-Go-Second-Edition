package main

import (
	"fmt"
	"net/http"
	"runtime"
	"sync/atomic"
)

var count int32

func handleAll(w http.ResponseWriter, r *http.Request) {
	temp := atomic.LoadInt32(&count)
	atomic.AddInt32(&count, 1)
	fmt.Println("count:", temp)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	http.HandleFunc("/", handleAll)
	http.ListenAndServe(":8080", nil)
}
