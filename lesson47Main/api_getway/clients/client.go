package clients

import (
	"log/slog"

	pb "api_getway/genproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientI struct {
	TransportClient pb.TransportServiceClient
	WeatherClient   pb.WheatherServiceClient
}

func NewClient() *ClientI {
	conn, err := grpc.NewClient("localhost:8088", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		slog.Error("Error while dialing", err)
	}

	weatherCon := pb.NewWheatherServiceClient(conn)
	transportCon := pb.NewTransportServiceClient(conn)

	return &ClientI{
		TransportClient: transportCon,
		WeatherClient:   weatherCon,
	}

}
