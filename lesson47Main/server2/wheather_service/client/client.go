package main

import (
	pb "Go11Group/Kupalov-Muhammadjon/lesson46/wheather_service/genproto/WheatherService"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:44444", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewWheatherServiceClient(conn)

	resp, err := c.GetCurrentWeather(context.Background(), &pb.CurrentWheatherRequest{City: "Tashkent"})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
