package transport

import (
	"Avito/model"
	"github.com/gin-gonic/gin"
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

}
