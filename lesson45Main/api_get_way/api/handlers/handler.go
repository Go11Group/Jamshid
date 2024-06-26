package handler

import (
	//pb "api_get_way/genproto"
	pb "get-way/genproto"

	"google.golang.org/grpc"
)

type Handler struct {
	borrow pb.BorrowServiceClient
	user   pb.UserServiceClient
	book   pb.LibraryServiceClient
}

func NewHandlerStruct(cl *grpc.ClientConn) *Handler {
	return &Handler{
		borrow: pb.NewBorrowServiceClient(cl),
		user:   pb.NewUserServiceClient(cl),
		book:   pb.NewLibraryServiceClient(cl),
	}
}
