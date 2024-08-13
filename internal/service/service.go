package service

import (
	"Avito/internal/repository"
	"Avito/model"
)

type Authorization interface {
	Register(user model.UserRegister) (string, error)
	Login(user model.UserLogin) (string, error)
	DummyLogin() string

	ParseToken(tokenString string) (string, error)
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

type Service struct {
	Authorization
	House
	Flat
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization),
		House:         NewHouseService(repository.House),
		Flat:          NewFlatService(repository.Flat),
	}
}
