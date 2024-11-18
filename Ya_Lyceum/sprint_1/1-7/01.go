package main

import (
	"fmt"
	"net/http"
	"regexp"
)

func StrangerHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "stranger"
	}
	if ok, _ := regexp.MatchString("[a-zA-Z]", name); !ok {
		name = "dirty hacker"
	}
	fmt.Fprintf(w, "hello %s", name)
}

func main() {
	http.HandleFunc("/", StrangerHandler)
	http.ListenAndServe(":8080", nil)
}
