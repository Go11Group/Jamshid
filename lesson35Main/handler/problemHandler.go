package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"my_project/model"
	"my_project/storage/postgres"
	"net/http"
)

func problemRepo() *postgres.ProblemRepository {
	db, err := postgres.ConnectionDB()
	if err != nil {
		panic(err)
	}
	problemHand := postgres.ProblemRepository{}
	problemHand.Db = db
	return &problemHand
}

func CreateProblemHandler(w http.ResponseWriter, r *http.Request) {

	problemHand := problemRepo()
	problem := model.Problem{}
	err := json.NewDecoder(r.Body).Decode(&problem)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		_, err = w.Write([]byte("problem   not created"))

	}
	err = problemHand.CreateProblem(problem)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
	}
	_, err = w.Write([]byte("problem    created"))

}

func DeleteProblemHandler(w http.ResponseWriter, r *http.Request) {
	problemHand := problemRepo()
	param := mux.Vars(r)
	id := param["id"]
	fmt.Println(id)
	err := problemHand.DeleteProblem(id)
	if err != nil {
		_, err = w.Write([]byte("problem is  not deleted"))
		w.WriteHeader(http.StatusBadRequest)
	}
	_, err = w.Write([]byte("problem    deleted"))

	w.WriteHeader(http.StatusOK)

}
func UpdateProblemHandler(w http.ResponseWriter, r *http.Request) {
	problemHand := problemRepo()
	param := mux.Vars(r)
	id := param["id"]
	problem := model.Problem{}
	err := json.NewDecoder(r.Body).Decode(&problem)
	if err != nil {
		_, err = w.Write([]byte("problem   not updated"))

		w.WriteHeader(http.StatusInternalServerError)
	}
	err = problemHand.UpdatedProblem(id, problem)
	if err != nil {
		_, err = w.Write([]byte("problem   not updated"))

		w.WriteHeader(http.StatusBadRequest)
	}
	_, err = w.Write([]byte("problem    updated"))

	w.WriteHeader(http.StatusOK)

}

func GetAllProblemHandler(w http.ResponseWriter, r *http.Request) {
	problemHand := problemRepo()
	a, err := problemHand.GetAllProblem()
	if err != nil {
		panic(err)
	}
	problems, err := json.Marshal(a)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	_, err = w.Write(problems)
	if err != nil {
		_, err = w.Write([]byte("problem   is getAll "))

		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)

}
