package main

import (
	"my_project/controller"
	"net/http"
)

func main() {
	mux := http.ServeMux{}
	controller.StudentController(&mux)
	controller.UniversityController(&mux)
	controller.GroupController(&mux)
	controller.FacultyController(&mux)

	err := http.ListenAndServe(":8088", &mux)
	if err != nil {
		panic(err)
	}
}
