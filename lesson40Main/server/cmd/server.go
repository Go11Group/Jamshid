package main

import (
	"database/sql"
	"fmt"
	"my_project/api"
	"my_project/storage/postgres"
	"net/http"
)

func main() {
	db, err := postgres.ConnectionDb()
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	fmt.Println("Connect!!!")
	user := postgres.NewUserRepository(db)
	router := api.RooterApi(user)
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
	fmt.Println("Listennig :8080")

}
