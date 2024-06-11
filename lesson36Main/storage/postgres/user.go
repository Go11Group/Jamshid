package postgres

import (
	"database/sql"
	"my_project/model"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (us *UserRepository) CreateUser(user model.User) error {
	_, err := us.DB.Exec("insert into  users(name ,email,phone_number,birthday,gender) values ($1,$2,$3,$4,$5)", &user.Name, &user.Email, &user.PhoneNumber, &user.Birthday, &user.Gender)
	return err
}
func (us *UserRepository) UpdateUser(id string, user model.User) error {
	_, err := us.DB.Exec("update  users set name=$1 ,email=$2,phone_number=$3,birthday=$4,gender=$5 where id=$5 ", user.Name, user.Email, user.PhoneNumber, user.Birthday, user.Gender, id)
	return err
}
func (us *UserRepository) DeletedUser(id string) error {
	_, err := us.DB.Exec("delete from users where id=$1", id)
	return err
}
func (us *UserRepository) GetAllUser() ([]model.User, error) {
	rows, err := us.DB.Query("select *from recruiters")
	if err != nil {
		return nil, err
	}
	users := []model.User{}
	for rows.Next() {
		user := model.User{}
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.PhoneNumber, &user.Birthday, &user.Gender, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
