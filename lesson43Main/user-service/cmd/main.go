package main

import (
	"user_service/api"
	config2 "user_service/config"
	"user_service/strorage"
	"user_service/strorage/postgres"
)

func main() {
	config := config2.Load()
	db, err := strorage.ConnectionDb(config)
	if err != nil {
		panic(err)
	}
	s := postgres.NewUserRepository(db)
	router := api.RouterApi(s)
	err = router.Run(config.HTTPPort)
	if err != nil {
		panic(err)
	}
}
