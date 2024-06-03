package postgres

import (
	"database/sql"
	"lesson1/model"
)

type ProductRepository struct {
	Db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{Db: db}

}

func (dbP *ProductRepository) CreateProduct(product model.Product) error {
	transaction, err := dbP.Db.Begin()
	if err != nil {
		return err
	}

	_, err = dbP.Db.Exec("insert into products(name,description,price,stock_quantity) values ($1,$2,$3,$4)", product.Name, product.Description, product.Price, product.StockQuantity)
	if err != nil {
		err = transaction.Rollback()
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
func (dbP *ProductRepository) DeleteProduct(id int) error {
	transaction, err := dbP.Db.Begin()
	if err != nil {
		return err
	}
	_, err = dbP.Db.Exec("delete from products where id=$1", id)
	if err != nil {
		err := transaction.Rollback()
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
func (dbP *ProductRepository) UpdatedProduct(id int, product *model.Product) error {
	transaction, err := dbP.Db.Begin()
	if err != nil {
		return err
	}

	_, err = dbP.Db.Exec("update products set name=$1,description=$2,price=$3,stock_quantity where id=$4", product.Name, product.Description, product.Price, product.StockQuantity, id)
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
func (dbP *ProductRepository) FindAllProduct() ([]model.Product, error) {
	products := []model.Product{}
	transaction, err := dbP.Db.Begin()
	if err != nil {
		return nil, err
	}
	rows, err := dbP.Db.Query("select name,description,price,stock_quantity from products")
	if err != nil {
		err := (transaction.Rollback())
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	for rows.Next() {
		product := model.Product{}
		err := rows.Scan(&product.Name, &product.Description, &product.Price, &product.StockQuantity)
		if err != nil {
			return nil, err
		}
		products = append(products, product)

	}
	err = (transaction.Commit())
	if err != nil {
		return nil, err
	}
	return products, err

}
