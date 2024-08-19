package service

import (
	"HomeService/internal/repository"
	"HomeService/model"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

type AuthService struct {
	repos repository.Authorization
}

const (
	key  = "myKey"
	salt = "myHASH"
)

type CustomClaim struct {
	jwt.RegisteredClaims
	UserType string    `json:"role"`
	UserId   uuid.UUID `json:"user_id"`
}

func NewAuthService(authorization repository.Authorization) *AuthService {
	return &AuthService{
		authorization,
	}
}

func (s *AuthService) Register(user model.UserRegister) (string, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repos.Register(user)
}

func (s *AuthService) Login(user model.UserLogin) (string, error) {
	user.Password = generatePasswordHash(user.Password)
	user_type, err := s.repos.Login(user)
	if err != nil {
		return "", err
	}
	token, err := s.generateToken(user_type, user.Id)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *AuthService) generateToken(user_type string, userId uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &CustomClaim{
		UserType: user_type,
		UserId:   userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})
	return token.SignedString([]byte(key))
}

func (s *AuthService) ParseToken(tokenString string) (model.UserLogin, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if err != nil {
		return model.UserLogin{}, err
	}
	claims, ok := token.Claims.(*CustomClaim)
	if !ok {
		return model.UserLogin{}, errors.New("token claims are not of type *tokenClaims")
	}
	return model.UserLogin{
		claims.UserId,
		"",
		claims.UserType,
	}, nil
}

func (s *AuthService) DummyLogin(user_type string, id uuid.UUID) (string, error) {
	return s.generateToken(user_type, id)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
