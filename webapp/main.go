package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "HelloWorld!", r.URL.Path)

}
func main() {
	//映射
	http.HandleFunc("/", handler)

	//路由
	http.ListenAndServe(":8080", nil)

}
