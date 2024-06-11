package postgres

import (
	"database/sql"
	"fmt"
	"my_project/model"
)

type CompanyRepository struct {
	DB *sql.DB
}

func NewCompanyRepository(db *sql.DB) *CompanyRepository {
	return &CompanyRepository{
		DB: db,
	}
}

func (comp *CompanyRepository) CreateCompany(company model.Company) error {
	_, err := comp.DB.Exec("insert into  companies(name ,location,workers) values ($1,$2,$3)", &company.Name, company.Location, company.Workers)
	return err
}
func (comp *CompanyRepository) UpdateCompany(id string, company model.Company) error {
	_, err := comp.DB.Exec("update  companies set name=$1 ,location=$2,workers=$3 where id=$4 ", &company.Name, &company.Location, &company.Workers, &id)
	return err
}
func (comp *CompanyRepository) DeletedCompany(id string) error {
	_, err := comp.DB.Exec("delete from companies where id=$1", id)
	return err
}
func (comp *CompanyRepository) GetAllCompany() ([]model.Company, error) {
	rows, err := comp.DB.Query("select * from companies")
	if err != nil {
		fmt.Println("-------------", err)
		return nil, err
	}
	companies := []model.Company{}
	for rows.Next() {
		company := model.Company{}
		err = rows.Scan(&company.Id, &company.Name, &company.Location, &company.Workers, &company.CreatedAt, &company.UpdatedAt, &company.DeletedAt)
		if err != nil {
			return nil, err
		}
		companies = append(companies, company)
	}
	return companies, nil
}

//func  (comp *CompanyRepository) GetById(id string)(*model.Company,error) {
//	err:=comp.DB.QueryRow("select name,loaction,workers,created_at,updated_at,deleted_at from companies where id=$1",&comp)
//	if err!=nil{
//		return nil,err
//	}
//}
