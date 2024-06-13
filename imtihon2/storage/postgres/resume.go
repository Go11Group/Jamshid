package postgres

import (
	"database/sql"
	"my_project/model"
)

type ResumeRepository struct {
	DB *sql.DB
}

func NewResumeRepository(db *sql.DB) *ResumeRepository {
	return &ResumeRepository{
		DB: db,
	}
}
func (ru *ResumeRepository) CreateResume(resume model.Resume) error {
	_, err := ru.DB.Exec("insert into  resumes(position ,experience,description,user_id) values ($1,$2,$3,$4)", &resume.Position, &resume.Experience, &resume.Description, &resume.UserID)
	return err
}
func (ru *ResumeRepository) UpdateResume(id string, resume model.Resume) error {
	_, err := ru.DB.Exec("update  resumes set position=$1 ,experience=$2,description=$3,user_id=$4 where id=$5 ", &resume.Position, &resume.Experience, &resume.Description, &resume.UserID, &id)
	return err
}
func (ru *ResumeRepository) DeletedResume(id string) error {
	_, err := ru.DB.Exec("delete from resumes where id=$1", &id)
	return err
}
func (ru *ResumeRepository) GetAllResume() ([]model.Resume, error) {
	rows, err := ru.DB.Query("select *from resumes")
	if err != nil {
		return nil, err
	}
	resumes := []model.Resume{}
	for rows.Next() {
		resume := model.Resume{}
		err = rows.Scan(&resume.Position, &resume.Experience, &resume.Description, &resume.UserID, &resume.CreatedAt, &resume.UpdatedAt, &resume.DeletedAt)
		if err != nil {
			return nil, err
		}
		resumes = append(resumes, resume)
	}
	return resumes, nil
}
