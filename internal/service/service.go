package service

import (
	"HomeService/internal/repository"
	"HomeService/model"
	_ "github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	Register(user model.UserRegister) (string, error)
	Login(user model.UserLogin) (string, error)
	DummyLogin(user_type string, user_id uuid.UUID) (string, error)
	//GenerateToken(user_type string, userId uuid.UUID ) (string, error)
	ParseToken(tokenString string) (model.UserLogin, error)
}

type House interface {
	Create(house model.House) (model.House, error)
	GetHouseModerFlatsList(houseId int) ([]model.Flat, error)
	GetHouseClientFlatsList(houseId int) ([]model.Flat, error)
	//SUBSCRIBE
}

type Flat interface {
	Create(flat model.Flat) (model.Flat, error)
	Update(id int, house_id int, status string, user_id uuid.UUID) (model.Flat, error)
	GetById(id, house_id int) (model.Flat, error)
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
