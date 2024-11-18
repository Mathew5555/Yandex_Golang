package main

import (
	"fmt"
	"net/http"
	"time"
)

var (
	prev = 0
	cur  = 1
	cnt  = 0
)

func FibonacciHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, prev)
	prev, cur = cur, prev+cur
	cnt++
}

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "rpc_duration_milliseconds_count %d", cnt)

}

func StartServer(t time.Duration) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", FibonacciHandler)
	mux.HandleFunc("/metrics", MetricsHandler)
	http.ListenAndServe(":8080", mux)
}

func main() {
	StartServer(time.Second)
}
