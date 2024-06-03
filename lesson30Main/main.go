package main

import (
	"database/sql"
	"fmt"
	_ "github.com/gorilla/mux"
	"lesson1/storage/postgres"
	_ "log"
	_ "net/http"
)

func main() {
	db, err := postgres.ConnectionDB()
	if err != nil {
		panic(err)

	}
	fmt.Println("is success")
	customerInfo := postgres.NewRepository(db)
	//customer := model.Customer{Username: "Jamshid", Email: "hatamovjamshid47@gmail.com", Password: "12345678"}
	//err = customerInfo.Create(customer)
	//if err != nil {
	//	panic(err)
	//}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	//find alll
	//a, err := customerInfo.FindAll()
	//if err != nil {
	//	panic(err)
	//}
	//for i := 0; i < len(a); i++ {
	//	fmt.Println(a[i])
	//}

	///deletetd

	//err = customerInfo.DeleteCustomer(1)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("deleted")
	//
	//customer := model.Customer{}
	//customer.Username = "Olim"
	//customer.Email = "Olimov"
	//customer.Password = "1234"
	//err = customerInfo.UpdatedCustomer(2, &customer)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("is updated")

	//product alll
	productInfo := postgres.NewProductRepository(db)
	products, err := productInfo.FindAllProduct()
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(products); i++ {
		fmt.Println(products[i].Name)
		fmt.Println(products[i].Price)
	}
	custumers, err := customerInfo.FindAllCustomer()
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(custumers); i++ {
		fmt.Println(custumers[i].Username)
	}

}
