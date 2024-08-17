package service

import (
	"HomeService/model"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type MockHouseRepo struct {
	mock.Mock
}

func (m *MockHouseRepo) Create(house model.House) (model.House, error) {
	args := m.Called(house)
	return args.Get(0).(model.House), args.Error(1)
}

func (m *MockHouseRepo) GetHouseModerFlatsList(houseId int) ([]model.Flat, error) {
	args := m.Called(houseId)
	return args.Get(0).([]model.Flat), args.Error(1)
}

func (m *MockHouseRepo) GetHouseClientFlatsList(houseId int) ([]model.Flat, error) {
	args := m.Called(houseId)
	return args.Get(0).([]model.Flat), args.Error(1)
}

func TestHouseService_GetHouseClientFlatsList(t *testing.T) {
	mockRepo := new(MockHouseRepo)
	services := NewHouseService(mockRepo)
	houseID := 1
	expectedFlats := []model.Flat{
		{Id: 1, HouseId: houseID, Price: 100000, Rooms: 3, Status: "approved"},
		{Id: 2, HouseId: houseID, Price: 120000, Rooms: 4, Status: "approved"},
	}

	mockRepo.On("GetHouseClientFlatsList", houseID).Return(expectedFlats, nil)

	flats, err := services.GetHouseClientFlatsList(houseID)
	assert.NoError(t, err)
	assert.Equal(t, expectedFlats, flats)
	mockRepo.AssertExpectations(t)
}

func TestHouseService_GetHouseModerFlatsList(t *testing.T) {
	mockRepo := new(MockHouseRepo)
	services := NewHouseService(mockRepo)
	houseID := 1
	expectedFlats := []model.Flat{
		{Id: 1, HouseId: houseID, Price: 100000, Rooms: 3, Status: "approved"},
		{Id: 2, HouseId: houseID, Price: 120000, Rooms: 4, Status: "approved"},
	}

	mockRepo.On("GetHouseModerFlatsList", houseID).Return(expectedFlats, nil)

	flats, err := services.GetHouseModerFlatsList(houseID)
	assert.NoError(t, err)
	assert.Equal(t, expectedFlats, flats)
	mockRepo.AssertExpectations(t)
}

func TestHouseService_CreateSuccess(t *testing.T) {
	mockRepo := new(MockHouseRepo)
	services := NewHouseService(mockRepo)
	house := model.House{
		Id:        1,
		Address:   "Moscow, ul. Nikolskaya 1",
		Year:      2024,
		Developer: "PIK",
		CreatedAt: time.Now().String(),
		UpdateAt:  time.Now().String(),
	}
	expectedHouse := house
	mockRepo.On("Create", house).Return(expectedHouse, nil)
	createdHouse, err := services.Create(house)
	assert.NoError(t, err)
	assert.Equal(t, expectedHouse, createdHouse)
	mockRepo.AssertExpectations(t)
}

func TestHouseService_CreateError(t *testing.T) {
	mockRepo := new(MockHouseRepo)
	services := NewHouseService(mockRepo)
	house := model.House{
		Id:        1,
		Address:   "Moscow, ul. Nikolskaya 1",
		Year:      2024,
		Developer: "PIK",
		CreatedAt: time.Now().String(),
		UpdateAt:  time.Now().String(),
	}
	expectedError := errors.New("some create error")
	mockRepo.On("Create", house).Return(model.House{}, expectedError)
	_, err := services.Create(house)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, expectedError))
	mockRepo.AssertExpectations(t)
}
