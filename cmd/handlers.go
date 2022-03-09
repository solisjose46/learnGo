package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func home(writer http.ResponseWriter, request *http.Request){
	// we want only to return this resource iff path matches /home

	if request.URL.Path != "/home"{
		http.NotFound(writer, request)
	}

	templateSet, err := template.ParseFiles("./ui/html/home.page.tmpl")

	if err != nil {
		log.Println(err.Error())
		http.Error(writer, "Internal Server Error", 500)
		return
	}

	// execute on template set writes template as response body
	err = templateSet.Execute(writer, nil)

	if err != nil {
		log.Println(err.Error())
		http.Error(writer, "Internal Server Error", 500)
	}


	writer.Write([]byte("Home"))
}

func test(writer http.ResponseWriter, request *http.Request){
	// want only post method
	if request.Method != http.MethodPost{
		writer.Header().Set("Allow", http.MethodPost)
		http.Error(writer, "Method Not Allowed", 405) // This route only supports POST request, does 
	}
	writer.Write([]byte("Home"))
}

// posting data via post and get (url parameter)
// for get route looks like /example?id=69
func postExample(writer http.ResponseWriter, request *http.Request){
	if request.Method == http.MethodGet{
		id, err := strconv.Atoi(request.URL.Query().Get("id"))
		if err != nil {
			http.NotFound(writer, request) // auto 404
		}
		fmt.Fprintf(writer, "ID is %d\n", id)
	}else if request.Method == http.MethodPost{
		// get value from body
	}


}