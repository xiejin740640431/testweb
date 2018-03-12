package main

import (
	"net/http"
	"io"
)

func test(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"<h1>this is 2.0 version!</h1>")
}

func main(){
	http.HandleFunc("/",test)
	http.ListenAndServe(":8000",nil)
}
