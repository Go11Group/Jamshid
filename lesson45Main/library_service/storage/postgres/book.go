package postgres

import (
	"database/sql"
	pb "library_service/genproto"
	strorage "library_service/storage"
	"time"
)

type BookRepo struct {
	Db *sql.DB
}

func (b *BookRepo) CreateBook(user *pb.CreateBookRequest) (*pb.Void, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BookRepo) UpdateBook(user *pb.UpdatedBookRequest) (*pb.Void, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BookRepo) DeleteBook(user *pb.ByIdRequest) (*pb.Void, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BookRepo) GeBook(user *pb.BookFilterRequest) (*pb.BooksResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewBookRepo(db *sql.DB) *BookRepo {
	return &BookRepo{Db: db}
}

func (b *BookRepo) Create(book *pb.CreateBookRequest) error {
	_, err := b.Db.Exec("insert into book (title,author, published, created_at) values ($1,$2,$3,$4)",
		book.Title, book.Author, book.Published, time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (b *BookRepo) GetById(id *pb.ByIdRequest) (*pb.BookResponse, error) {
	var book pb.BookResponse

	err := b.Db.QueryRow("select title,author,published from book where id = $1", id).Scan(&book.Title, &book.Author, &book.Published)
	if err != nil {
		return nil, err
	}
	return &book, err
}

func (b *BookRepo) Update(book *pb.BookResponse, id string) error {
	_, err := b.Db.Exec("update book set tittle = $1,author = $2,published = $3,updated_at = $4 where id = $5 and deleted_at = 0", book.Title, book.Author, book.Published, time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}

func (b *BookRepo) Delete(id string) error {
	_, err := b.Db.Exec("update book set deleted_at = $1 where id = $2", id)
	if err != nil {
		return err
	}
	return nil
}

func (c *BookRepo) GetAll(b *pb.BookFilterRequest) (*pb.BooksResponse, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset string
	)
	query := `select id, title,author, published, created_at,updated_at,deleted_at from book where deleted_at = 0 `
	filter := ` where true `

	if len(b.Title) > 0 {
		params["title"] = b.Title
		filter += "and title = :title "
	}

	if len(b.Author) > 0 {
		params["author"] = b.Author
		filter += "and author = :author "
	}
	if (b.Published) > 0 {
		params["published"] = b.Published
		filter += "and published = :published "
	}

	if (b.LimitOffset.Limit) > 0 {
		params["limit"] = b.LimitOffset.Offset
		filter += "and limit = :limit "
	}

	if (b.LimitOffset.Offset) > 0 {
		params["offset"] = b.LimitOffset.Offset
		filter += "and offset = :offset "
	}

	query = query + filter + limit + offset
	query, arr = strorage.ReplaceQueryParams(query, params)
	rows, err := c.Db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	var books []*pb.BookResponse
	for rows.Next() {
		var book pb.BookResponse
		err := rows.Scan(&book.Id, &book.Title, &book.Author, book.Published, book.CreatedAt, book.UpdatedAt, book.DeletedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}
	return &pb.BooksResponse{BooksResponse: books}, err
}
