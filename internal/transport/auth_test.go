package transport

import (
	"HomeService/internal/service"
	"HomeService/internal/service/mocks"
	"HomeService/model"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuth := mock_service.NewMockAuthorization(ctrl)
	handler := &Handler{
		service: &service.Service{
			Authorization: mockAuth,
		},
	}

	router := gin.Default()
	router.POST("/register", handler.register)

	user := model.UserRegister{
		Email:    "test@example.com",
		UserType: "user",
		Password: "password",
	}

	mockAuth.EXPECT().Register(user).Return("123", nil)

	body, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "123", response["user_id"])
}

func TestHandler_login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuth := mock_service.NewMockAuthorization(ctrl)
	handler := &Handler{
		service: &service.Service{
			Authorization: mockAuth,
		},
	}

	router := gin.Default()
	router.POST("/login", handler.login)

	user := model.UserLogin{
		Id:       uuid.New(),
		Password: "password",
	}

	mockAuth.EXPECT().Login(user).Return("token", nil)

	body, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response string
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "token", response)
}
