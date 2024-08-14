package model

import "github.com/google/uuid"

type Flat struct {
	Id          int       `json:"id"`
	HouseId     int       `json:"house_id" binding:"required"`
	Price       int       `json:"price" binding:"required"`
	Rooms       int       `json:"rooms" binding:"required"`
	ModeratorId uuid.UUID `json:"-"`
	Status      string    `json:"status"`
}
