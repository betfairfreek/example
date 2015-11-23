package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	."github.com/example/cloudplatform/handler"
)

const PORT string = ":8080"

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/getMe", GetMe)
	r.HandleFunc("/loggin",Loggin)

	fmt.Println("Http-Proxy Listening on port", PORT)
	http.ListenAndServe(PORT, r)
}
