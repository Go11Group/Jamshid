package postgres

import (
	"database/sql"
	"fmt"
	"my_project/model"
)

type UniversityRepository struct {
	Db *sql.DB
}

func NewUniversityRepository(db *sql.DB) *UniversityRepository {
	return &UniversityRepository{
		Db: db,
	}
}

func (un *UniversityRepository) CreateUniversity(university model.University) error {
	_, err := un.Db.Exec("insert into  universities(university_name,rector,university_email) values ($1,$2,$3)", university.UniversityName, university.Rector, university.UniversityEmail)
	if err != nil {
		return err
	}
	return nil
}

func (un *UniversityRepository) UpdateUniversity(id string, university model.University) error {
	_, err := un.Db.Exec("update universities set university_name=$1,rector=$2,university_email=$3 where id=$4", university.UniversityName, university.Rector, university.UniversityEmail, id)
	if err != nil {
		return err
	}
	return nil
}

func (un *UniversityRepository) DeleteUniversity(id string) error {
	_, err := un.Db.Exec("delete  from  universities where id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func (un *UniversityRepository) ReadAllUniversity() ([]model.University, error) {
	rows, err := un.Db.Query("select id , university_name,rector, university_email ,created_at,updated_at,deleted_at  from universities")
	universities := []model.University{}
	if err != nil {
		fmt.Println("edds", err)
		return nil, err
	}
	university := model.University{}
	for rows.Next() {
		err := rows.Scan(&university.Id, &university.UniversityName, &university.Rector, &university.UniversityEmail, &university.CreatedAt, &university.UpdatedAt, &university.DeletedAt)

		if err != nil {
			fmt.Println("has", err)
			return nil, err
		}

		universities = append(universities, university)
	}
	fmt.Println("salom", universities)

	return universities, err
}
