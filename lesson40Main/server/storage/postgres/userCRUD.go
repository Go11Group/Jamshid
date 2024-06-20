package postgres

import (
	"database/sql"
	"my_project/models"
	strorage "my_project/storage"
	"time"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (us *UserRepository) CreateUser(user models.User) error {
	_, err := us.DB.Exec("insert into users (name,age,email) values ($1,$2,$3)", user.Name, user.Age, user.Email)
	return err
}
func (us *UserRepository) UpdateUser(id string, user models.User) error {
	_, err := us.DB.Exec("update users set name=$1, age=$2, email=$3,updated_at=$4 where id=$5", user.Name, user.Age, user.Email, time.Now(), id)
	return err
}
func (us *UserRepository) DeleteUser(id string) error {
	_, err := us.DB.Exec("update   users deleted_at=$1 where id=$2", time.Now(), id)
	return err
}
func (us *UserRepository) Get(f models.Filter) ([]models.User, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset string
	)
	query := `select id, name, age, email,created_at, updated_at, deleted_at
	 	from users where true`
	filter := ""

	if len(f.Name) > 0 {
		params["name"] = f.Name
		filter += " and name = :name "
	}

	if f.Age > 0 {
		params["age"] = f.Age
		filter += " and age = :age "
	}
	if f.Limit > 0 {
		params["limit"] = f.Limit
		limit = ` LIMIT :limit`
	}

	if f.Offset > 0 {
		params["offset"] = f.Offset
		offset = ` OFFSET :offset`
	}

	query = query + filter + limit + offset

	query, arr = strorage.ReplaceQueryParams(query, params)
	rows, err := us.DB.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	users := []models.User{}
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Age, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil

}
