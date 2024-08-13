package transport

import (
	"Avito/model"
	"errors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) flatCreate(c *gin.Context) {
	//usrt,_:=c.Get("user_type")
	//user_type:=usrt.(string)
	var flat model.Flat
	err := c.BindJSON(&flat)
	if err != nil {
		newErrorResponse(c, err, errorResponse{"Invalid data type", "request_id", 400})
		return
	}
	flat, err = h.service.Flat.Create(flat)
	if err != nil {
		newErrorResponse(c, err, errorResponse{"Что то пошло не так", "request_id", 500})
		return
	}
	c.JSON(200, flat)
}

type flatUpdate struct {
	Id     int    `json:"id" binding:"required"`
	Status string `json:"status" binding:"required"`
}

func (h *Handler) flatUpdate(c *gin.Context) {
	usr, _ := c.Get("user_type")
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
	flat, err := h.service.Flat.Update(flatUp.Id, flatUp.Status)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			newErrorResponse(c, err, errorResponse{"Flat with this id not found", "request_id", 500})
			return
		}
		newErrorResponse(c, err, errorResponse{"Что-то пошло не так", "request_id", 500})
		return
	}
	c.JSON(200, flat)
}
