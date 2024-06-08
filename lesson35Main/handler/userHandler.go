package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"my_project/model"
	"my_project/storage/postgres"
	"net/http"
)

func UserRepo() *postgres.UserRepository {
	db, err := postgres.ConnectionDB()
	if err != nil {
		panic(err)
	}
	userHand := postgres.UserRepository{}
	userHand.Db = db
	return &userHand
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	studentHand := UserRepo()
	user := model.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return
	}
	err = studentHand.CreateUser(user)
	if err != nil {
		fmt.Println("User create is not ")
		_, err = w.Write([]byte("User is not  created"))
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Println("User create is ")
	_, err = w.Write([]byte("User created"))
	w.WriteHeader(http.StatusOK)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {

	studentHand := UserRepo()
	param := mux.Vars(r)
	fmt.Println(param["id"])
	id := param["id"]
	err := studentHand.DeleteUser(id)
	if err != nil {
		fmt.Println("user is not deleted")
		_, err = w.Write([]byte("Is  not success deleted user"))
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Println("user is deleted ")
	_, err = w.Write([]byte("Is   success deleted user"))
	w.WriteHeader(http.StatusOK)

}
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	studentHand := UserRepo()
	param := mux.Vars(r)
	id := param["id"]
	user := model.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
	}
	err = studentHand.UpdatedUser(id, user)
	if err != nil {
		_, err = w.Write([]byte("is user  not updated "))
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(" user  is  updated "))

}

func GetAllUserHandler(w http.ResponseWriter, r *http.Request) {
	studentHand := UserRepo()
	a, err := studentHand.GetAllUser()
	if err != nil {
		panic(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Println("a========", a)
	users, err := json.Marshal(a)
	fmt.Println("user   =====", users)
	if err != nil {
		fmt.Println("Marshal errr-r---")
		w.WriteHeader(http.StatusBadRequest)
	}
	_, err = w.Write(users)
	if err != nil {
		_, err = w.Write([]byte("is user  not getAll "))
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)

}
