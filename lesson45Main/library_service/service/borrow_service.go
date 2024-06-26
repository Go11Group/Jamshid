package service

import (
	"context"
	pb "library_service/genproto"
	strorage "library_service/storage"
)

type BorrowService struct {
	Serv strorage.StorageI
	pb.UnimplementedBorrowServiceServer
}

func NewBorrowService(st strorage.StorageI) *BorrowService {
	return &BorrowService{Serv: st}
}
func (serv *BorrowService) CreateBorrow(ctc context.Context, in *pb.CreateBorrowRequest) (*pb.Void, error) {
	return serv.Serv.BorrowI.CreateBorrow(in)
}
func (serv *BorrowService) UpdateBorrow(ctc context.Context, in *pb.UpdatedBorrowRequest) (*pb.Void, error) {
	return serv.Serv.BorrowI.UpdateBorrow(in)
}
func (serv *BorrowService) DeleteBorrow(ctc context.Context, in *pb.ByIdRequest) (*pb.Void, error) {
	return serv.Serv.BorrowI.DeleteBorrow(in)
}
func (serv *BorrowService) GetBorrow(ctc context.Context, in *pb.BorrowFilterRequest) (*pb.BorrowsResponse, error) {
	return serv.Serv.BorrowI.GeBorrow(in)
}
