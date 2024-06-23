package main

import (
	"bill_service/api"
	config2 "bill_service/config"
	strorage "bill_service/storage"
	"bill_service/storage/postgres"
)

func main() {
	config := config2.Load()
	db, err := strorage.ConnectionDb(config)
	if err != nil {
		panic(err)
	}
	router := api.RouterApi(postgres.NewCardRepository(db), postgres.NewStationRepository(db), postgres.NewTerminalRepository(db), postgres.NewTransactionRepository(db))

	panic(router.Run(config.HTTPPort))

}
