package model

import "github.com/google/uuid"

type UserRegister struct {
	Id       uuid.UUID `json:"id"`
	Email    string    `json:"email" binding:"required"`
	Password string    `json:"password" binding:"required"`
	UserType string    `json:"user_type" binding:"required"`
}

type UserLogin struct {
	Id       uuid.UUID `json:"id" binding:"required"`
	Password string    `json:"password" binding:"required"`
	UserType string    `json:"user_type"`
}
