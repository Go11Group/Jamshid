package main

import (
	"fmt"
	"github.com/go-faker/faker/v4"
	_ "github.com/go-faker/faker/v4"
	"my_project/storage/postgres"
)

func main() {
	db, err := postgres.ConnectionDb()
	if err != nil {
		panic(err)
	}
	fmt.Println("Is success ")
	for i := 0; i < 100000; i++ {
		_, err := db.Exec(""+
			"insert into students(name,surname,phone_number,email,password) values ($1,$2,$3,$4,$5)", faker.FirstName(), faker.LastName(), faker.PhoneNumber, faker.Email(), faker.Password())
		if err != nil {
			panic(err)
		}
		if i%1000 == 0 {
			fmt.Println(i)
		}
	}
}
