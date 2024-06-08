package controller

import (
	"github.com/gorilla/mux"
	"my_project/handler"
	"net/http"
)

func RestApi() {
	router := mux.NewRouter()
	sr := router.PathPrefix("/api/user").Subrouter()

	sr.HandleFunc("/create", handler.CreateUserHandler).Methods(http.MethodPost)
	sr.HandleFunc("/get", handler.GetAllUserHandler).Methods(http.MethodGet)
	sr.HandleFunc("/delete/{id}", handler.DeleteUserHandler).Methods(http.MethodDelete)
	sr.HandleFunc("/update/{id}", handler.UpdateUserHandler).Methods(http.MethodPut)

	gr := router.PathPrefix("/api/problem").Subrouter()
	gr.HandleFunc("/create", handler.CreateProblemHandler).Methods(http.MethodPost)
	gr.HandleFunc("/get", handler.GetAllProblemHandler).Methods(http.MethodGet)
	gr.HandleFunc("/delete/{id}", handler.DeleteProblemHandler).Methods(http.MethodDelete)
	gr.HandleFunc("/update/{id}", handler.UpdateProblemHandler).Methods(http.MethodPut)

	grs := router.PathPrefix("/api/problem_solve").Subrouter()
	grs.HandleFunc("/create", handler.CreateSolvingProblemHandler).Methods(http.MethodPost)
	grs.HandleFunc("/get", handler.GetAllSolvingProblemHandler).Methods(http.MethodGet)
	grs.HandleFunc("/delete/{id}", handler.DeleteSolvingProblemHandler).Methods(http.MethodDelete)
	grs.HandleFunc("/update/{id}", handler.DeleteSolvingProblemHandler).Methods(http.MethodPut)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}

}
