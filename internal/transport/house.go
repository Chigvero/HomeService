package transport

import (
	"HomeService/model"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

type flatsList struct {
	Flats []model.Flat `json:"flats"`
}

func (h *Handler) getHouseFlatsList(c *gin.Context) {
	usr, _ := c.Get("user_type")
	user_type, _ := usr.(string)
	var (
		flats flatsList
	)
	if user_type != "moderator" && user_type != "client" {
		err := errors.New("Insufficient access rights ")
		newErrorResponse(c, err, errorResponse{"Unauthorized access", "request_id", 401})
		return
	}
	id := c.Param("id")
	house_id, err := strconv.Atoi(id)
	if err != nil {
		newErrorResponse(c, err, errorResponse{"Invalid id data type", "request_id", 400})
		return
	}
	if user_type == "moderator" {
		flats.Flats, err = h.service.House.GetHouseModerFlatsList(house_id)
		if err != nil {
			newErrorResponse(c, err, errorResponse{"что-то пошло не так", "request_id", 500})
			return
		}
		c.JSON(200, flats)
		return
	}
	if user_type == "client" {
		flats.Flats, err = h.service.House.GetHouseClientFlatsList(house_id)
		if err != nil {
			newErrorResponse(c, err, errorResponse{"что-то пошло не так", "request_id", 500})
			return
		}
		c.JSON(200, flats)
		return
	}
}

func (h *Handler) createHouse(c *gin.Context) {
	usr, _ := c.Get("user_type")
	user_type, _ := usr.(string)
	if user_type != "moderator" {
		err := errors.New("Insufficient access rights ")
		newErrorResponse(c, err, errorResponse{err.Error(), "request_id", 401})
		return
	}
	var house model.House
	err := c.BindJSON(&house)
	if err != nil {
		newErrorResponse(c, err, errorResponse{"Invalid data type", "request_id", 400})
		return
	}
	house, err = h.service.House.Create(house)
	if err != nil {
		newErrorResponse(c, err, errorResponse{"Что то пошло не так", "request_id", 500})
		return
	}
	c.JSON(200, house)
}

func (h *Handler) subscribeHouse(c *gin.Context) {

}
