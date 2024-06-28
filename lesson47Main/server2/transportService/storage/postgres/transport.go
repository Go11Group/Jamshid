package postgres

import (
	"database/sql"
	"errors"
	"github.com/lib/pq"
	"math/rand"
)

type TransportRepo struct {
	DB          *sql.DB
	TrafficJams []string
}

func NewTransportRepo(db *sql.DB) *TransportRepo {
	return &TransportRepo{DB: db}
}

func (t *TransportRepo) CreateTrafficjam(report string) error {
	t.TrafficJams = append(t.TrafficJams, report)
	return nil
}

func (t *TransportRepo) GetBusSchedule(busNumber int32) ([]string, error) {
	schedule := []string{}
	err := t.DB.QueryRow("SELECT stations FROM transports WHERE bus_number = $1", busNumber).Scan(pq.Array(&schedule))
	return schedule, err
}

func (t *TransportRepo) TrackBusLocation(busNumber int32) (string, error) {
	schedule := []string{}
	err := t.DB.QueryRow("SELECT stations FROM transports WHERE bus_number = $1", busNumber).Scan(pq.Array(&schedule))
	if err != nil {
		return "", err
	}
	if len(schedule) == 0 {
		return "", errors.New("no location found")
	}
	randomIndex := rand.Intn(len(schedule))
	return schedule[randomIndex], nil
}
