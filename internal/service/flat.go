package service

import (
	"Avito/internal/repository"
	"Avito/model"
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
func (s *FlatService) Update(id int, status string, user_id uuid.UUID) (model.Flat, error) {
	return s.repos.Update(id, status, user_id)
}

func (s *FlatService) GetById(id int) (model.Flat, error) {
	return s.repos.GetById(id)
}
