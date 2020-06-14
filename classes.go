package main

import (
	"fmt"
	"net/http"
)

func getClassesFunction(w http.ResponseWriter, r *http.Request) {
	var message = "Hello from classes script, you method was: " + r.Method
	w.Write([]byte(message))

	switch method := r.Method; method {
	case "GET":
		fmt.Println("get")
	case "PUT":
		fmt.Println("PUT")
	case "POST":
		fmt.Println("POST")
	case "DELETE":
		fmt.Println("DELETE")
	default:
		fmt.Printf("%s.\n", method)
	}
}

func startClassesService() {
	fmt.Println("Print en consola desde classes")
	http.HandleFunc("/classes", getClassesFunction)
}