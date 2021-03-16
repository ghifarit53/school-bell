package main

import (
	"fmt"
	"net/http"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.URL.Path)
    p := "." + r.URL.Path
    if p == "./" {
        p = "./static/index.html"
    }
    http.ServeFile(w, r, p)
}

func main() {
	http.HandleFunc("/", handlerIndex)

	var address = "localhost:1337"
	fmt.Printf("server started at %s\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
