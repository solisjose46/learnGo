package main

import (
	"log"
	"net/http"
)

// Notes
// Go auto sets response for the following: Date, Content-length and Content-type
// but should manually set content-type for json

// Del() does not suppress header, to supress header use nil ex writer.Header()["Date"] = nil

func main(){
	mux := http.NewServeMux() // Defailt Mux is global and can be abused by adding rotes by compromised packages
	// subtree path / is like /*
	// path must end in slash
	mux.HandleFunc("/", home)

	// fixed path does not end in slash
	// looks for exact path
	mux.HandleFunc("/test", test)

	serverError := http.ListenAndServe(":8080", mux)

	log.Fatal(serverError)
}