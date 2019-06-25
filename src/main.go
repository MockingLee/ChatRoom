package main

import (
	"fmt"
	"net/http"
)

func sayHello(writer http.ResponseWriter , request *http.Request){
	request.ParseForm()
	fmt.Println("receive request")
	fmt.Println("method : " , request.Method)
	fmt.Println(request.Form)
	fmt.Fprintln(writer , "hello")
}

func main() {
	fmt.Println("Starting web server...")
	http.HandleFunc("/",sayHello)
	err := http.ListenAndServe("0.0.0.0:9999" , nil)
	if err == nil{
		fmt.Println(err)
	}
}
