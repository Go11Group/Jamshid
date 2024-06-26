package postgres

import (
	"database/sql"
	pb "library_service/genproto"
	strorage "library_service/storage"
	"time"
)

type BorrowRepo struct {
	Db *sql.DB
}

func (b *BorrowRepo) CreateBorrow(user *pb.CreateBorrowRequest) (*pb.Void, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BorrowRepo) UpdateBorrow(user *pb.UpdatedBorrowRequest) (*pb.Void, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BorrowRepo) DeleteBorrow(user *pb.ByIdRequest) (*pb.Void, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BorrowRepo) GeBorrow(user *pb.BorrowFilterRequest) (*pb.BorrowsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewBorrowRepo(db *sql.DB) *BorrowRepo {
	return &BorrowRepo{Db: db}
}

func (b *BorrowRepo) Create(borrow *pb.CreateBorrowRequest) error {
	_, err := b.Db.Exec("insert into borrow (user_id,book_id, created_at) values ($1,$2,$3)",
		borrow.Userid, borrow.BookId, time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (b *BorrowRepo) GetById(id *pb.ByIdRequest) (*pb.BorrowResponse, error) {
	var borrow pb.BorrowResponse

	err := b.Db.QueryRow("select user_id,book_id from borrow where id = $1", id).Scan(&borrow.Userid, &borrow.BookId)
	if err != nil {
		return nil, err
	}
	return &borrow, err
}

func (b *BorrowRepo) Update(borrow *pb.UpdatedBorrowRequest) error {
	_, err := b.Db.Exec("update borrows set user_id = $1,book_id = $2,updated_at = $3 where id = $4 and deleted_at = 0", borrow.Userid, borrow.BookId, time.Now(), borrow.Id)
	if err != nil {
		return err
	}
	return nil
}

func (b *BorrowRepo) Delete(id *pb.ByIdRequest) error {
	_, err := b.Db.Exec("update borrows set deleted_at = $1 where id = $2", id)
	if err != nil {
		return err
	}
	return nil
}

func (b *BorrowRepo) GetAll(f *pb.BorrowFilterRequest) (*pb.BorrowsResponse, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset string
	)
	query := `select id, user_id,book_id,created_at,updated_at,deleted_at from borrow where deleted_at = 0 `
	filter := ` where true `

	if len(f.Userid) > 0 {
		params["user_id"] = f.Userid
		filter += "and user_id = :user_id "
	}

	if len(f.BookId) > 0 {
		params["book_id"] = f.BookId
		filter += "and book_id = :book_id "
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
	rows, err := b.Db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	var borrows []*pb.BorrowResponse
	for rows.Next() {
		var borrow pb.BorrowResponse
		err := rows.Scan(&borrow.Id, &borrow.Userid, &borrow.BookId, borrow.CreatedAt, borrow.UpdatedAt, borrow.DeletedAt)
		if err != nil {
			return nil, err
		}
		borrows = append(borrows, &borrow)
	}
	return &pb.BorrowsResponse{BorrowsResponse: borrows}, nil
}
