package postgres

import (
	"database/sql"
	"my_project/model"
)

type StudentRepository struct {
	Db *sql.DB
}

func NewStudentRepository(db *sql.DB) *StudentRepository {
	return &StudentRepository{Db: db}
}

func (db *StudentRepository) InsertIntoStudent(student model.Student) error {
	_, err := db.Db.Exec("insert into student (id,name,age) values ($1,$2,$3)", student.Id, student.Name, student.Age)
	if err != nil {
		return err
	}
	return nil

}

func (db *StudentRepository) UpdateStudent(id string, student model.Student) error {
	_, err := db.Db.Exec("update student set id=$1 ,name=$2,age=$3 where  id=$4", student.Id, student.Name, student.Age, id)
	if err != nil {
		return err
	}
	return nil

}

func (db *StudentRepository) ReadStudent() ([]model.Student, error) {
	students := []model.Student{}
	rows, err := db.Db.Query("select * from student")

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		student := model.Student{}
		rows.Scan(&student.Id, &student.Name, &student.Age)
		students = append(students, student)
	}
	return students, nil
}
func (db *StudentRepository) DeleteStudent(id string) error {
	_, err := db.Db.Exec("delete from student where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
