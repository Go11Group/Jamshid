package postgres

import (
	"database/sql"
	"lesson1/model"
)

type CustomerRepository struct {
	Db *sql.DB
}

func NewRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{Db: db}
}

func (dbC *CustomerRepository) CreateCustomer(customer model.Customer) error {
	transaction, err := dbC.Db.Begin()
	if err != nil {
		return err
	}
	_, err = dbC.Db.Exec("insert into customers(username,email,password) values ($1,$2,$3)", customer.Username, customer.Email, customer.Password)
	if err != nil {
		err := (transaction.Rollback())
		if err != nil {
			return err
		}
		return err
	}
	panic(transaction.Commit())
	return nil

}
func (dbC *CustomerRepository) DeleteCustomer(id int) error {
	transaction, err := dbC.Db.Begin()
	if err != nil {
		return err
	}
	_, err = dbC.Db.Exec("delete from customers where id=$1", id)
	if err != nil {
		err := (transaction.Rollback())
		if err != nil {
			return err
		}
		return err
	}
	err = (transaction.Commit())
	if err != nil {
		return err
	}
	return nil

}
func (dbC *CustomerRepository) UpdatedCustomer(id int, customer *model.Customer) error {
	transaction, err := dbC.Db.Begin()
	if err != nil {
		return err
	}

	_, err = dbC.Db.Exec("update customers set username=$1,email=$2,password=$3 where id=$4", customer.Username, customer.Email, customer.Password, id)
	if err != nil {
		panic(transaction.Rollback())
		return err
	}
	err = (transaction.Commit())
	if err != nil {
		return err
	}
	return nil

}
func (dbC *CustomerRepository) FindAllCustomer() ([]model.Customer, error) {
	transaction, err := dbC.Db.Begin()
	if err != nil {
		return nil, err
	}
	customers := []model.Customer{}
	rows, err := dbC.Db.Query("select username,email,password from customers")
	if err != nil {
		err := (transaction.Rollback())
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	for rows.Next() {
		customer := model.Customer{}
		err := rows.Scan(&customer.Username, &customer.Email, &customer.Password)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)

	}
	err = (transaction.Commit())
	if err != nil {
		return nil, err
	}
	return customers, nil

}
