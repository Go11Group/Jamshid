package strorage

import (
	pb "library_service/genproto"
)

type StorageI struct {
	UserI   UserStorage
	BookI   BookStorage
	BorrowI BorrowStorage
}

func (s StorageI) CreateUser(user *pb.CreateUserRequest) (*pb.Void, error) {
	//TODO implement me
	panic("implement me")
}

func (s StorageI) UpdateUser(user *pb.UpdatedUserRequest) (*pb.Void, error) {
	//TODO implement me
	panic("implement me")
}

func (s StorageI) DeleteUser(user *pb.ByIdRequest) (*pb.Void, error) {
	//TODO implement me
	panic("implement me")
}

func (s StorageI) GetUser(user *pb.UserFilterRequest) (*pb.UsersResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s StorageI) CreateBorrow(user *pb.CreateBorrowRequest) (*pb.Void, error) {
	//TODO implement me
	panic("implement me")
}

func (s StorageI) UpdateBorrow(user *pb.UpdatedBorrowRequest) (*pb.Void, error) {
	//TODO implement me
	panic("implement me")
}

func (s StorageI) DeleteBorrow(user *pb.ByIdRequest) (*pb.Void, error) {
	//TODO implement me
	panic("implement me")
}

func (s StorageI) GeBorrow(user *pb.BorrowFilterRequest) (*pb.BorrowsResponse, error) {
	//TODO implement me
	panic("implement me")
}

type UserStorage interface {
	CreateUser(user *pb.CreateUserRequest) (*pb.Void, error)
	UpdateUser(user *pb.UpdatedUserRequest) (*pb.Void, error)
	DeleteUser(user *pb.ByIdRequest) (*pb.Void, error)
	GetUser(user *pb.UserFilterRequest) (*pb.UsersResponse, error)
}
type BookStorage interface {
	CreateBook(user *pb.CreateBookRequest) (*pb.Void, error)
	UpdateBook(user *pb.UpdatedBookRequest) (*pb.Void, error)
	DeleteBook(user *pb.ByIdRequest) (*pb.Void, error)
	GeBook(user *pb.BookFilterRequest) (*pb.BooksResponse, error)
}
type BorrowStorage interface {
	CreateBorrow(user *pb.CreateBorrowRequest) (*pb.Void, error)
	UpdateBorrow(user *pb.UpdatedBorrowRequest) (*pb.Void, error)
	DeleteBorrow(user *pb.ByIdRequest) (*pb.Void, error)
	GeBorrow(user *pb.BorrowFilterRequest) (*pb.BorrowsResponse, error)
}
