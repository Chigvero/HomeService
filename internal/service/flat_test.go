package service

import (
	"HomeService/model"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockFlatRepo struct {
	mock.Mock
}

func (m *MockFlatRepo) Create(flat model.Flat) (model.Flat, error) {
	args := m.Called(flat)
	return args.Get(0).(model.Flat), args.Error(1)
}

func (m *MockFlatRepo) Update(id int, house_id int, status string, user_id uuid.UUID) (model.Flat, error) {
	args := m.Called(id, house_id, status, user_id)
	return args.Get(0).(model.Flat), args.Error(1)
}

func (m *MockFlatRepo) GetById(id, house_id int) (model.Flat, error) {
	args := m.Called(id, house_id)
	return args.Get(0).(model.Flat), args.Error(1)
}

func TestFlatService_CreateSuccess(t *testing.T) {
	mockRepo := new(MockFlatRepo)
	services := NewFlatService(mockRepo)
	flat := model.Flat{
		Id:      113,
		HouseId: 1,
		Price:   120000,
		Rooms:   3,
	}
	expectedFlat := flat
	mockRepo.On("Create", flat).Return(expectedFlat, nil)
	createdFlat, err := services.repos.Create(flat)
	assert.NoError(t, err)
	assert.Equal(t, expectedFlat, createdFlat)
	mockRepo.AssertExpectations(t)
}

func TestFlatService_CreateError(t *testing.T) {
	mockRepo := new(MockFlatRepo)
	services := NewFlatService(mockRepo)

	flat := model.Flat{
		Id:      113,
		HouseId: 1,
		Price:   120000,
		Rooms:   3,
	}

	expectedError := errors.New("some create error")
	mockRepo.On("Create", flat).Return(model.Flat{}, expectedError)
	_, err := services.Create(flat)

	assert.Error(t, err)
	assert.True(t, errors.Is(err, expectedError))
	mockRepo.AssertExpectations(t)
}
