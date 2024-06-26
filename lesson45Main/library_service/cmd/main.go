package main

import (
	"google.golang.org/grpc"
	"library_service/config"
	pb "library_service/genproto"
	"library_service/service"
	strorage "library_service/storage"

	//"library_service/storage/postgres"
	"log"
	"net"
)

func main() {
	db, err := strorage.ConnectionDb(config.Config{})
	if err != nil {
		panic(err)
	}
	listen, err := net.Listen("tcp", ":8088")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterBorrowServiceServer(s, service.NewBorrowService(db))
	pb.RegisterLibraryServiceServer(s, service.NewService(db))
	pb.RegisterLibraryServiceServer(s, service.NewService(db))

	log.Printf("server listening at %v", listen.Addr())

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
