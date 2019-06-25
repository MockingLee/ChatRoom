package main

import (
	"fmt"
	"net/http"
)

func sayHello(writer http.ResponseWriter , request *http.Request){
	request.ParseForm()
	fmt.Println("receive request")
	fmt.Println("method : " , request.Method)
	fmt.Println()
	fmt.Println(request.Form)
}

func main() {
	fmt.Println("Starting web server...")
	http.HandleFunc()
}
