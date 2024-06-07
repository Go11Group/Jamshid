package controller

import (
	"my_project/Service"
	"net/http"
)

func FacultyController(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/faculty/all", Service.Faculty_ReadAll_Handler)
	mux.HandleFunc("POST /api/faculty/create", Service.Create_Faculty_Handler)
	mux.HandleFunc("PUT /api/faculty/update", Service.Update_Faculty_Handler)
	mux.HandleFunc("DELETE /api/faculty/delete", Service.Delete_Faculty_Handler)

}
