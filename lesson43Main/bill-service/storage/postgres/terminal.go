package postgres

import (
	"bill_service/models"
	strorage "bill_service/storage"
	"database/sql"
	"time"
)

type TerminalRepository struct {
	Db *sql.DB
}

func NewTerminalRepository(db *sql.DB) *TerminalRepository {
	return &TerminalRepository{
		Db: db,
	}
}

func (repo *TerminalRepository) CreateTerminal(terminal *models.Terminal) error {
	_, err := repo.Db.Exec("insert into terminals(station_id,created_at)values ($1,$2)", terminal.StationId, time.Now())
	return err
}
func (repo *TerminalRepository) UpdateTerminal(id string, terminal *models.Terminal) error {
	_, err := repo.Db.Exec("update terminals set station_id=$1,updated_at=$2 where id=$3", terminal.StationId, time.Now(), id)
	return err
}
func (repo *TerminalRepository) DeletedTerminal(id string) error {
	_, err := repo.Db.Exec("update terminals set deleted_at=$1 where id=$2", 1, id)
	return err
}
func (repo *TerminalRepository) GetTerminal(terminalFilter models.Filter) (*[]models.Terminal, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset string
	)
	query := "select id,station_id,created_at,updated_at from terminals where  deleted_at=0 "
	filter := ""
	if len(terminalFilter.StationId) > 0 {
		params["station_id"] = terminalFilter.StationId
		filter += " and station_id = :station_id "

	}

	if terminalFilter.Limit > 0 {
		params["limit"] = terminalFilter.Limit
		limit = ` LIMIT :limit`

	}
	if terminalFilter.Offset > 0 {
		params["offset"] = terminalFilter.Offset
		limit = ` OFFSET :offset`

	}
	query = query + filter + limit + offset
	query, arr = strorage.ReplaceQueryParams(query, params)
	rows, err := repo.Db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	var terminals []models.Terminal
	for rows.Next() {
		var terminal models.Terminal
		err := rows.Scan(&terminal.Id, &terminal.StationId, &terminal.CreatedAt, &terminal.UpdatedAt)
		if err != nil {
			return nil, err
		}
		terminals = append(terminals, terminal)
	}
	return &terminals, err

}

func (repo *TerminalRepository) GetTerminalById(id string) (*models.Terminal, error) {
	var terminal models.Terminal
	rows, err := repo.Db.Query("select id,station_id,created_at,updated_at,deleted_at from terminals where id=$1 and deleted_at=0", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&terminal.Id, &terminal.StationId, &terminal.CreatedAt, &terminal.UpdatedAt, &terminal.DeletedAt)
		if err != nil {
			return nil, err
		}
		return &terminal, err
	}
	return nil, err
}
