package main

import (
	"fmt"
	"net/http"
)

func getUsersFunction(w http.ResponseWriter, r *http.Request) {
	var message = "Hello from users script, you method was: " + r.Method
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

func startUsersService() {
	fmt.Println("Print en consola desde users")
	http.HandleFunc("/users", getUsersFunction)
}
