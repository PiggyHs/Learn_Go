package main

import (
	"fmt"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter,  r *http.Request)  {
	r.PostForm()

	fmt.Println(r.Form)
	fmt.Println("path:", r.URL.Path)
	fmt.Println("scheme:", r.URL.Scheme)
	fmt.Println(r.Form["url_log"])

	for k,v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("value:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello, HeShuai")
}
