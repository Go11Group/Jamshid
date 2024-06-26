package postgres

import (
	"database/sql"
	pb "library_service/genproto"
	strorage "library_service/storage"
	"time"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (u *UserRepository) Create(user *pb.CreateUserRequest) (*pb.Void, error) {
	_, err := u.Db.Exec("insert into user (fist_name,last_name, created_at) values ($1,$2,$3)",
		user.Name, user.LastName, time.Now())
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (u *UserRepository) GetById(id *pb.ByIdRequest) (*pb.UserResponse, error) {
	var user pb.UserResponse

	err := u.Db.QueryRow("select name,surname from user where id = $1", id).Scan(&user.Name, &user.Lastname)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (u *UserRepository) Update(user *pb.UpdatedUserRequest, id string) error {
	_, err := u.Db.Exec("update user set name = $1,surname = $2,updated_at = $3 where id = $4 and deleted_at = 0", user.Name, user.LasName, time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) Delete(id *pb.ByIdRequest) error {
	_, err := u.Db.Exec("update user set deleted_at = $1 where id = $2", id)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) GetAll(f *pb.UserFilterRequest) (*pb.UsersResponse, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset string
	)
	query := `select id, name,surname,created_at,updated_at,deleted_at from user where deleted_at = 0 `
	filter := ` where true `

	if len(f.Name) > 0 {
		params["name"] = f.Name
		filter += "and name = :name "
	}

	if len(f.LastName) > 0 {
		params["surname"] = f.LastName
		filter += "and surname = :surname "
	}

	if (f.LimitOffset.Limit) > 0 {
		params["limit"] = f.LimitOffset.Limit
		filter += "and limit = :limit "
	}

	if (f.LimitOffset.Offset) > 0 {
		params["offset"] = f.LimitOffset.Offset
		filter += "and offset = :offset "
	}

	query = query + filter + limit + offset
	query, arr = strorage.ReplaceQueryParams(query, params)
	rows, err := u.Db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	var users []*pb.UserResponse
	for rows.Next() {
		var user pb.UserResponse
		err := rows.Scan(&user.Id, &user.Name, &user.Lastname, user.CreatedAt, user.UpdatedAt, user.DeletedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return &pb.UsersResponse{UsersResponse: users}, nil
}
