package postgres

import (
	"database/sql"
	"fmt"
	"my_project/model"
	"my_project/strorage"
	"time"
)

type CourseRepository struct {
	Db *sql.DB
}

func NewCourseRepository(db *sql.DB) *CourseRepository {
	return &CourseRepository{Db: db}
}

func (repo CourseRepository) CreateCourse(course model.Course) error {
	_, err := repo.Db.Exec("insert into courses(title,description) values ($1,$2)", course.Title, course.Description)
	return err

}
func (repo CourseRepository) UpdateCourse(id string, course model.Course) error {
	_, err := repo.Db.Exec("update courses  set title=$1,description=$2 where id =$3", course.Title, course.Description, id)
	return err
}

func (repo CourseRepository) DeleteCourse(id string) error {
	fmt.Println("+++++++++++++++++", id)
	_, err := repo.Db.Exec("delete from courses where id=$1", id)
	now := time.Now()
	unixTime := now.Unix()
	DeletedAt := int(unixTime)
	_, err = repo.Db.Exec("update  courses set deleted_at=$1 where id=$2", DeletedAt, id)
	return err
}

func (repo *CourseRepository) GetCourses(f model.CourseFilter) ([]model.Course, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset string
	)
	query := `select id, title, description, created_at, updated_at, deleted_at
	 	from courses where true`
	filter := ""

	if len(f.Title) > 0 {
		params["title"] = f.Title
		filter += " and title = :title "
	}

	if len(f.Description) > 0 {
		params["description"] = f.Description
		filter += " and description = :description "
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
	courses := []model.Course{}
	for rows.Next() {
		course := model.Course{}
		err := rows.Scan(&course.Id, &course.Title, &course.Description, &course.CreatedAt, &course.UpdatedAt, &course.DeletedAt)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil

}

func (repo CourseRepository) GetLessonByCourseId(courseId string) (string, []model.Lesson, error) { // 2-task bajarildi
	rows, err := repo.Db.Query("SELECT course_id, title, content FROM lessons WHERE course_id = $1", courseId)
	if err != nil {
		return "", nil, err
	}
	defer rows.Close()

	var lessons []model.Lesson
	for rows.Next() {
		lesson := model.Lesson{}
		err := rows.Scan(&lesson.CourseId, &lesson.Title, &lesson.Content)
		if err != nil {
			return "", nil, err
		}
		lessons = append(lessons, lesson)
	}

	return courseId, lessons, nil
}
func (repo CourseRepository) GetUserByCourseIdWithEnrollment(course_id string) (string, []model.User, error) { //3-task bajarildi
	var user_id string
	err := repo.Db.QueryRow("select  user_id from enrollments where course_id = $1", course_id).Scan(&user_id)
	if err != nil {
		return "", nil, err
	}
	rows, err := repo.Db.Query("select id,name,email from  users where id=$1", &user_id)
	if err != nil {
		return "", nil, err
	}
	users := []model.User{}
	for rows.Next() {
		user := model.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			return "", nil, err
		}
		users = append(users, user)
	}
	return course_id, users, nil
}

func (repo CourseRepository) GetById(id string) (*model.Course, error) {
	course := model.Course{}
	err := repo.Db.QueryRow("select id,title,description from courses where id=$1", id).Scan(&course.Id, &course.Title, &course.Description)
	if err != nil {
		fmt.Println("+++++++", err)
		return nil, err
	}
	return &course, nil
}

func (repo CourseRepository) GetPopularyCourse(startTime, endTime time.Time) (int, *model.Course, error) {
	course := model.Course{}
	var enrollmentsCount int
	err := repo.Db.QueryRow(`select c.id, c.title, COUNT(*)from courses c inner join  enrollments e ON c.id = e.course_id where e.enrollment_date BETWEEN $1 AND $2 group by  c.id, c.title order by  count(*) desc limit 1`, startTime, endTime).Scan(&course.Id, &course.Title, &enrollmentsCount)
	if err != nil {
		fmt.Println("+++++++", err)
		return 0, nil, err
	}
	return enrollmentsCount, &course, nil
}
