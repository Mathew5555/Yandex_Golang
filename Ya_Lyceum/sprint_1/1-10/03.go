package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Person struct {
	Name string `json:"name"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	person := Person{Name: name}
	jsonBytes, _ := json.Marshal(person)
	w.Write(jsonBytes)
	log.Printf("%s %s", time.Now().Format("2006/02/01 03:04:05"), string(jsonBytes))
}
