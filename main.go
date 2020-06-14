package main

import (
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message + " with method " + r.Method

	w.Write([]byte(message))
}

func main() {
	//http.HandleFunc("/", sayHello)

	// starts http classes service
	startClassesService()

	// starts http Modules service
	startModulesService()

	// starts http Courses service
	startCoursesService()

	// starts http Users service
	startUsersService()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
