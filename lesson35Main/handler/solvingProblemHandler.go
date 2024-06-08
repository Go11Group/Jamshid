package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"my_project/model"
	"my_project/storage/postgres"
	"net/http"
)

func solve_problemRepo() *postgres.SolvedProblemRepository {
	db, err := postgres.ConnectionDB()
	if err != nil {
		panic(err)
	}
	problem_solve_hand := postgres.SolvedProblemRepository{}
	problem_solve_hand.Db = db
	return &problem_solve_hand
}

func CreateSolvingProblemHandler(w http.ResponseWriter, r *http.Request) {

	problemSolveHand := solve_problemRepo()
	solve_problem := model.SolvedProblem{}

	err := json.NewDecoder(r.Body).Decode(&solve_problem)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadGateway)
	}
	fmt.Println(solve_problem)
	err = problemSolveHand.CreateSolvedProblem(solve_problem)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadGateway)
	}

}

func DeleteSolvingProblemHandler(w http.ResponseWriter, r *http.Request) {
	problemSolveHand := solve_problemRepo()
	param := mux.Vars(r)
	id := param["id"]
	fmt.Println(id)
	err := problemSolveHand.DeleteSolvedProblem(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)

}
func UpdateSolvingProblemHandler(w http.ResponseWriter, r *http.Request) {
	problemSolveHand := solve_problemRepo()
	param := mux.Vars(r)
	id := param["id"]
	solve_problem := model.SolvedProblem{}
	err := json.NewDecoder(r.Body).Decode(&solve_problem)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	err = problemSolveHand.UpdatedSolvedProblem(id, solve_problem)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)

}

func GetAllSolvingProblemHandler(w http.ResponseWriter, r *http.Request) {
	problemSolveHand := solve_problemRepo()
	a, err := problemSolveHand.GetAllSolvedProblem()
	if err != nil {
		panic(err)
	}
	fmt.Println(a)
	solve_problems, err := json.Marshal(a)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	_, err = w.Write(solve_problems)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)

}
