package model

type Flat struct {
	Id      int    `json:"id"`
	HouseId int    `json:"house_id" binding:"required"`
	Price   int    `json:"price" binding:"required"`
	Rooms   int    `json:"rooms" binding:"required"`
	Status  string `json:"status"`
}
