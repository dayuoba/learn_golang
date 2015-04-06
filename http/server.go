package main

import (
	"fmt"
	"log"
	"net/http"
	"io"
)

type Hello struct{}

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, "Hello!")
	go fmt.Println("request");
}

func HelloServe(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello,world!\n")
}

func main() {
	var h Hello
	http.HandleFunc("/hello", HelloServe)
	err := http.ListenAndServe("localhost:4000", h)
	
	if err != nil {
		log.Fatal(err)
	}
}
