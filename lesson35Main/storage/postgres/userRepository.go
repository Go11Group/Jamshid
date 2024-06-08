package postgres

import (
	"database/sql"
	"fmt"
	"my_project/model"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		Db: db,
	}
}

func (sr *UserRepository) CreateUser(user model.User) error {

	_, err := sr.Db.Exec("insert into  users(first_name,last_name,age,email,phone) values ($1,$2,$3,$4,$5) ", user.FirstName, user.LastName, user.Age, user.Email, user.Phone)
	if err != nil {
		fmt.Println("error --------------")
	}
	return err
}
func (sr *UserRepository) DeleteUser(id string) error {
	_, err := sr.Db.Exec("delete from users where id=$1", id)
	if err != nil {
		fmt.Println("error is database")
	}
	return err
}
func (sr *UserRepository) UpdatedUser(id string, user model.User) error {
	_, err := sr.Db.Exec("update  users set first_name=$1,last_name=$2,age=$3,email=$4,phone=$5 where id=$6", user.FirstName, user.LastName, user.Age, user.Email, user.Phone, id)
	return err
}
func (sr *UserRepository) GetAllUser() ([]model.User, error) {
	rows, err := sr.Db.Query("select first_name,last_name,age,email,phone,created_at,updated_at,deleted_at from users")
	if err != nil {
		fmt.Println("Error qye========")
		return nil, err
	}
	users := []model.User{}
	for rows.Next() {
		user := model.User{}
		err := rows.Scan(&user.FirstName, &user.LastName, &user.Age, &user.Email, &user.Phone, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			fmt.Println("row error1-390847")
			return nil, err
		}
		users = append(users, user)
	}
	fmt.Println("user-------", users)
	return users, nil
}
