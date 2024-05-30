package main

import (
	"fmt"
	"my_project/storage/postgres"
)

func main() {
	db, err := postgres.ConnectionDB()
	if err != nil {
		panic(err)
	}
	postgres.NewStudentRepository(db)

	// uztoz bu student table va course table referance qilingan many to many qilib shunining uchun course cruddida course_student table shu crudlar bajariladi

	//Read

	//studentInfo := postgres.NewStudentRepository(db)
	//students, err := studentInfo.ReadStudent()
	//if err != nil {
	//	panic(err)
	//}
	//for _, student := range students {
	//	fmt.Println(student)
	//}

	/// update

	//student := model.Student{"16cb01a2-179e-4053-addc-f9c88d27309a", "Ahmad", 12}
	//err = studentInfo.UpdateStudent("0a7ecf43-b73b-4dc9-9d7a-7b1611cead5c", student)
	//if err != nil {
	//	fmt.Println(fmt.Errorf("update student error: %v", err))
	//	panic(err)
	//}
	//fmt.Println("student updated")

	// delete

	//err = studentInfo.DeleteStudent("16cb01a2-179e-4053-addc-f9c88d27309a")
	//if err != nil {
	//	fmt.Println(fmt.Errorf("delete student error: %v", err))
	//	panic(err)
	//}
	//fmt.Println("student deleted")

	// insert into

	//err = studentInfo.InsertIntoStudent(student)
	//if err != nil {
	//	fmt.Println(fmt.Errorf("insert student error: %v", err))
	//	panic(err)
	//}
	//fmt.Println("student inserted")

	//courseInfo := postgres.CourseRepository{db}

	// read course
	//courses, err := courseInfo.ReadCourse()

	//if err != nil {
	//	fmt.Println(fmt.Errorf("is not found course"))
	//	panic(err)
	//}
	//for _, course := range courses {
	//	fmt.Println(course)
	//}

	//err = courseInfo.DeleteCourse("28c2f838-ac98-4086-8689-ee35b99a4167")
	//if err != nil {
	//	fmt.Println(fmt.Errorf("is not deleted course"))
	//	panic(err)
	//}
	//fmt.Println("deleted course")
	//course := model.Course{Id: "d91f67b4-8baa-43a7-a9cd-bc848662e3cd", CourseName: "Foundation1"}
	//courseStudent := model.CourseStudent{"54987563-0965-424c-93b3-32d3e9ca8492", "16cb01a2-179e-4053-addc-f9c88d27309a", "16cb01a2-179e-4053-addc-f9c88d27309a"}
	//

	// insert

	//err = courseInfo.InsertIntoCourse(course, courseStudent)
	//if err != nil {
	//	fmt.Println(fmt.Errorf("is not inserted course"))
	//	panic(err)
	//}
	//fmt.Println("inserted course")
	//
	//

	// updated
	//err = courseInfo.UpdateCourse()
	if err != nil {
		fmt.Println(fmt.Errorf("error updating course: %v", err))
		panic(err)
	}
	fmt.Println("is updated successfully")

}
