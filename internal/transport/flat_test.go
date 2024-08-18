package transport

import (
	"HomeService/internal/service"
	"HomeService/internal/service/mocks"
	"HomeService/model"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_flatCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFlat := mock_service.NewMockFlat(ctrl)
	handler := &Handler{
		service: &service.Service{
			Flat: mockFlat,
		},
	}

	router := gin.Default()
	router.POST("/flatCreate", handler.flatCreate)

	flat := model.Flat{
		Id:      1,
		HouseId: 1,
		Price:   1000,
		Rooms:   3,
		Status:  "available",
	}

	mockFlat.EXPECT().Create(flat).Return(flat, nil)

	body, _ := json.Marshal(flat)
	req, _ := http.NewRequest("POST", "/flatCreate", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response model.Flat
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, flat, response)
}

func TestHandler_flatCreate_InvalidData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFlat := mock_service.NewMockFlat(ctrl)
	handler := &Handler{
		service: &service.Service{
			Flat: mockFlat,
		},
	}

	router := gin.Default()
	router.POST("/flatCreate", handler.flatCreate)

	invalidFlat := model.Flat{
		Id:      0,
		HouseId: 0,
		Price:   0,
		Rooms:   0,
		Status:  "",
	}

	body, _ := json.Marshal(invalidFlat)
	req, _ := http.NewRequest("POST", "/flat/create", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}

func TestHandler_flatCreate_DuplicateKey(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFlat := mock_service.NewMockFlat(ctrl)
	handler := &Handler{
		service: &service.Service{
			Flat: mockFlat,
		},
	}

	router := gin.Default()
	router.POST("/flat/create", handler.flatCreate)

	flat := model.Flat{
		Id:      1,
		HouseId: 1,
		Price:   1000,
		Rooms:   3,
		Status:  "available",
	}

	mockFlat.EXPECT().Create(flat).Return(model.Flat{}, errors.New("pq: duplicate key value violates unique constraint \"flats_pkey\""))

	body, _ := json.Marshal(flat)
	req, _ := http.NewRequest("POST", "/flat/create", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
