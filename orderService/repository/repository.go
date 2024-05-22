package repository

import "database/sql"

type OrderRepository struct {
	DB *sql.DB
}

func NewOrderRepository(DB *sql.DB) OrderRepositoryInterface {
	return OrderRepository{
		DB: DB,
	}
}

