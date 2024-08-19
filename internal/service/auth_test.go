package service

import (
	"HomeService/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockAuthRepo struct {
	mock.Mock
}

func (m *MockAuthRepo) Register(user model.UserRegister) (string, error) {
	args := m.Called(user)
	return args.Get(0).(string), args.Error(1)
}
func (m *MockAuthRepo) Login(user model.UserLogin) (string, error) {
	args := m.Called(user)
	return args.Get(0).(string), args.Error(1)
}

func TestAuthService_Register(t *testing.T) {
	mockRepo := new(MockAuthRepo)
	services := NewAuthService(mockRepo)

	user := model.UserRegister{
		Email:    "test@example.com",
		UserType: "user",
		Password: "password",
	}

	hashedPassword := generatePasswordHash(user.Password)
	userWithHashedPassword := user
	userWithHashedPassword.Password = hashedPassword

	mockRepo.On("Register", userWithHashedPassword).Return("123", nil)

	id, err := services.Register(user)
	assert.NoError(t, err)
	assert.Equal(t, "123", id)
	mockRepo.AssertExpectations(t)
}

func TestAuthService_ParseToken(t *testing.T) {
	services := NewAuthService(nil)

	token, err := services.generateToken("user", uuid.New())
	assert.NoError(t, err)

	user, err := services.ParseToken(token)
	assert.NoError(t, err)
	assert.Equal(t, "user", user.UserType)
}

func TestAuthService_DummyLogin(t *testing.T) {
	services := NewAuthService(nil)

	token, err := services.DummyLogin("user", uuid.New())
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}
