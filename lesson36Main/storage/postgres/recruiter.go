package postgres

import (
	"database/sql"
	"my_project/model"
)

type RecruiterRepository struct {
	DB *sql.DB
}

func NewRecruiterRepository(db *sql.DB) *RecruiterRepository {
	return &RecruiterRepository{
		DB: db,
	}
}

func (rc *RecruiterRepository) CreateRecruiter(recruiter model.Recruiter) error {
	_, err := rc.DB.Exec("insert into  recruiters(name ,email,phone_number,birthday,gender,company_id) values ($1,$2,$3,$4,$5,$6)", &recruiter.Name, &recruiter.Email, &recruiter.PhoneNumber, &recruiter.Birthday, &recruiter.Gender, &recruiter.CompanyId)
	return err
}
func (rc *RecruiterRepository) UpdateRecruiter(id string, recruiter model.Recruiter) error {
	_, err := rc.DB.Exec("update  recruiters set name=$1 ,email=$2,phone_number=$3,birthday=$4,gender=$5,company_id=$6 where id=$7 ", &recruiter.Name, &recruiter.Email, &recruiter.PhoneNumber, &recruiter.Birthday, &recruiter.Gender, &recruiter.CompanyId, &id)
	return err
}
func (rc *RecruiterRepository) DeletedRecruiter(id string) error {
	_, err := rc.DB.Exec("delete from recruiters where id=$1", id)
	return err
}
func (rc *RecruiterRepository) GetAllRecruiter() ([]model.Recruiter, error) {
	rows, err := rc.DB.Query("select *from recruiters")
	if err != nil {
		return nil, err
	}
	recruiters := []model.Recruiter{}
	for rows.Next() {
		recruiter := model.Recruiter{}
		err = rows.Scan(&recruiter.Id, &recruiter.Name, &recruiter.Email, &recruiter.PhoneNumber, &recruiter.Birthday, &recruiter.Gender, recruiter.CompanyId, &recruiter.CreatedAt, &recruiter.UpdatedAt, &recruiter.DeletedAt)
		if err != nil {
			return nil, err
		}
		recruiters = append(recruiters, recruiter)
	}
	return recruiters, nil
}
