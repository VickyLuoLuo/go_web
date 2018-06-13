package main

import (
	"net/http"
	"fmt"
	"log"
	"strings"
)

func sayhello(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	fmt.Println("param_form: ",r.Form)
	fmt.Println("path: ", r.URL.Path)
	fmt.Println("scheme: ", r.URL.Scheme)
	fmt.Println("one_param_value: ", r.Form["hello"])
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("value: ",strings.Join(v," "))
	}
	fmt.Fprintf(w, "hello go")
}

func main()  {
	http.HandleFunc("/", sayhello)
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}