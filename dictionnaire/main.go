package main

import (
	"dictionnaire/dict"
	"net/http"
	"fmt"
)


func main() {
    s := dict.New()
	fmt.Println("Server is running...")
	http.HandleFunc("/get",  s.GetHandler)
	go http.HandleFunc("/post", s.PostHandler)
	http.HandleFunc("/update",  s.UpdateHandler)
	http.HandleFunc("/delete",  s.DeleteHandler)
	http.ListenAndServe(":8080", nil)
}