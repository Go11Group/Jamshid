package postgres

import (
	"database/sql"
	"fmt"
	"my_project/model"
	"strconv"
	"strings"
)

var (
	params = make(map[string]interface{})
	arr    []interface{}
	limit  string
)

type StudentRepository struct {
	Db *sql.DB
}

func NewStudentRepository(db *sql.DB) *StudentRepository {
	return &StudentRepository{
		Db: db,
	}
}

func (stu *StudentRepository) GetAll(f model.Filter) ([]model.Student, error) {

	query := `select id, first_name, last_name, age, gender, nation, field, parent_name, city
	 	from students where true`
	filter := ""
	if len(f.Gender) > 0 {
		params["gender"] = f.Gender
		filter += " and gender = :gender "
	}

	if len(f.Nation) > 0 {
		params["nation"] = f.Gender
		filter += " and nation = :nation "
	}

	if len(f.Field) > 0 {
		params["field"] = f.Gender
		filter += " and field = :field "
	}

	if f.Age > 0 {
		params["age"] = f.Gender
		filter += " and age = :age "
	}

	query = query + filter + limit // + offset

	query, arr = ReplaceQueryParams(query, params)

	rows, err := stu.Db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	students := []model.Student{}
	for rows.Next() {
		student := model.Student{}
		err := rows.Scan(&student.FirstName)
		if err != nil {
			fmt.Printf("+++++++++++++", err)
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil

}

func ReplaceQueryParams(namedQuery string, params map[string]interface{}) (string, []interface{}) {
	var (
		i    int = 1
		args []interface{}
	)

	for k, v := range params {
		if k != "" && strings.Contains(namedQuery, ":"+k) {
			namedQuery = strings.ReplaceAll(namedQuery, ":"+k, "$"+strconv.Itoa(i))
			args = append(args, v)
			i++
		}
	}

	return namedQuery, args
}
