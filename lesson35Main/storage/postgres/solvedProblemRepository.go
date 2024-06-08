package postgres

import (
	"database/sql"
	"my_project/model"
)

type SolvedProblemRepository struct {
	Db *sql.DB
}

func NewSolvedProblemRepository(db *sql.DB) *SolvedProblemRepository {
	return &SolvedProblemRepository{
		Db: db,
	}
}

func (slp *SolvedProblemRepository) CreateSolvedProblem(solved_solution model.SolvedProblem) error {
	_, err := slp.Db.Exec("insert into  solved_problems (solution,problem_solution,user_solution) values ($1,$2,$3)", solved_solution.Solution, solved_solution.ProblemSolution, solved_solution.UserSolution)
	return err
}
func (slp *SolvedProblemRepository) DeleteSolvedProblem(id string) error {
	_, err := slp.Db.Exec("delete from solved_problems where id=$1", id)
	return err
}
func (slp *SolvedProblemRepository) UpdatedSolvedProblem(id string, solved_solution model.SolvedProblem) error {
	_, err := slp.Db.Exec("update  solved_problems set solution=$1,problem_solution=$2,user_solution=$3 where id=$4", solved_solution.Solution, solved_solution.ProblemSolution, solved_solution.UserSolution, id)
	return err
}
func (slp *SolvedProblemRepository) GetAllSolvedProblem() ([]model.SolvedProblem, error) {
	rows, err := slp.Db.Query("select id,solution,problem_solution,user_solution,created_at,updated_at,deleted_at from  solved_problems")
	if err != nil {
		return nil, err
	}
	solved_problems := []model.SolvedProblem{}
	for rows.Next() {
		solv_pro := model.SolvedProblem{}
		err := rows.Scan(&solv_pro.Id, &solv_pro.Solution, &solv_pro.ProblemSolution, &solv_pro.UserSolution, &solv_pro.CreatedAt, &solv_pro.UpdatedAt, &solv_pro.DeletedAt)
		if err != nil {
			return nil, err
		}
		solved_problems = append(solved_problems, solv_pro)
	}
	return solved_problems, nil
}
