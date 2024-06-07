package postgres

import (
	"database/sql"
	"fmt"
	"my_project/model"
)

type FacultyRepository struct {
	Db *sql.DB
}

func NewFacultyRepository(db *sql.DB) *FacultyRepository {
	return &FacultyRepository{
		Db: db,
	}
}

func (ft *FacultyRepository) CreateFaculty(faculty model.Faculty) error {
	_, err := ft.Db.Exec("insert into  faculties(faculty_name,group_count,decant,faculty_email,faculty_phone) values ($1,$2,$3,$4,$5)", faculty.Name, faculty.GroupCount, faculty.Decant, faculty.FacultyEmail, faculty.FacultyPhone)
	if err != nil {
		return err
	}
	return nil
}

func (ft *FacultyRepository) UpdateFaculty(id string, faculty model.Faculty) error {
	_, err := ft.Db.Exec("update faculties set faculty_name=$1,group_count=$2,decant=$3,faculty_email=$4,faculty_phone =$5 where id=$6", faculty.Name, faculty.GroupCount, faculty.Decant, faculty.FacultyEmail, faculty.FacultyPhone, id)
	if err != nil {
		fmt.Println("salom")
		return err
	}
	return nil
}

func (ft *FacultyRepository) DeleteFaculty(id string) error {
	_, err := ft.Db.Exec("delete  from  faculties where id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func (ft *FacultyRepository) ReadAllFaculty() ([]model.Faculty, error) {
	rows, err := ft.Db.Query("select id,faculty_name,group_count,decant,faculty_email,faculty_phone,created_at,updated_at,deleted_at from faculties")
	faculties := []model.Faculty{}
	if err != nil {
		return nil, err
	}
	faculty := model.Faculty{}
	for rows.Next() {
		err := rows.Scan(&faculty.Id, &faculty.Name, &faculty.GroupCount, &faculty.Decant, &faculty.FacultyEmail, &faculty.FacultyPhone, &faculty.CreatedAt, &faculty.UpdatedAt, &faculty.DeletedAt)

		if err != nil {
			fmt.Println("has", err)
			return nil, err
		}

		faculties = append(faculties, faculty)
	}
	fmt.Println("salom", faculties)

	return faculties, err
}
