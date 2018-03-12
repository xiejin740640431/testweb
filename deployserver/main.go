package main

import (
	"os/exec"
	"log"
	"net/http"
	"io"
)

func reLauncher() {
	cmd := exec.Command("sh", "./deploy.sh")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
}

//git注释
func test(w http.ResponseWriter,r *http.Request)  {
	io.WriteString(w,"<h1>Hello ,this is deploy server!</h1>")
	reLauncher()
}

func main() {
	http.HandleFunc("/",test)
	http.ListenAndServe(":5000",nil)
}
