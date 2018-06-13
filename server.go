package main

import (
	"net/http"
	"fmt"
	"log"
)

func sayhello(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	fmt.Println("path:", r.URL.Path)
	fmt.Fprintf(w, "hello go")
}

func main()  {
	http.HandleFunc("/hello", sayhello)
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}