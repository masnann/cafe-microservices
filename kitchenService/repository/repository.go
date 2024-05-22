package repository

import "database/sql"

type KitchenRepository struct {
	DB *sql.DB
}

func NewKitchenRepository(DB *sql.DB) KitchenRepositoryInterface {
	return KitchenRepository{
		DB: DB,
	}
}

