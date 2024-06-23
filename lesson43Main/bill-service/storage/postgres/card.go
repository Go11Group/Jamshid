package postgres

import "C"
import (
	"bill_service/models"
	strorage "bill_service/storage"
	"database/sql"
	"fmt"
	"time"
)

type CardRepository struct {
	Db *sql.DB
}

func NewCardRepository(db *sql.DB) *CardRepository {
	return &CardRepository{
		Db: db,
	}
}

func (repo *CardRepository) CreateCard(card *models.Card) error {
	_, err := repo.Db.Exec("insert into cards(number,user_id,created_at)values ($1,$2,$3)", card.Number, card.Userid, time.Now())
	return err
}
func (repo *CardRepository) UpdateCard(id string, card *models.Card) error {
	_, err := repo.Db.Exec("update cards set number=$1,user_id=$2,updated_at=$3 where id=$4", card.Number, card.Userid, time.Now(), id)
	return err
}
func (repo *CardRepository) DeletedCard(id string) error {
	_, err := repo.Db.Exec("update cards set deleted_at=$1 where id=$2", 1, id)
	return err
}
func (repo *CardRepository) GetCard(cardFilter models.Filter) (*[]models.Card, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset string
	)
	query := "select id,number,user_id,created_at,updated_at from cards where  deleted_at=0 "
	filter := ""
	if len(cardFilter.Number) > 0 {
		params["number"] = cardFilter.Number
		filter += " and number = :number "

	}

	if cardFilter.Limit > 0 {
		params["limit"] = cardFilter.Limit
		limit = ` LIMIT :limit`

	}
	if cardFilter.Offset > 0 {
		params["offset"] = cardFilter.Offset
		limit = ` OFFSET :offset`

	}
	query = query + filter + limit + offset
	query, arr = strorage.ReplaceQueryParams(query, params)
	rows, err := repo.Db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	var cards []models.Card
	for rows.Next() {
		var card models.Card
		err := rows.Scan(&card.Id, &card.Number, &card.Userid, &card.CreatedAt, &card.UpdatedAt)
		if err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}
	return &cards, err

}

func (repo *CardRepository) GetCardById(id string) (*models.Card, error) {
	var card models.Card
	rows, err := repo.Db.Query("select *from where id=$1 and deleted_at=0", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&card.Id, &card.Number, &card.Userid, &card.CreatedAt, &card.UpdatedAt, &card.DeletedAt)
		if err != nil {
			return nil, err
		}
		return &card, err
	}
	return nil, err
}
func (repo *CardRepository) GetStationByCardId(id string) (*[]models.CardTransactionStation, error) {
	var cardTransactionStations []models.CardTransactionStation
	rows, err := repo.Db.Query("select  c.id,c.number,t.id,t.transaction_type,s.name from  cards c inner join   transactions t on c.id=t.card_id inner join terminals  on t.terminal_id = terminals.id inner join stations s on terminals.station_id = s.id where c.id=$1 and c.deleted_at=0 and t.deleted_at=0 and terminals.deleted_at=0 and s.deleted_at=0", id)
	if err != nil {

		return nil, err
	}
	for rows.Next() {
		var cardTransactionStation *models.CardTransactionStation
		err = rows.Scan(&cardTransactionStation.CardId, &cardTransactionStation.CardNumber, &cardTransactionStation.TransactionId, &cardTransactionStation.TransactionType, &cardTransactionStation.StationName, &cardTransactionStation)
		if err != nil {
			fmt.Println("++++++++++++", err)
			return nil, err
		}
		cardTransactionStations = append(cardTransactionStations, *cardTransactionStation)

	}
	return &cardTransactionStations, err
}
