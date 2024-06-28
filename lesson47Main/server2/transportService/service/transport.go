package service

import (
	pb "Go11Group/Kupalov-Muhammadjon/lesson46/transportService/genproto/TransportService"
	"Go11Group/Kupalov-Muhammadjon/lesson46/transportService/storage/postgres"
	"context"
	"database/sql"
	"fmt"
)

type TransportService struct {
	pb.UnimplementedTransportServiceServer
	TransportRepo *postgres.TransportRepo
}

func NewTransportService(db *sql.DB) *TransportService {
	tr := postgres.NewTransportRepo(db)
	return &TransportService{TransportRepo: tr}
}

func (t *TransportService) GetBusSchedule(ctx context.Context, r *pb.BusScheduleRequest) (*pb.BusScheduleResponse, error) {
	st, err := t.TransportRepo.GetBusSchedule(r.BusNumber)
	if err != nil {
		return nil, err
	}
	return &pb.BusScheduleResponse{Stations: st}, nil
}
func (t *TransportService) TrackBusLocation(ctx context.Context, r *pb.BusLocationRequest) (*pb.BusLocationResponse, error) {
	st, err := t.TransportRepo.TrackBusLocation(r.BusNumber)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &pb.BusLocationResponse{Station: st}, nil
}
func (t *TransportService) ReportTrafficJam(ctx context.Context, r *pb.TrafficJamRequest) (*pb.TrafficJamResponse, error) {
	err := t.TransportRepo.CreateTrafficjam(r.Report)
	if err != nil {
		return &pb.TrafficJamResponse{IsReported: false}, err
	}
	return &pb.TrafficJamResponse{IsReported: true}, nil
}
