package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	//http.HandleFunc("/", sayHello)
	var port string = os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// starts http classes service
	startClassesService()

	// starts http Modules service
	startModulesService()

	// starts http Courses service
	startCoursesService()

	// starts http Users service
	startUsersService()

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
