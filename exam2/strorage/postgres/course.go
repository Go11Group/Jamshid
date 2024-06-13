package postgres

import (
	"database/sql"
	"my_project/model"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (repo UserRepository) CreateUser(user model.User) error {
	_, err := repo.Db.Exec("insert into users(name,email,birthday,password) values ($1,$2,$3,$4)", user.Name, user.Email, user.Birthday, user.Password)
	return err

}
func (repo UserRepository) UpdateUser(id string, user model.User) error {
	_, err := repo.Db.Exec("update users  set name=$1,email=$2,birthday=$3,password=$4 where id =$5", user.Name, user.Email, user.Birthday, user.Password, id)
	return err
}

func (repo UserRepository) DeleteUser(id string) error {
	_, err := repo.Db.Exec("delete from users where id=$1", id)
	return err
}
func (repo UserRepository) GetAllUsers(id string) ([]model.User, error) {
	rows, err := repo.Db.Query("select *from users")
	if err != nil {
		return nil, err
	}
	users := []model.User{}
	for rows.Next() {
		user := model.User{}
		err := rows.Scan(&user.Id,&user.Name, &user.Email, &user.Birthday, &user.Password, &user.CreatedAt, &user.UpdatedAt, user.DeletedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
