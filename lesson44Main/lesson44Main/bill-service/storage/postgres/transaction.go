package postgres

import (
	"bill_service/models"
	strorage "bill_service/storage"
	"database/sql"
	"fmt"
	"time"
)

type TransactionRepository struct {
	Db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{
		Db: db,
	}
}

func (repo *TransactionRepository) CreateTransaction(transaction *models.Transaction) (bool, error) {
	if transaction.TransactionType == "deposit" {
		query := "select  sum(amount)  as totaly_summ from  transactions where transaction_type='deposit' or transaction_type='credit'"
		var totalSum int
		err := repo.Db.QueryRow(query).Scan(&totalSum)
		if err != nil {
			fmt.Println("++++++++", err)
			return true, nil

		}
		if totalSum <= 0 {
			fmt.Println("+++++++++", transaction.Amount)
			return false, nil
		}
		return true, nil
	}

	_, err := repo.Db.Exec("insert into transactions(card_id,amount,terminal_id,transaction_type, created_at)values ($1,$2,$3,$4,$5)", transaction.CardId, transaction.Amount, transaction.TerminalId, transaction.TransactionType, time.Now())
	return true, err
}
func (repo *TransactionRepository) UpdateTransaction(id string, transaction *models.Transaction) error {
	_, err := repo.Db.Exec("update transactions set card_id=$1,amount=$2,terminal_id=$3,transaction_type=$4,updated_at=$5 where id=$6", transaction.CardId, transaction.Amount, transaction.TerminalId, transaction.TransactionType, time.Now(), id)
	return err
}
func (repo *TransactionRepository) DeletedTransaction(id string) error {
	_, err := repo.Db.Exec("update transactions set deleted_at=$1 where id=$2", 1, id)
	return err
}
func (repo *TransactionRepository) GetTransaction(transactionFilter models.Filter) (*[]models.Transaction, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset string
	)
	query := "select id,card_id,amount,terminal_id,transaction_type,created_at,updated_at from transactions where  deleted_at=0 "
	filter := ""
	if len(transactionFilter.CardId) > 0 {
		params["card_id"] = transactionFilter.CardId
		filter += " and card_id = :card_id "

	}
	if len(transactionFilter.TerminalId) > 0 {
		params["terminal_id"] = transactionFilter.TerminalId
		filter += " and terminal_id = :terminal_id "

	}
	if len(transactionFilter.TransactionType) > 0 {
		params["transaction_type"] = transactionFilter.TransactionType
		filter += " and transaction_type = :transaction_type "

	}
	if transactionFilter.Amount > 0 {
		params["amount"] = transactionFilter.Amount
		filter += " and amount = :amount "

	}

	if transactionFilter.Limit > 0 {
		params["limit"] = transactionFilter.Limit
		limit = ` LIMIT :limit`

	}
	if transactionFilter.Offset > 0 {
		params["offset"] = transactionFilter.Offset
		limit = ` OFFSET :offset`

	}
	query = query + filter + limit + offset
	query, arr = strorage.ReplaceQueryParams(query, params)
	fmt.Println("++++++++++", query)
	rows, err := repo.Db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	var transactions []models.Transaction
	for rows.Next() {
		var transaction models.Transaction
		err := rows.Scan(&transaction.Id, &transaction.CardId, &transaction.Amount, &transaction.TerminalId, &transaction.TransactionType, &transaction.CreatedAt, &transaction.UpdatedAt)
		if err != nil {
			return nil, err
		}
		transaction.Amount = (-1) * transaction.Amount
		transactions = append(transactions, transaction)
	}
	return &transactions, err

}

func (repo *TransactionRepository) GetTransactionById(id string) (*models.Transaction, error) {
	var transaction models.Transaction
	rows, err := repo.Db.Query("select id,card_id,amount,terminal_id,transaction_type,created_at,updated_at,deleted_at from transactions where id=$1 and deleted_at=0", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&transaction.Id, &transaction.CardId, &transaction.Amount, &transaction.TerminalId, &transaction.TransactionType, &transaction.CreatedAt, &transaction.UpdatedAt, &transaction.DeletedAt)
		if err != nil {
			return nil, err
		}
		return &transaction, err
	}
	return nil, err
}
