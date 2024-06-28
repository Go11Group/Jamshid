package handler

import (
	pb "api_getway/genproto"
	"google.golang.org/grpc"
)

type Handler struct {
	transport pb.TransportServiceClient
	weather   pb.WheatherServiceClient
}

func NewHandlerStruct(cl *grpc.ClientConn) *Handler {
	return &Handler{
		borrow: pb.NewBorrowServiceClient(cl),
		user:   pb.NewUserServiceClient(cl),
		book:   pb.NewLibraryServiceClient(cl),
	}
}
