package postgres

import (
	"database/sql"
	"fmt"
	"my_project/model"
	"my_project/strorage"
	"time"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (repo UserRepository) CreateUser(user model.User) error {
	parsedTime, err := time.Parse("2006-01-02", user.Birthday)
	if err != nil {
		fmt.Println("Vaqtni o'zgartirishda xatolik:", err)
		return nil
	}
	_, err = repo.Db.Exec("insert into users(name,email,birthday,password) values ($1,$2,$3,$4)", user.Name, user.Email, parsedTime, user.Password)
	return err

}
func (repo UserRepository) UpdateUser(id string, user model.User) error {
	_, err := repo.Db.Exec("update users  set name=$1,email=$2,birthday=$3,password=$4 where id =$5", user.Name, user.Email, user.Birthday, user.Password, id)
	return err
}

func (repo UserRepository) DeleteUser(id string) error {
	_, err := repo.Db.Exec("delete from users where id=$1", id)
	now := time.Now()
	unixTime := now.Unix()
	DeletedAt := int(unixTime)
	_, err = repo.Db.Exec("update  users set deleted_at=$1 where id=$2", DeletedAt, id)
	return err
}

func (repo *UserRepository) GetUser(f model.UserFilter) ([]model.User, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset string
	)
	query := `select id, name, email, birthday,password, created_at, updated_at, deleted_at
	 	from users where true`
	filter := ""
	if len(f.Name) > 0 {
		params["name"] = f.Name
		filter += " and name = :name "
	}

	if len(f.Email) > 0 {
		params["email"] = f.Email
		filter += " and email = :email "
	}

	if len(f.Password) > 0 {
		params["password"] = f.Password
		filter += " and password = :password "
	}
	if len(f.Birthday) > 0 {
		params["birthday"] = f.Birthday
		filter += " and birthday = :birthday "
	}
	if f.Limit > 0 {
		params["limit"] = f.Limit
		limit = ` LIMIT :limit`
	}

	if f.Offset > 0 {
		params["offset"] = (f.Offset - 1) * f.Limit
	}

	query = query + filter + limit + offset

	query, arr = strorage.ReplaceQueryParams(query, params)

	rows, err := repo.Db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	users := []model.User{}
	for rows.Next() {
		user := model.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Birthday, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil

}

func (repo UserRepository) GetEnrollmentByCourseId(user_id string) (string, []model.Course, error) { //1-task bajarildi

	var course_id string
	err := repo.Db.QueryRow("select  course_id from enrollments where user_id = $1", &user_id).Scan(&course_id)
	if err != nil {
		fmt.Println("----------", course_id)
		fmt.Println("))))))))/////", err)
		return "", nil, err
	}
	rows, err := repo.Db.Query("select id,title,description from  courses where id=$1", &course_id)
	if err != nil {
		fmt.Println("--------", err)
		return "", nil, err
	}
	courses := []model.Course{}
	for rows.Next() {
		course := model.Course{}
		err := rows.Scan(&course.Id, &course.Title, &course.Description)
		if err != nil {
			fmt.Println("00000000", err)
			return "", nil, err
		}
		courses = append(courses, course)
	}
	return user_id, courses, nil
}

func (repo *UserRepository) GetUserByEmailOrName(name, email string) ([]model.User, error) {

	rows, err := repo.Db.Query("select id,name,email from users where name=$1 or email=$2", &name, &email)
	if err != nil {
		return nil, err
	}
	users := []model.User{}
	for rows.Next() {
		user := model.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (repo UserRepository) GetById(id string) (*model.User, error) {
	user := model.User{}
	err := repo.Db.QueryRow("select id, name,email,password from users where id=$1", id).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		fmt.Println("_____________", err)
		return nil, err
	}
	return &user, nil
}
