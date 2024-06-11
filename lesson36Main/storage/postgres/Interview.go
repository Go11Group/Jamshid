package postgres

import (
	"database/sql"
	"my_project/model"
)

type InterviewRepository struct {
	DB *sql.DB
}

func NewInterviewRepository(db *sql.DB) *InterviewRepository {
	return &InterviewRepository{
		DB: db,
	}
}

func (trv *InterviewRepository) CreateInterview(interview model.Interview) error {
	_, err := trv.DB.Exec("insert into  interviews(user_id ,vacancy_id,recruiter_id,interview_date) values ($1,$2,$3,$4)", &interview.UserId, &interview.VacancyId, &interview.RecruiterId, &interview.InterviewDate)
	return err
}
func (trv *InterviewRepository) UpdateInterview(id string, interview model.Interview) error {
	_, err := trv.DB.Exec("update  interviews set user_id=$1 ,vacancy_id=$2,recruiter_id=$3, interview_date=$4 where id=$5 ", &interview.UserId, &interview.VacancyId, &interview.RecruiterId, &interview.InterviewDate, &id)
	return err
}
func (trv *InterviewRepository) DeletedInterview(id string) error {
	_, err := trv.DB.Exec("delete from interviews where id=$1", &id)
	return err
}
func (trv *InterviewRepository) GetAllInterview() ([]model.Interview, error) {
	rows, err := trv.DB.Query("select *from interviews")
	if err != nil {
		return nil, err
	}
	interviews := []model.Interview{}
	for rows.Next() {
		interview := model.Interview{}
		err = rows.Scan(&interview.Id, &interview.UserId, &interview.VacancyId, &interview.RecruiterId, &interview.InterviewDate, &interview.CreatedAt, &interview.UpdatedAt, &interview.DeletedAt)
		if err != nil {
			return nil, err
		}
		interviews = append(interviews, interview)
	}
	return interviews, nil
}
