package Service

import (
	"encoding/json"
	"fmt"
	_ "fmt"
	_ "github.com/gorilla/mux"
	"my_project/model"
	"my_project/storage/postgres"
	"net/http"
)

func Group_Handler_Repo() *postgres.GroupRepository {
	db, err := postgres.ConnectionDb()
	if err != nil {
		panic(err)
	}
	groupInfo := postgres.NewGroupRepository(db)
	return groupInfo
}
func Group_ReadAll_Handler(w http.ResponseWriter, r *http.Request) {
	studentInfo := Group_Handler_Repo()
	groups, err := studentInfo.ReadAllGroup()
	fmt.Println(groups)
	data, err := json.Marshal(groups)
	if err != nil {
		_, err = w.Write([]byte(string(rune(http.StatusInternalServerError))))
	}
	_, err = w.Write(data)

}

func Group_Deleted_Handler(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query()
	id := param["id"][0]

	groupInfo := Group_Handler_Repo()
	err := groupInfo.DeleteGroup(id)
	if err != nil {
		http.Error(w, "Group is not  deleted", http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte("Group is deleted"))
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
	}
}

func Create_Group_Handler(w http.ResponseWriter, r *http.Request) {

	studentInfo := Group_Handler_Repo()
	group := model.Group{}
	err := json.NewDecoder(r.Body).Decode(&group)
	fmt.Println(group)
	err = studentInfo.CreateGroup(group)
	if err != nil {
		_, err = w.Write([]byte("Is   is not create group"))
	} else {
		_, err = w.Write([]byte("Is  create group"))
	}
}

func Update_Group_Handler(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query()
	id := param["id"][0]
	groupInfo := Group_Handler_Repo()

	group := model.Group{}
	err := json.NewDecoder(r.Body).Decode(&group)
	fmt.Println(group)
	if err != nil {
		_, err = w.Write([]byte("Internal Server Exception"))
	}

	err = groupInfo.UpdateGroup(id, group)
	if err != nil {
		_, err = w.Write([]byte(string(rune(http.StatusBadRequest))))
	}
	_, err = w.Write([]byte("Group is updated"))
}
