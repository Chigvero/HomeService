package transport

import (
	"HomeService/model"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (h *Handler) flatCreate(c *gin.Context) {
	var flat model.Flat
	err := c.BindJSON(&flat)
	if err != nil {
		newErrorResponse(c, err, errorResponse{"Invalid data type", "request_id", 400})
		return
	}
	flat, err = h.service.Flat.Create(flat)
	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"flats_pkey\"" {
			newErrorResponse(c, err, errorResponse{"Flat with this id was created earlier", "request_id", 500})
			return
		}
		newErrorResponse(c, err, errorResponse{"Что то пошло не так", "request_id", 500})
		return
	}
	c.JSON(200, flat)
}

type flatUpdate struct {
	Id      int    `json:"id" binding:"required"`
	HouseID int    `json:"house_id" binding:"required"`
	Status  string `json:"status" binding:"required"`
}

func (h *Handler) flatUpdate(c *gin.Context) {
	usr, ok := c.Get("user_type")
	usrId, ok := c.Get("user_id")
	if !ok {
		logrus.Print("user_type not found")
		return
	}
	if !ok {
		logrus.Print("user_id not found")
		return
	}

	user_id := usrId.(uuid.UUID)
	user_type := usr.(string)

	if user_type != "moderator" {
		err := errors.New("Insufficient access rights ")
		newErrorResponse(c, err, errorResponse{err.Error(), "Request_id", 401})
		return
	}
	var flatUp flatUpdate
	err := c.BindJSON(&flatUp)
	if err != nil {
		newErrorResponse(c, err, errorResponse{"Invalid data type", "request_id", 400})
		return
	}
	flat, err := h.service.Flat.GetById(flatUp.Id, flatUp.HouseID)
	if err != nil {
		newErrorResponse(c, err, errorResponse{"Flat with this id not found", "request_id", 500})
		return
	}
	if flat.Status == "on moderation" && flat.ModeratorId != user_id {
		err = errors.New("error: Flat is already in moderation or has been processed")
		newErrorResponse(c, err, errorResponse{"Flat is already in moderation or has been processed", "request_id", 400})
		return
	}
	flat, err = h.service.Flat.Update(flatUp.Id, flatUp.HouseID, flatUp.Status, user_id)
	if err != nil {
		newErrorResponse(c, err, errorResponse{"Что-то пошло не так", "request_id", 500})
		return
	}
	c.JSON(200, flat)
}
