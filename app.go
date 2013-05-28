package main

import (
	//"fmt"
	//"github.com/nrml/membership-go/endpoints"
	"github.com/nrml/membership-go/service"
	"github.com/nrml/rpc-go"
	"log"
	//"net/http"
	"os"
	"strconv"
)

func main() {
	//http.HandleFunc("/favicon.ico", endpoints.StaticHandler)
	//http.HandleFunc("/login", endpoints.LoginHandler)
	//http.HandleFunc("/", endpoints.Handler)

	port := os.Args[1]

	//fmt.Println("listening on port:", port)

	// err := http.ListenAndServe(":"+port, nil)
	// if err != nil {
	// 	fmt.Printf("{error: \"%s\"}", err.Error())
	// }

	//set up to listen for rpc

	iport, err := strconv.ParseInt(port, 10, 64)

	if err != nil {
		log.Printf("port error: %s\n", err.Error())
	}

	//for rdp, we don't know the key or namespace at this point
	svc := service.NewMembershipService("", "")
	svr, err := rpc.NewServer("Membership", svc, iport)
	for {
		svr.Accept()
	}

	if err != nil {
		log.Fatal("server error:", err)
	}
}
