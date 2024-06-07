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

func Student_Handler_Repo() *postgres.StudentRepository {
	db, err := postgres.ConnectionDb()
	if err != nil {
		panic(err)
	}
	studentInfo := postgres.NewStudentRepository(db)
	return studentInfo
}
func Student_ReadAll_Handler(w http.ResponseWriter, r *http.Request) {
	studentInfo := Student_Handler_Repo()
	students, err := studentInfo.ReadAllStudent()
	fmt.Println(students)
	data, err := json.Marshal(students)
	fmt.Println(students)
	if err != nil {
		_, err = w.Write([]byte(string(rune(http.StatusInternalServerError))))
	}
	_, err = w.Write(data)

}

func Delete_Student_Handler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query()
	studentInfo := Student_Handler_Repo()
	fmt.Println(id["id"][0])
	err := studentInfo.DeleteStudent(id["id"][0])
	if err != nil {
		fmt.Println(" deleting student:", err)
		http.Error(w, "Student is not  deleted", http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte("Student has been deleted"))
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func Create_Student_Handler(w http.ResponseWriter, r *http.Request) {

	studentInfo := Student_Handler_Repo()
	student := model.Student{}
	err := json.NewDecoder(r.Body).Decode(&student)
	fmt.Println(student)
	err = studentInfo.CreateStudent(student)
	if err != nil {
		fmt.Println("sealom", err)
		_, err = w.Write([]byte("Is not create students"))
	} else {
		_, err = w.Write([]byte("Is  create students"))
	}
}

func Update_Student_Handler(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query()
	id := param["id"][0]
	fmt.Println(id)
	studentInfo := Student_Handler_Repo()

	student := model.Student{}
	err := json.NewDecoder(r.Body).Decode(&student)
	fmt.Println(student)
	if err != nil {
		_, err = w.Write([]byte("Internal Server Exception"))
	}

	err = studentInfo.UpdateStudent(id, &student)
}
