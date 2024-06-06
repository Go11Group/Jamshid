package Service

import (
	"encoding/json"
	"fmt"
	_ "fmt"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"my_project/model"
	"my_project/storage/postgres"
	"net/http"
)

func Universitiy_Handler_Repo() *postgres.UniversityRepository {
	db, err := postgres.ConnectionDb()
	if err != nil {
		panic(err)
	}
	univerInfo := postgres.NewUniversityRepository(db)
	return univerInfo
}
func Universitiy_ReadAll_Handler(w http.ResponseWriter, r *http.Request) {
	univerInfo := Student_Handler_Repo()
	universities, err := univerInfo.ReadAllStudent()
	fmt.Println(universities)
	data, err := json.Marshal(universities)
	if err != nil {
		_, err = w.Write([]byte(string(rune(http.StatusInternalServerError))))
	}
	_, err = w.Write(data)

}

func Universitiy_Delete_Handler(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id := param["id"]

	univerInfo := Universitiy_Handler_Repo()
	err := univerInfo.DeleteUniversity(id)
	if err != nil {
		fmt.Println("salom", err)
		http.Error(w, "Student is not  deleted", http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte("Student is deleted"))
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func Create_Universitiy_Handler(w http.ResponseWriter, r *http.Request) {

	univerInfo := Universitiy_Handler_Repo()
	universities := model.University{}
	err := json.NewDecoder(r.Body).Decode(&universities)
	fmt.Println(universities)
	err = univerInfo.CreateUniversity(universities)
	if err != nil {
		fmt.Println("sealom", err)
		_, err = w.Write([]byte("Is not create students"))
	} else {
		_, err = w.Write([]byte("Is  create students"))
	}
}

func Update_Universitiy_Handler(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id := param["id"]
	univerInfo := Universitiy_Handler_Repo()

	university := model.University{}
	err := json.NewDecoder(r.Body).Decode(&university)
	fmt.Println(university)
	if err != nil {
		_, err = w.Write([]byte("Internal Server Exception"))
	}

	err = univerInfo.UpdateUniversity(id, university)
}
