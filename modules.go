package main

import (
	"fmt"
	"net/http"
)

func getModulesFunction(w http.ResponseWriter, r *http.Request) {
	var message = "Hello from modules script, you method was: " + r.Method
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

func startModulesService() {
	fmt.Println("Print en consola desde modules")
	http.HandleFunc("/modules", getModulesFunction)
}
