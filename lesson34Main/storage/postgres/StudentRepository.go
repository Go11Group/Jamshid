package postgres

import (
	"database/sql"
	"fmt"
	"my_project/model"
)

type StudentRepository struct {
	Db *sql.DB
}

func NewStudentRepository(db *sql.DB) *StudentRepository {
	return &StudentRepository{
		Db: db,
	}
}

func (st *StudentRepository) CreateStudent(student model.Student) error {
	_, err := st.Db.Exec("insert into  students(first_name,last_name,age,email,phone_number) values ($1,$2,$3,$4,$5)", student.FirstName, student.LastName, student.Age, student.Email, student.PhoneNumber)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (st *StudentRepository) UpdateStudent(id string, student *model.Student) error {
	_, err := st.Db.Exec("update students set first_name=$1,last_name=$2,age=$3,email=$4 where id=$5", student.FirstName, student.LastName, student.Age, student.Email, id)
	if err != nil {
		return err
	}
	return nil
}

func (st *StudentRepository) DeleteStudent(id string) error {
	_, err := st.Db.Exec("delete  from  students where id=$1", id)
	if err != nil {
		fmt.Println("is hello")
		return err
	}
	return nil
}

func (st *StudentRepository) ReadAllStudent() ([]model.Student, error) {
	rows, err := st.Db.Query("select id,first_name,last_name,age ,email,phone_number from students")
	students := []model.Student{}
	if err != nil {
		fmt.Println("student into error")
		return nil, err
	}
	student := model.Student{}
	for rows.Next() {
		err := rows.Scan(&student.Id, &student.FirstName, &student.LastName, &student.Age, &student.Email, &student.PhoneNumber)

		if err != nil {
			fmt.Println("has", err)
			return nil, err
		}

		students = append(students, student)
	}
	fmt.Println("salom", students)

	return students, err
}
