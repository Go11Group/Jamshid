package postgres

import (
	"database/sql"
	"my_project/model"
	"my_project/strorage"
	"time"
)

type EnrollmentRepository struct {
	Db *sql.DB
}

func NewEnrollmentRepository(db *sql.DB) *EnrollmentRepository {
	return &EnrollmentRepository{Db: db}
}

func (repo EnrollmentRepository) CreateEnrollment(enrollment model.Enrollment) error {
	_, err := repo.Db.Exec("insert into enrollments(user_id,course_id,enrollment_date) values ($1,$2,$3)", enrollment.UserId, enrollment.CourseId, enrollment.EnrollmentDate)
	return err

}
func (repo EnrollmentRepository) UpdateEnrollment(id string, enrollment model.Enrollment) error {
	_, err := repo.Db.Exec("update enrollments  set user_id=$1,course_id=$2,enrollment_date=$3 where id =$4", enrollment.UserId, enrollment.CourseId, enrollment.EnrollmentDate, id)
	return err
}

func (repo EnrollmentRepository) DeleteEnrollment(id string) error {
	_, err := repo.Db.Exec("delete from enrollments where id=$1", id)
	now := time.Now()
	unixTime := now.Unix()
	DeletedAt := int(unixTime)
	_, err = repo.Db.Exec("update  enrollments set deleted_at=$1 where id=$2", DeletedAt, id)
	return err
}

func (repo *EnrollmentRepository) GetEnrollments(f model.EnrollmentFilter) ([]model.Enrollment, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset string
	)
	query := `select id, user_id, course_id,enrollment_date, created_at, updated_at, deleted_at
	 	from courses where true`
	filter := ""

	if len(f.User_id) > 0 {
		params["user_id"] = f.User_id
		filter += " and user_id = :user_id "
	}

	if len(f.Course_id) > 0 {
		params["course_id"] = f.Course_id
		filter += " and course_id = :course_id "
	}
	if len(f.EnrollmentDate) > 0 {
		params["enrollment_date"] = f.EnrollmentDate
		filter += " and enrollment_date = :enrollment_date "
	}
	if f.Limit > 0 {
		params["limit"] = f.Limit
		limit = ` LIMIT :limit`
	}

	if f.Offset > 0 {
		params["offset"] = (f.Offset - 1) * f.Limit
		offset = ` OFFSET :offset`
	}

	query = query + filter + limit + offset

	query, arr = strorage.ReplaceQueryParams(query, params)

	rows, err := repo.Db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	enrollments := []model.Enrollment{}
	for rows.Next() {
		enrollment := model.Enrollment{}
		err := rows.Scan(&enrollment.Id, &enrollment.UserId, &enrollment.CourseId, &enrollment.EnrollmentDate, &enrollment.CreatedAt, &enrollment.UpdatedAt, &enrollment.DeletedAt)
		if err != nil {
			return nil, err
		}
		enrollments = append(enrollments, enrollment)
	}
	return enrollments, nil

}
func (repo EnrollmentRepository) GetById(id string) (*model.Enrollment, error) {
	enrollment := model.Enrollment{}
	err := repo.Db.QueryRow("select id,user_id,course_id,enrollment_date from enrollments where id=$1", id).Scan(&enrollment.Id, &enrollment.UserId, &enrollment.CourseId, &enrollment.EnrollmentDate)
	if err != nil {
		return nil, err
	}
	return &enrollment, nil
}
