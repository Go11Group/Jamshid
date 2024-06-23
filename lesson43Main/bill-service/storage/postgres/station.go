package postgres

import (
	"bill_service/models"
	strorage "bill_service/storage"
	"database/sql"
	"fmt"
	"time"
)

type StationRepository struct {
	Db *sql.DB
}

func NewStationRepository(db *sql.DB) *StationRepository {
	return &StationRepository{
		Db: db,
	}
}

func (repo *StationRepository) CreateStation(station *models.Station) error {
	_, err := repo.Db.Exec("insert into stations(name,created_at)values ($1,$2)", station.Name, time.Now())
	return err
}
func (repo *StationRepository) UpdateStation(id string, station *models.Station) error {
	_, err := repo.Db.Exec("update stations set name=$1,updated_at=$2 where id=$3", station.Name, time.Now(), id)
	return err
}
func (repo *StationRepository) DeletedStation(id string) error {
	_, err := repo.Db.Exec("update stations set deleted_at=$1 where id=$2", 1, id)
	return err
}
func (repo *StationRepository) GetStation(station models.Filter) (*[]models.Station, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset string
	)
	query := "select id,name,created_at,updated_at from stations where  deleted_at=0 "
	filter := ""
	if len(station.Name) > 0 {
		params["name"] = station.Name
		filter += " and name = :name "

	}

	if station.Limit > 0 {
		params["limit"] = station.Limit
		limit = ` LIMIT :limit`

	}
	if station.Offset > 0 {
		params["offset"] = station.Offset
		limit = ` OFFSET :offset`

	}
	query = query + filter + limit + offset

	query, arr = strorage.ReplaceQueryParams(query, params)
	rows, err := repo.Db.Query(query, arr...)
	fmt.Println("+++++++++", arr)
	fmt.Println("+++++++++", query)
	if err != nil {
		return nil, err
	}
	var stations []models.Station
	for rows.Next() {
		var station models.Station
		err := rows.Scan(&station.Id, &station.Name, &station.CreatedAt, &station.UpdatedAt)
		if err != nil {
			return nil, err
		}
		stations = append(stations, station)
	}
	return &stations, err

}

func (repo *StationRepository) GetStationById(id string) (*models.Station, error) {
	var station models.Station
	rows, err := repo.Db.Query("select id,name,created_at,updated_at,deleted_at from stations where id=$1 and deleted_at=0", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&station.Id, &station.Name, &station.CreatedAt, &station.UpdatedAt, &station.DeletedAt)
		if err != nil {
			return nil, err
		}
		return &station, err
	}
	return nil, err
}

func (repo *StationRepository) GetTerminalsByStationId(id string) (*[]models.StationTerminal, error) {
	var stationTerminal models.StationTerminal
	rows, err := repo.Db.Query("select s.id,s.name,t.id from stations  s inner join  terminals t on s.id=t.station_id where s.id=$1 and s.deleted_at=0 and t.deleted_at=0", id)
	if err != nil {

		return nil, err
	}
	var stationTerminals []models.StationTerminal
	for rows.Next() {
		err := rows.Scan(&stationTerminal.StationId, &stationTerminal.StationName, &stationTerminal.TerminalId)
		if err != nil {
			fmt.Println("-----", err)
			return nil, err
		}
		stationTerminals = append(stationTerminals, stationTerminal)
	}
	return &stationTerminals, err
}
