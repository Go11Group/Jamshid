package main

import (
	pb "Go11Group/Kupalov-Muhammadjon/lesson46/transportService/genproto/TransportService"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:55555", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := pb.NewTransportServiceClient(conn)
	resp, err := c.TrackBusLocation(context.Background(), &pb.BusLocationRequest{BusNumber: 114})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
