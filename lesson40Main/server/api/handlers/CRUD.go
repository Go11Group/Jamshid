package handlers

import (
	"encoding/json"
	"fmt"
	"my_project/models"
	"net/http"
	"strconv"
	"strings"
)

func (us *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		_, err := w.Write([]byte("Error"))
		if err != nil {
			return
		}
		return
	}
	err = us.UserRepo.CreateUser(user)
	if err != nil {
		fmt.Println("+++++++", err)
		_, err := w.Write([]byte("Error Internal server error"))
		if err != nil {
			return
		}
	}
}

func (us *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/user/update")
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		_, err := w.Write([]byte("Error"))
		if err != nil {
			return
		}
		return
	}
	err = us.UserRepo.UpdateUser(id, user)
	if err != nil {
		return
	}
}

func (us *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/user/update")
	err := us.UserRepo.DeleteUser(id)
	if err != nil {
		_, err := w.Write([]byte("Error"))
		if err != nil {
			return
		}
		return
	}
}
func (us *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	age := r.URL.Query().Get("age")
	son, err := strconv.Atoi(age)
	if err != nil {
		_, err := w.Write([]byte("Error"))
		if err != nil {
			return
		}
		return
	}

	filter := models.Filter{
		Name:  r.URL.Query().Get("name"),
		Age:   son,
		Email: r.URL.Query().Get("email"),
	}
	users, err := us.UserRepo.Get(filter)
	if err != nil {
		_, err := w.Write([]byte("Error"))
		if err != nil {
			return
		}
	}
	err = json.NewEncoder(w).Encode(&users)
	if err != nil {
		_, err := w.Write([]byte("Error"))
		if err != nil {
			return
		}
		return
	}

}
