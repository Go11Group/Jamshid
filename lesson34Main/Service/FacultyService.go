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

func Faculty_Handler_Repo() *postgres.FacultyRepository {
	db, err := postgres.ConnectionDb()
	if err != nil {
		panic(err)
	}
	facultyInfo := postgres.NewFacultyRepository(db)
	return facultyInfo
}
func Faculty_ReadAll_Handler(w http.ResponseWriter, r *http.Request) {
	facultyInfo := Faculty_Handler_Repo()
	faculties, err := facultyInfo.ReadAllFaculty()
	fmt.Println(faculties)
	data, err := json.Marshal(faculties)
	if err != nil {
		_, err = w.Write([]byte(string(rune(http.StatusInternalServerError))))
	}
	_, err = w.Write(data)

}

func Delete_Faculty_Handler(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query()
	id := param["id"][0]
	if id == "" {
		http.Error(w, "ID is Empty", http.StatusBadRequest)
		return
	}
	fmt.Println("Id ", id)

	facultyInfo := Faculty_Handler_Repo()
	err := facultyInfo.DeleteFaculty(id)
	if err != nil {
		fmt.Println("salom", err)
		http.Error(w, "faculty is not  deleted", http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte("faculty is deleted"))
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func Create_Faculty_Handler(w http.ResponseWriter, r *http.Request) {

	facultyInfo := Faculty_Handler_Repo()
	faculty := model.Faculty{}
	err := json.NewDecoder(r.Body).Decode(&faculty)
	fmt.Println(faculty)
	err = facultyInfo.CreateFaculty(faculty)
	if err != nil {
		fmt.Println("sealom", err)
		_, err = w.Write([]byte("Is not faculty students"))
	} else {
		_, err = w.Write([]byte("Is  create faculty"))
	}
}

func Update_Faculty_Handler(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query()
	id := param["id"][0]
	facultyInfo := Faculty_Handler_Repo()

	faculty := model.Faculty{}
	err := json.NewDecoder(r.Body).Decode(&faculty)
	fmt.Println(faculty)
	if err != nil {
		_, err = w.Write([]byte("Faculty is not updated"))
	}

	err = facultyInfo.UpdateFaculty(id, faculty)
	_, err = w.Write([]byte("Faculty is  updated"))
}
