package controller

import (
	"my_project/Service"
	"net/http"
)

func GroupController(mux *http.ServeMux) {

	mux.HandleFunc("GET /api/group/all", Service.Group_ReadAll_Handler)
	mux.HandleFunc("POST /api/group/create", Service.Create_Group_Handler)
	mux.HandleFunc("PUT /api/group/update", Service.Update_Group_Handler)
	mux.HandleFunc("DELETE /api/group/delete", Service.Group_Deleted_Handler)

}
