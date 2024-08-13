package service

import (
	"Avito/internal/repository"
	"Avito/model"
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
func (s *FlatService) Update(id int, status string) (model.Flat, error) {
	return s.repos.Update(id, status)
}
