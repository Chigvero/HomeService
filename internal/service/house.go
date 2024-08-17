package service

import (
	"HomeService/internal/repository"
	"HomeService/model"
)

type HouseService struct {
	repos repository.House
}

func NewHouseService(repos repository.House) *HouseService {
	return &HouseService{
		repos: repos,
	}
}

func (s *HouseService) Create(house model.House) (model.House, error) {
	return s.repos.Create(house)
}

func (s *HouseService) GetHouseModerFlatsList(houseId int) ([]model.Flat, error) {
	return s.repos.GetHouseModerFlatsList(houseId)
}
func (s *HouseService) GetHouseClientFlatsList(houseId int) ([]model.Flat, error) {
	return s.repos.GetHouseClientFlatsList(houseId)
}
