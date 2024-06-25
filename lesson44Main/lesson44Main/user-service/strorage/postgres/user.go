package postgres

import (
	"database/sql"
	"fmt"
	"time"
	pb "user_service/proto"
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

func (repo *UserRepository) CreateUser(user *pb.UserCreateRequest) (*pb.Void, error) {
	_, err := repo.Db.Exec("insert into users(name,phone,age,created_at) values ($1,$2,$3,$4)", user.Name, user.Phone, user.Age, time.Now())
	response := &pb.Void{}
	return response, err
}
func (repo *UserRepository) UpdateUser(user *pb.UserUpdatedRequest) (*pb.Void, error) {
	_, err := repo.Db.Exec("update  users set name=$1,phone=$2,age=$3,updated_at=$4 where id=$5", user.Name, user.Phone, &user.Age, time.Now(), user.Id)
	response := &pb.Void{}
	return response, err
}
func (repo *UserRepository) DeleteUser(id *pb.GetByIdRequest) (*pb.Void, error) {
	_, err := repo.Db.Exec("update  users set deleted_at=$1 where id=$2", 1, id)
	response := &pb.Void{}
	return response, err
}

func (repo *UserRepository) GetUser(userFilter *pb.UserFilterRequest) ([]*pb.UserResponse, error) {
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
	var users []*pb.UserResponse
	for rows.Next() {
		var user pb.UserResponse
		err := rows.Scan(&user.Id, &user.Name, &user.Phone, &user.Phone, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, err

}

//func (repo *UserRepository) GetById(id string) (*models.User, error) {
//	user := models.User{}
//	rows, err := repo.Db.Query("select id,name,phone,age,created_at,deleted_at,updated_at from users where  id=$1 and   deleted_at=0", id)
//	if err != nil {
//		return nil, err
//	}
//	for rows.Next() {
//		err := rows.Scan(&user.Id, &user.Name, &user.Phone, &user.Age, &user.CreatedAt, &user.DeletedAt, &user.UpdatedAt)
//		if err != nil {
//			return nil, err
//		}
//		return &user, err
//	}
//	return nil, err
//}

func (repo *UserRepository) GetUserCardAmount(id *pb.GetByIdRequest) (*pb.UserAmount, error) {
	rows, err := repo.Db.Query("select u.name,sum(t.amount) from users  u inner join  cards   c on u.id=c.user_id inner join transactions t on  c.id=t.card_id where t.transaction_type='deposit' or t.transaction_type='credit' and u.deleted_at=0 and c.deleted_at=0 and t.deleted_at=0 and u.id=$1 group by u.name ", id)
	if err != nil {
		return nil, err
	}
	var userCard pb.UserAmount
	for rows.Next() {
		err := rows.Scan(&userCard.UserName, &userCard.Amount)
		if err != nil {
			return nil, err
		}
		return &userCard, err
	}
	return nil, err
}
func (repo *UserRepository) GetUserCard(id string) ([]*pb.UserCard, error) {
	rows, err := repo.Db.Query("select u.name,c.number from users u inner join cards c on u.id=c.user_id and u.id=$1 and u.deleted_at=0 and c.deleted_at=0", id)
	if err != nil {
		return nil, err
	}
	var userCards []*pb.UserCard
	for rows.Next() {
		var userCard pb.UserCard
		err := rows.Scan(&userCard.UserName, &userCard.CardNumber)
		if err != nil {
			return nil, err
		}

		userCards = append(userCards, &userCard)
	}
	return userCards, err
}
func (repo *UserRepository) GetUserDepositCard(id *pb.GetByIdRequest) ([]*pb.UserDeposit, error) {
	rows, err := repo.Db.Query("select u.name,c.number,t.transaction_type from  users u inner join  cards c on u.id=c.user_id inner join transactions  t on c.id=t.card_id where u.id=$1 and u.deleted_at=0 and c.deleted_at=0 and t.deleted_at=0 and t.transaction_type='deposit'", id)
	if err != nil {
		return nil, err
	}
	var userCards []*pb.UserDeposit
	for rows.Next() {
		var userCard pb.UserDeposit
		err := rows.Scan(&userCard.UserName, &userCard.CardNumber, &userCard.CardDeposit)
		if err != nil {
			return nil, err
		}

		userCards = append(userCards, &userCard)
	}
	return userCards, err
}
func (repo *UserRepository) GetUserCreditCard(id *pb.GetByIdRequest) ([]*pb.UserCredit, error) {
	rows, err := repo.Db.Query("select u.name,c.number,t.transaction_type from  users u inner join  cards c on u.id=c.user_id inner join transactions  t on c.id=t.card_id where u.id=$1 and u.deleted_at=0 and c.deleted_at=0 and t.deleted_at=0 and t.transaction_type='credit'", id)
	if err != nil {
		return nil, err
	}
	var userCards []*pb.UserCredit
	for rows.Next() {
		var userCard pb.UserCredit
		err := rows.Scan(&userCard.UserName, &userCard.CardNumber, &userCard.CardCredit)
		if err != nil {
			return nil, err
		}

		userCards = append(userCards, &userCard)
	}
	return userCards, err
}
