package postgres

import (
	"database/sql"
	"my_project/model"
)

type CourseRepository struct {
	DB *sql.DB
}

func (db *CourseRepository) InsertIntoCourse(course model.Course, courseStudent model.CourseStudent) error {

	_, err := db.DB.Exec("insert into  course_student(id,course_id,student_id) values ($1,$2,$3)", courseStudent.Id, courseStudent.CourseId, courseStudent.StudentId)
	if err != nil {
		return err
	}
	_, err = db.DB.Exec("insert into course (id,name,) values ($1,$2)", course.Id, course.CourseName)
	if err != nil {
		return err
	}
	return nil

}

func (db *CourseRepository) UpdateCourse(id string, course model.Course) error {
	_, err := db.DB.Exec("update course_student set course_id=$1, where id=$2", course.Id, id)
	_, err = db.DB.Exec("update course set id=$1 ,course_name=$2 where  id=$3 and course_id=%4", course.Id, course.CourseName, id)
	if err != nil {
		return err
	}
	return nil

}

func (db *CourseRepository) DeleteCourse(id string) error {
	_, err := db.DB.Exec("delete from  course_student where course_id = $1", id)
	if err != nil {
		return err
	}

	_, err = db.DB.Exec("delete from course where id=$1", id)
	if err != nil {
		return err
	}
	return nil

}
func (db *CourseRepository) ReadCourse() ([]model.Course, error) {
	courses := []model.Course{}
	rows, err := db.DB.Query("select * from course")

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		course := model.Course{}
		rows.Scan(&course.Id, &course.CourseName)
		courses = append(courses, course)
	}
	return courses, nil
}
