package postgres

import (
	"database/sql"
	"my_project/model"
)

type ProblemRepository struct {
	Db *sql.DB
}

func NewProblemRepository(db *sql.DB) *ProblemRepository {
	return &ProblemRepository{
		Db: db,
	}
}

func (slp *ProblemRepository) CreateProblem(problem model.Problem) error {
	_, err := slp.Db.Exec("insert into  problems (description,type) values ($1,$2)", problem.Description, problem.Type)
	return err
}
func (slp *ProblemRepository) DeleteProblem(id string) error {
	_, err := slp.Db.Exec("delete from problems where id=$1", id)
	return err
}
func (slp *ProblemRepository) UpdatedProblem(id string, problem model.Problem) error {
	_, err := slp.Db.Exec("update  problems set description=$1,type=$2 where id=$3", problem.Description, problem.Type, id)
	return err
}
func (slp *ProblemRepository) GetAllProblem() ([]model.Problem, error) {
	rows, err := slp.Db.Query("select id,description,type,created_at,updated_at,deleted_at from  problems")
	if err != nil {
		return nil, err
	}
	problems := []model.Problem{}
	for rows.Next() {
		problem := model.Problem{}
		err := rows.Scan(&problem.Id, &problem.Description, &problem.Type, &problem.CreatedAt, &problem.UpdatedAt, &problem.DeletedAt)
		if err != nil {
			return nil, err
		}
		problems = append(problems, problem)
	}
	return problems, nil
}
