package service

import (
	pb "Go11Group/Kupalov-Muhammadjon/lesson46/wheather_service/genproto/WheatherService"
	"Go11Group/Kupalov-Muhammadjon/lesson46/wheather_service/storage/postgres"
	"context"
	"database/sql"
)

type WheatherService struct {
	pb.UnimplementedWheatherServiceServer
	WheatherRepo *postgres.WheatherRepo
}

func NewWheatherService(db *sql.DB) *WheatherService {
	w := postgres.NewWheatherRepo(db)
	return &WheatherService{WheatherRepo: w}
}

func (w *WheatherService) GetCurrentWeather(ctx context.Context, r *pb.CurrentWheatherRequest) (*pb.CurrentWheatherResponse, error) {
	cur, err := w.WheatherRepo.GetCurrentWheather(r.City)
	if err != nil {
		return nil, err
	}
	return cur, nil
}

func (w *WheatherService) GetWeatherForecast(ctx context.Context, r *pb.ForecastWheatherRequest) (*pb.ForecastWheatherResponse, error) {
	cur, err := w.WheatherRepo.GetWheatherForecast(r.Days, r.City)
	if err != nil {
		return nil, err
	}
	return &pb.ForecastWheatherResponse{Days: *cur}, nil
}

func (w *WheatherService) ReportWeatherCondition(ctx context.Context, r *pb.ReportWheatherRequest) (*pb.ReportWheatherResponse, error) {
	cur, err := w.WheatherRepo.CreateWheatherReport(r)
	if err != nil {
		return nil, err
	}
	return &pb.ReportWheatherResponse{IsReported: cur}, nil
}
