package transport

import (
	"HomeService/model"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func (h *Handler) register(c *gin.Context) {
	var user model.UserRegister
	err := c.BindJSON(&user)
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatusJSON(400, map[string]interface{}{"error": "Invalid data type"}) //?????
		return
	}
	id, err := h.service.Authorization.Register(user)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			logrus.Error(err)
			c.AbortWithStatusJSON(409, map[string]interface{}{"error": "User with this email already exists"})
			return
		}
		resp := errorResponse{
			"Что-то пошло не так",
			"request_id", //Надо разобраться
			500,
		}
		newErrorResponse(c, err, resp)
		return
	}
	c.JSON(200, map[string]interface{}{
		"user_id": id,
	})
}

func (h *Handler) login(c *gin.Context) {
	var user model.UserLogin
	err := c.BindJSON(&user)
	if err != nil {
		logrus.Println(err)
		c.AbortWithStatusJSON(400, map[string]interface{}{
			"error": "Invalid data type",
		})
		return
	}
	token, err := h.service.Authorization.Login(user)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.AbortWithStatusJSON(404, map[string]interface{}{
				"error": "User not found",
			})
			return
		}
		resp := errorResponse{
			"Что-то пошло не так",
			"request_id", //Надо разобраться
			500,
		}
		newErrorResponse(c, err, resp)
		return
	}
	c.JSON(200, token)
}

func (h *Handler) dummyLogin(c *gin.Context) {
	userType := c.Query("user_type")
	if userType != "moderator" && userType != "client" {
		err := errors.New("Incorrect user_type")
		newErrorResponse(c, err, errorResponse{err.Error(), "request_id", 400})
		return
	}
	user_id := uuid.New()
	token, err := h.service.Authorization.DummyLogin(userType, user_id)
	if err != nil {
		newErrorResponse(c, err, errorResponse{"что то пошло не так", "request_id", 500})
	}
	c.JSON(200, token)
}
