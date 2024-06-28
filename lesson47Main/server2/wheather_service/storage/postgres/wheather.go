package postgres

import (
	pb "Go11Group/Kupalov-Muhammadjon/lesson46/wheather_service/genproto/WheatherService"
	"database/sql"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type WheatherRepo struct {
	Db *sql.DB
}

func NewWheatherRepo(db *sql.DB) *WheatherRepo {
	return &WheatherRepo{Db: db}
}

func (w *WheatherRepo) GetCurrentWheather(city string) (*pb.CurrentWheatherResponse, error) {
	cur := &pb.CurrentWheatherResponse{}
	query := `
		SELECT
			local_time, country, tempC, windKmph, humidity
		FROM
			wheather
		WHERE 
			city = $1
	`
	var localTime time.Time

	row := w.Db.QueryRow(query, city)
	err := row.Scan(&localTime, &cur.Country, &cur.TempC, &cur.WindKmph, &cur.Humidity)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	cur.LocalTime = timestamppb.New(localTime)

	return cur, nil
}

func (w *WheatherRepo) CreateWheatherReport(request *pb.ReportWheatherRequest) (bool, error) {
	query := `
	insert into wheather(city, country, local_time, tempC, windKmph, humidity)
values
    ($1, $2, $3, $4, $5, $6)
`
	_, err := w.Db.Exec(query, request.City, request.Country, request.LocalTime.AsTime(),
		request.TempC, request.WindKmph, request.Humidity)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (w *WheatherRepo) GetWheatherForecast(days int32, city string) (*[]*pb.Forecast, error) {
	forecasts := []*pb.Forecast{}
	query := `
	select
		local_time, country, tempC, windKmph, humidity
	from
	    wheather
	where 
	    city = $1
	limit $2
`
	rows, err := w.Db.Query(query, city, days)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		cur := pb.Forecast{}
		err := rows.Scan(&cur.LocalTime, &cur.Country, &cur.WindKmph, &cur.TempC)
		if err != nil {
			return nil, err
		}
		forecasts = append(forecasts, &cur)
	}

	return &forecasts, nil
}
