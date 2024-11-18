package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

var (
	t = make(map[string]string)
)

type Response struct {
	Status string            `json:"status"`
	Result map[string]string `json:"result"`
}

func HelloHandler2(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	t["greetings"] = "hello"
	t["name"] = name
}

func Sanitize2(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if notOk, _ := regexp.MatchString("[^a-zA-Z]", name); notOk {
			panic("")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func SetDefaultName2(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			t["greetings"] = "hello"
			t["name"] = "stranger"
			return
		}
		next.ServeHTTP(w, r)
	})
}

func RPC(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				response := Response{Status: "error", Result: make(map[string]string)}
				t = make(map[string]string)
				jsonBytes, _ := json.Marshal(response)
				fmt.Fprint(w, string(jsonBytes))
				return
			}
		}()
		next.ServeHTTP(w, r)
		response := Response{Status: "ok", Result: t}
		jsonBytes, _ := json.Marshal(response)
		fmt.Fprintf(w, string(jsonBytes))
	})
}

func main() {
	http.HandleFunc("/", RPC(SetDefaultName2(Sanitize2(HelloHandler))))
	http.ListenAndServe(":8080", nil)
}
