package main

import (
	pb "Go11Group/Kupalov-Muhammadjon/lesson46/transportService/genproto/TransportService"
	"Go11Group/Kupalov-Muhammadjon/lesson46/transportService/service"
	"Go11Group/Kupalov-Muhammadjon/lesson46/transportService/storage/postgres"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":55555")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	trService := service.NewTransportService(db)
	s := grpc.NewServer()
	pb.RegisterTransportServiceServer(s, trService)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
