package controller

import (
	"my_project/Service"
	"net/http"
)

func StudentController(mux *http.ServeMux) {

	mux.HandleFunc("GET /api/student/all", Service.Student_ReadAll_Handler)
	mux.HandleFunc("POST /api/student/create", Service.Create_Student_Handler)
	mux.HandleFunc("PUT /api/student/update", Service.Update_Student_Handler)
	mux.HandleFunc("DELETE /api/student/delete", Service.Delete_Student_Handler)

}
