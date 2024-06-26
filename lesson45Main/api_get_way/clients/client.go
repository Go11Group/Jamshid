package clients

import (
	pb "get-way/genproto"
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientI struct {
	UserClient   pb.UserServiceClient
	BookClient   pb.LibraryServiceClient
	BorrowClient pb.BorrowServiceClient
}

func NewClient() *ClientI {
	conn, err := grpc.NewClient("localhost:8088", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		slog.Error("Error while dialing", err)
	}

	eduC := pb.NewUserServiceClient(conn)
	exC := pb.NewLibraryServiceClient(conn)
	prC := pb.NewBorrowServiceClient(conn)

	return &ClientI{
		UserClient:   eduC,
		BookClient:   exC,
		BorrowClient: prC,
	}

}
