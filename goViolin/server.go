//referred from
//https://blog.scottlogic.com/2017/02/28/building-a-web-app-with-go.html

package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! \n This is Bhoomika.")
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8080", nil)
}
