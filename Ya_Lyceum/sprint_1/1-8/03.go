package main

import (
	"fmt"
	"net/http"
)

func answerHandler(w http.ResponseWriter, r *http.Request) {

}

func Authorization(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ok := r.Header.Get("Authorization")
		if ok == "" {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return

		}
		fmt.Fprint(w, "The answer is 42")
	})
}

func main() {
	http.HandleFunc("/answer/", Authorization(answerHandler))
	http.ListenAndServe(":8080", nil)
}
