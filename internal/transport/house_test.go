package transport

import (
	"HomeService/internal/service"
	"HomeService/internal/service/mocks"
	"HomeService/model"
	"bytes"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_getHouseFlatsList_InvalidID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHouse := mock_service.NewMockHouse(ctrl)
	handler := &Handler{
		service: &service.Service{
			House: mockHouse,
		},
	}

	router := gin.Default()
	router.GET("/house/:id", handler.getHouseFlatsList)

	userType := "moderator"

	ctx := context.WithValue(context.Background(), "user_type", userType)
	req, _ := http.NewRequestWithContext(ctx, "GET", "/house/invalid", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestHandler_createHouse_Client(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHouse := mock_service.NewMockHouse(ctrl)
	handler := &Handler{
		service: &service.Service{
			House: mockHouse,
		},
	}

	router := gin.Default()
	router.POST("/house/create", handler.createHouse)

	userType := "client"
	house := model.House{
		Id:        1,
		Address:   "123 Street",
		Year:      2023,
		Developer: "PIK",
	}

	ctx := context.WithValue(context.Background(), "user_type", userType)
	body, _ := json.Marshal(house)
	req, _ := http.NewRequestWithContext(ctx, "POST", "/house/create", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestHandler_createHouse_InvalidData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHouse := mock_service.NewMockHouse(ctrl)
	handler := &Handler{
		service: &service.Service{
			House: mockHouse,
		},
	}

	router := gin.Default()
	router.POST("/house/create", handler.createHouse)

	userType := "moderator"
	invalidHouse := model.House{
		Id:        0,
		Address:   "",
		Year:      0,
		Developer: "",
	}

	ctx := context.WithValue(context.Background(), "user_type", userType)
	body, _ := json.Marshal(invalidHouse)
	req, _ := http.NewRequestWithContext(ctx, "POST", "/house/create", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
}
