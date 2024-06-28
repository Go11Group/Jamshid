package main

import (
	pb "Go11Group/Kupalov-Muhammadjon/lesson46/wheather_service/genproto/WheatherService"
	"Go11Group/Kupalov-Muhammadjon/lesson46/wheather_service/service"
	"Go11Group/Kupalov-Muhammadjon/lesson46/wheather_service/storage/postgres"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":44444")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	wService := service.NewWheatherService(db)
	s := grpc.NewServer()
	pb.RegisterWheatherServiceServer(s, wService)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
