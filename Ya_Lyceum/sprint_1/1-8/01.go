package main

import (
	"fmt"
	"net/http"
	"regexp"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, "Hello, %s!", name)
}

func Sanitize(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if notOk, _ := regexp.MatchString("[^a-zA-Z]", name); notOk {
			fmt.Fprint(w, "Hello, dirty hacker!")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func SetDefaultName(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			fmt.Fprint(w, "Hello, stranger!")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	http.HandleFunc("/", SetDefaultName(Sanitize(HelloHandler)))
	http.ListenAndServe(":8080", nil)
}
