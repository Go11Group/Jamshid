package controller

import (
	"my_project/Service"
	"net/http"
)

func UniversityController(mux *http.ServeMux) {

	mux.HandleFunc("GET /api/university/all", Service.Universitiy_ReadAll_Handler)
	mux.HandleFunc("POST /api/university/create", Service.Create_Universitiy_Handler)
	mux.HandleFunc("PUT /api/university/update", Service.Update_Universitiy_Handler)
	mux.HandleFunc("DELETE /api/university/delete", Service.Universitiy_Delete_Handler)

}
