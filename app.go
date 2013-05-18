package main

import (
	"fmt"
	"github.com/nrml/membership-go/endpoints"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/favicon.ico", endpoints.StaticHandler)
	http.HandleFunc("/login", endpoints.LoginHandler)
	http.HandleFunc("/", endpoints.Handler)

	port := os.Args[1]

	fmt.Println("listening on port:", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("{error: \"%s\"}", err.Error())
	}
}
