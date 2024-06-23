package postgres

import (
	"database/sql"
	"fmt"
	"time"
	"user_service/models"
	"user_service/strorage"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		Db: db,
	}
}

func (repo *UserRepository) CreateUser(user *models.User) error {
	_, err := repo.Db.Exec("insert into users(name,phone,age,created_at) values ($1,$2,$3,$4)", user.Name, user.Phone, user.Age, time.Now())
	return err
}
func (repo *UserRepository) UpdateUser(id string, user *models.User) error {
	_, err := repo.Db.Exec("update  users set name=$1,phone=$2,age=$3,updated_at=$4 where id=$5", user.Name, user.Phone, &user.Age, time.Now(), id)
	return err
}
func (repo *UserRepository) DeleteUser(id string) error {
	_, err := repo.Db.Exec("update  users set deleted_at=$1 where id=$2", 1, id)
	return err
}

func (repo *UserRepository) GetUser(userFilter models.Filter) (*[]models.User, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset string
	)
	query := "select id,name,phone,age,created_at,deleted_at,updated_at from users where  deleted_at=0 "
	filter := ""
	if len(userFilter.Name) > 0 {
		params["name"] = userFilter.Name
		filter += " and name = :name "

	}
	if userFilter.Age > 0 {
		params["age"] = userFilter.Age
		filter += "and age=:age"
	}
	if len(userFilter.Phone) > 0 {
		params["phone"] = userFilter.Phone
		filter += "and phone =:phone"
	}
	if userFilter.Limit > 0 {
		params["limit"] = userFilter.Limit
		limit = ` LIMIT :limit`

	}
	if userFilter.Offset > 0 {
		params["offset"] = userFilter.Limit
		limit = ` OFFSET :offset`

	}
	query = query + filter + limit + offset
	query, arr = strorage.ReplaceQueryParams(query, params)
	fmt.Println("----------------", query, arr)
	rows, err := repo.Db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Name, &user.Phone, &user.Phone, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return &users, err

}
func (repo *UserRepository) GetById(id string) (*models.User, error) {
	user := models.User{}
	rows, err := repo.Db.Query("select id,name,phone,age,created_at,deleted_at,updated_at from users where  id=$1 and   deleted_at=0", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Phone, &user.Age, &user.CreatedAt, &user.DeletedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		return &user, err
	}
	return nil, err
}

func (repo *UserRepository) GetUserCardAmount(id string) (*models.UserCard, error) {
	rows, err := repo.Db.Query("select u.name,sum(t.amount) from users  u inner join  cards   c on u.id=c.user_id inner join transactions t on  c.id=t.card_id where t.transaction_type='deposit' or t.transaction_type='credit' and u.deleted_at=0 and c.deleted_at=0 and t.deleted_at=0 and u.id=$1 group by u.name ", id)
	if err != nil {
		return nil, err
	}
	var userCard models.UserCard
	for rows.Next() {
		err := rows.Scan(&userCard.UserName, &userCard.CardAmount)
		if err != nil {
			return nil, err
		}
		return &userCard, err
	}
	return nil, err
}
func (repo *UserRepository) GetUserCard(id string) (*[]models.UserCard, error) {
	rows, err := repo.Db.Query("select u.name,c.number from users u inner join cards c on u.id=c.user_id and u.id=$1 and u.deleted_at=0 and c.deleted_at=0", id)
	if err != nil {
		return nil, err
	}
	var userCards []models.UserCard
	for rows.Next() {
		var userCard models.UserCard
		err := rows.Scan(&userCard.UserName, &userCard.CardNumber)
		if err != nil {
			return nil, err
		}

		userCards = append(userCards, userCard)
	}
	return &userCards, err
}
func (repo *UserRepository) GetUserDepositCard(id string) (*[]models.UserCard, error) {
	rows, err := repo.Db.Query("select u.name,c.number,t.transaction_type from  users u inner join  cards c on u.id=c.user_id inner join transactions  t on c.id=t.card_id where u.id=$1 and u.deleted_at=0 and c.deleted_at=0 and t.deleted_at=0 and t.transaction_type='deposit'", id)
	if err != nil {
		return nil, err
	}
	var userCards []models.UserCard
	for rows.Next() {
		var userCard models.UserCard
		err := rows.Scan(&userCard.UserName, &userCard.CardNumber, &userCard.CardDeposit)
		if err != nil {
			return nil, err
		}

		userCards = append(userCards, userCard)
	}
	return &userCards, err
}
func (repo *UserRepository) GetUserCreditCard(id string) (*[]models.UserCard, error) {
	rows, err := repo.Db.Query("select u.name,c.number,t.transaction_type from  users u inner join  cards c on u.id=c.user_id inner join transactions  t on c.id=t.card_id where u.id=$1 and u.deleted_at=0 and c.deleted_at=0 and t.deleted_at=0 and t.transaction_type='credit'", id)
	if err != nil {
		return nil, err
	}
	var userCards []models.UserCard
	for rows.Next() {
		var userCard models.UserCard
		err := rows.Scan(&userCard.UserName, &userCard.CardNumber, &userCard.CardCredit)
		if err != nil {
			return nil, err
		}

		userCards = append(userCards, userCard)
	}
	return &userCards, err
}
