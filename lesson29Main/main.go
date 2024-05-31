package main

import (
	"fmt"
	"lesson29Main/storage/postgres"
)

func main() {
	db, err := postgres.ConnectionDb()
	if err != nil {
		panic(err)
	}
	fmt.Println("Is success")
	employeeInfo := postgres.NewWriterRepository(db)

	//insert int

	//writer := model.Writer{FirstName: "Nodir", LastName: "Nodirov", Email: "nodir@gmail.com", Password: "124", Age: 34, Field: "developer", Gender: "male", IsEmployee: false}
	//err = employeeInfo.InsetInto(writer)
	//if err != nil {
	//	fmt.Println(fmt.Errorf("insert not success: %v", err))
	//	panic(err)
	//}

	// delete
	//err = employeeInfo.DeleteId(2)
	//if err != nil {
	//	fmt.Println(fmt.Errorf("delete not success: %v", err))
	//	panic(err)
	//}
	//fmt.Println("delete is success ")

	//employees, err := employeeInfo.FindByIsEmployee(true)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(employees)

	//employees, err := employeeInfo.FindByGender("male")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(employees)

	//employees, err := employeeInfo.FindByAgeFromAgeTo(10, 15)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(employees)

	//employees, err := employeeInfo.FindALl(true)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(employees)

	//employees, err := employeeInfo.FindByName("Nodir")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(employees)
}
