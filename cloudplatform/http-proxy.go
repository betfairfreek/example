package main

import (
	"fmt"
	"net/http"
)

const PORT string = ":8080"

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,"<h1>HomePage</h1>")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/login", loginHandler)

	fmt.Println("Http-Server Started on port", PORT)
	http.ListenAndServe(PORT, nil)
}
