package main

import (
	"fmt"
	"net/http"
	"runtime"
	"sync/atomic"
)

var count int32

func handleAll(w http.ResponseWriter, r *http.Request) {
	atomic.AddInt32(&count, 1)
}

func getCounter(w http.ResponseWriter, r *http.Request) {
	temp := atomic.LoadInt32(&count)
	fmt.Println("Count:", temp)
	fmt.Fprintf(w, "<h1 align=\"center\">%d</h1>", count)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	http.HandleFunc("/getCounter", getCounter)
	http.HandleFunc("/", handleAll)
	http.ListenAndServe(":8080", nil)
}
