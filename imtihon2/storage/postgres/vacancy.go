package postgres

import (
	"database/sql"
	"my_project/model"
)

type VacancyRepository struct {
	DB *sql.DB
}

func NewVacancyRepository(db *sql.DB) *VacancyRepository {
	return &VacancyRepository{
		DB: db,
	}
}

func (vc *VacancyRepository) CreateVacancy(vacancy model.Vacancy) error {
	_, err := vc.DB.Exec("insert into  vacancies(name ,min_experience,company_id,description) values ($1,$2,$3,$4)", &vacancy.Name, &vacancy.MinExperience, &vacancy.CompanyID, &vacancy.Description)
	return err
}
func (vc *VacancyRepository) UpdateVacancy(id string, vacancy model.Vacancy) error {
	_, err := vc.DB.Exec("update  vacancies set name=$1 ,min_experience=$2,company_id=$3,description where id=$5 ", &vacancy.Name, &vacancy.MinExperience, &vacancy.CompanyID, &vacancy.Description, &id)
	return err
}
func (vc *VacancyRepository) DeletedVacancy(id string) error {
	_, err := vc.DB.Exec("delete from vacancies where id=$1", &id)
	return err
}
func (vc *VacancyRepository) GetAllVacancy() ([]model.Vacancy, error) {
	rows, err := vc.DB.Query("select *from vacancies")
	if err != nil {
		return nil, err
	}
	vacancies := []model.Vacancy{}
	for rows.Next() {
		vacancy := model.Vacancy{}
		err = rows.Scan(&vacancy.Id, &vacancy.Name, &vacancy.MinExperience, &vacancy.CompanyID, &vacancy.Description, &vacancy.CreatedAt, &vacancy.UpdatedAt, &vacancy.DeletedAt)
		if err != nil {
			return nil, err
		}
		vacancies = append(vacancies, vacancy)
	}
	return vacancies, nil
}
