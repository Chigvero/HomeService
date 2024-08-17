package service

import (
	"HomeService/internal/repository"
	"HomeService/model"
	"github.com/google/uuid"
)

type FlatService struct {
	repos repository.Flat
}

func NewFlatService(repos repository.Flat) *FlatService {
	return &FlatService{
		repos: repos,
	}
}

func (s *FlatService) Create(flat model.Flat) (model.Flat, error) {
	return s.repos.Create(flat)
}
func (s *FlatService) Update(id int, house_id int, status string, user_id uuid.UUID) (model.Flat, error) {
	return s.repos.Update(id, house_id, status, user_id)
}

func (s *FlatService) GetById(id, house_id int) (model.Flat, error) {
	return s.repos.GetById(id, house_id)
}
