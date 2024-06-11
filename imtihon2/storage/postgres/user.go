package postgres

import (
	"database/sql"
	"fmt"
	"my_project/model"
)

type InterviewRepository struct {
	DB *sql.DB
}

func NewInterviewRepository(db *sql.DB) *CompanyRepository {
	return &CompanyRepository{
		DB: db,
	}
}

func (comp *CompanyRepository) CreateInterview(company model.Company) error {
	_, err := comp.DB.Exec("insert into  companies(name ,location,workers) values ($1,$2,$3)", company.Name, company.Location, company.Workers)
	return err
}
func (comp *CompanyRepository) UpdateInterview(id string, company model.Company) error {
	_, err := comp.DB.Exec("update  companies set name=$1 ,location=$2,workers=$3 where id=$5 ", company.Name, company.Location, company.Workers, id)
	return err
}
func (comp *CompanyRepository) DeletedInterview(id string) error {
	_, err := comp.DB.Exec("delete from companies where id=$1", id)
	return err
}
func (comp *CompanyRepository) GetAllInterview(column interface{}, values string) error {
	query := fmt.Sprintf("select *from %s=$1", column, values)
	_, err := comp.DB.Exec(query)
	return err
}
