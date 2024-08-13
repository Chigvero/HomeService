package service

import (
	"Avito/internal/repository"
	"Avito/model"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type AuthService struct {
	repos repository.Authorization
}

const (
	key = "myKey"
)

type CustomClaim struct {
	jwt.RegisteredClaims
	UserType string `json:"role"`
}

func NewAuthService(authorization repository.Authorization) *AuthService {
	return &AuthService{
		authorization,
	}
}

func (s *AuthService) Register(user model.UserRegister) (string, error) {
	return s.repos.Register(user)
}

func (s *AuthService) Login(user model.UserLogin) (string, error) {
	user_type, err := s.repos.Login(user)
	if err != nil {
		return "", err
	}
	token, err := s.generateToken(user_type)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *AuthService) generateToken(user_type string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &CustomClaim{
		UserType: user_type,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})
	return token.SignedString([]byte(key))
}

func (s *AuthService) ParseToken(tokenString string) (string, error) {
	//Почему ?&?CustomClaim
	//Возможно понадобится еще id
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(*CustomClaim)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UserType, nil
}

func (s *AuthService) DummyLogin() string {
	return ""
}
