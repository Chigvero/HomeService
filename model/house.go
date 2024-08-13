package model

type House struct {
	Id        int    `json:"id"`
	Address   string `json:"address" binding:"required"`
	Year      int    `json:"year" binding:"required"`
	Developer string `json:"developer" binding:"required"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"update_at"`
}
