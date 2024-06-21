package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	_ "github.com/gin-gonic/gin"
	"my_project/model"
	"net/http"
	"strconv"
	"strings"
)

/* bu yerda create user qiladi    */
func (handle *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	err := json.NewDecoder(r.Body).Decode(&user) // blindjson - bu body kelgan malumotni jsonga parse qilib
	if err != nil {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
	}
	userJson, err := json.Marshal(&user)
	if err != nil {

		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
		return
	}
	s, err := http.NewRequest(http.MethodPost, "http://localhost:8080/api/user/create", bytes.NewBuffer(userJson))
	if err != nil {

		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
		return
	}
	//err = handle.UserRepo.CreateUser(user) // create user sql -query code ga otadi  user struct ti berib yuboriladi
	response, err := handle.client.Do(s)
	if err != nil {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		} // agar bazaga saqlansamasa response=Ok
		return
	}
	if response.StatusCode == 200 || response.StatusCode == 201 {
		_, err := w.Write([]byte("Ok success"))
		if err != nil {
			return
		}
	} else {
		fmt.Println("----------", userJson)

		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
	}

}

/* bu yerda updated user qiladi */
func (handle *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	err := json.NewDecoder(r.Body).Decode(&user) // blindjson - bu body kelgan malumotni jsonga parse qilib
	if err != nil {
		fmt.Println("+++++++", err)
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
	}
	userJson, err := json.Marshal(&user)
	if err != nil {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
		return
	}
	id := strings.TrimPrefix(r.URL.Path, "/api/user/update/")
	s, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:8080/api/user/update/%s", id), bytes.NewBuffer(userJson))
	//err = handle.UserRepo.UpdateUser(gn.Param("id"), user)
	if err != nil {
		fmt.Println("+++++++++", err)
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		} // agar bazaga saqlansamasa response=internalservererror
	}
	response, err := handle.client.Do(s)
	if err != nil {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
		return
	}
	if response.StatusCode == 200 {
		_, err := w.Write([]byte("Ok success "))
		if err != nil {
			return
		}
	} else {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
	}
}

/* bu yerda deleted user qiladi */

func (handle *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/api/user/delete/")

	//err := handle.UserRepo.DeleteUser(gn.Param("id"))
	s, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:8080/api/user/delete/%s", id), nil)
	if err != nil {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		} // agar bazaga saqlansamasa response=internalservererror
	}

	response, err := handle.client.Do(s)
	if err != nil {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
		return
	}
	if response.StatusCode == 200 {
		_, err := w.Write([]byte("Ok success "))
		if err != nil {
			return
		}
	} else {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
	}
}

/* bu yerda filter va getAll user qiladi */
func (handle *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	userFilter := model.Filter{}
	userFilter.Name = r.URL.Query().Get("name")
	userFilter.Email = r.URL.Query().Get("email")
	userFilter.Birthday = r.URL.Query().Get("birthday")
	userFilter.Password = r.URL.Query().Get("password")
	Limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	Offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	fmt.Println("-------------", r.URL.Query().Get("offset"))
	fmt.Println("-------------", Offset)
	//err := gn.ShouldBindQuery(userFilter)

	userFilter.Limit = Limit
	userFilter.Offset = Offset
	//userFilterJson, err := json.Marshal(&userFilter)
	if err != nil {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
		return
	}
	s, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/user/get/?name=%s&email=%s&birthday=%s&password=%s&limit=%d&offset=%d", &userFilter.Name, userFilter.Email, &userFilter.Password, &userFilter.Limit, &userFilter.Offset), nil)
	//users, err := handle.UserRepo.GetUser(userFilter)
	if err != nil {
		fmt.Println("+++++++++++", err)
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}

	}
	client := http.Client{}
	response, err := client.Do(s)
	if err != nil {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
	}
	if response.StatusCode == 200 {
		users := []model.User{}
		err = json.NewDecoder(response.Body).Decode(&users)
		if err != nil {
			_, err := w.Write([]byte("error internal server  error"))
			if err != nil {
				return
			}
			return
		}
		json.NewEncoder(w).Encode(&users)
	} else {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
	}

}

/* bu yerda search user_id boyicha  user qiladi */
func (handle *Handler) GetCourseByUserId(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/user/courses/")
	s, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/user/courses/%s", id), nil)
	//userId, courses, err := handle.UserRepo.GetEnrollmentByCourseId(gn.Param("id"))
	if err != nil {
		fmt.Println("+++++++++++", err)
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
	}
	response, err := handle.client.Do(s)
	if err != nil {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
		return
	}
	if response.StatusCode == 200 {
		courses := []model.Course{}
		err = json.NewDecoder(response.Body).Decode(&courses)
		if err != nil {
			_, err := w.Write([]byte("error internal server  error"))
			if err != nil {
				return
			}
			return
		}
		json.NewEncoder(w).Encode(&courses)
	} else {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
	}

}

/* bu yerda Email yoki Name boyicha qidiradi user qiladi */

func (handle *Handler) GetUserByEmailOrName(w http.ResponseWriter, r *http.Request) {
	//users, err := handle.UserRepo.GetUserByEmailOrName(gn.Param("name"), gn.Param("email"))
	name := r.URL.Query().Get("name")
	email := r.URL.Query().Get("email")
	s, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/user/get/%s/%s", name, email), nil)
	if err != nil {
		fmt.Println("+++++++++++", err)
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
	}
	response, err := handle.client.Do(s)
	if err != nil {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
		return
	}
	if response.StatusCode == 200 {
		users := []model.User{}
		err := json.NewDecoder(response.Body).Decode(&users)
		if err != nil {
			_, err := w.Write([]byte("error internal server  error"))
			if err != nil {
				return
			}
			return
		}
		json.NewEncoder(w).Encode(&users)
	} else {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
	}
}

// user id boyich search qiladi

func (handle *Handler) GetUserById(w http.ResponseWriter, r *http.Request) {
	//user, err := handle.UserRepo.GetById(gn.Param("id"))
	id := strings.TrimPrefix(r.URL.Path, "/api/user/id/")
	s, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/api/user/id/%s", id), nil)
	if err != nil {
		fmt.Println("+++++++++++", err)
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
	}
	response, err := handle.client.Do(s)
	if err != nil {
		return
	}
	if response.StatusCode == 200 {
		users := []model.User{}
		err = json.NewDecoder(response.Body).Decode(&users)
		if err != nil {
			_, err := w.Write([]byte("error internal server  error"))
			if err != nil {
				return
			}
			return
		}
		json.NewEncoder(w).Encode(&users)
	} else {
		_, err := w.Write([]byte("error internal server  error"))
		if err != nil {
			return
		}
	}

}
