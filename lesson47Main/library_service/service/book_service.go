package service

import (
	"context"
	pb "library_service/genproto"
	strorage "library_service/storage"
)

type BookService struct {
	Serv strorage.StorageI
	pb.UnimplementedLibraryServiceServer
}

func NewBookService(st strorage.StorageI) *BookService {
	return &BookService{Serv: st}
}

func (serv *BookService) CreatBook(ctc context.Context, in *pb.CreateBookRequest) (*pb.Void, error) {
	return serv.Serv.BookI.CreateBook(in)
}
func (serv *BookService) UpdateBook(ctc context.Context, in *pb.UpdatedBookRequest) (*pb.Void, error) {
	return serv.Serv.BookI.UpdateBook(in)
}
func (serv *BookService) DeleteBook(ctc context.Context, in *pb.ByIdRequest) (*pb.Void, error) {
	return serv.Serv.BookI.DeleteBook(in)
}
func (serv *BookService) GetBook(ctc context.Context, in *pb.BookFilterRequest) (*pb.BooksResponse, error) {
	return serv.Serv.BookI.GeBook(in)
}
