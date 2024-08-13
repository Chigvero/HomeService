package repository

import (
	"Avito/internal/repository/postgres"
	"Avito/model"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	//DummyLogin(user_type string)
	Register(user model.UserRegister) (string, error)
	Login(user model.UserLogin) (string, error)
}

type House interface {
	Create(house model.House) (model.House, error)
	GetHouseModerFlatsList(houseId int) ([]model.Flat, error)
	GetHouseClientFlatsList(houseId int) ([]model.Flat, error)
	//SUBSCRIBE
}

type Flat interface {
	Create(flat model.Flat) (model.Flat, error)
	Update(id int, status string) (model.Flat, error)
}

type Repository struct {
	Authorization
	House
	Flat
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: postgres.NewAuthPostgres(db),
		House:         postgres.NewHousePostgres(db),
		Flat:          postgres.NewFlatPostgres(db),
	}
}
