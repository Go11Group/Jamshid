package main

import (
	"google.golang.org/grpc"
	pb "library_service/genproto"
	"library_service/service"
	"library_service/storage/postgres"

	//"library_service/storage/postgres"
	"log"
	"net"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	listen, err := net.Listen("tcp", ":8088")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterBorrowServiceServer(s, service.NewBorrowService(db))
	pb.RegisterLibraryServiceServer(s, service.NewBookService(db))
	pb.RegisterUserServiceServer(s, service.NewUserService(db))

	log.Printf("server listening at %v", listen.Addr())

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
