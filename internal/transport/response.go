package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message   string `json:"message"`
	RequestId string `json:"request_id"`
	Code      int    `json:"code"`
}

func newErrorResponse(c *gin.Context, err error, response errorResponse) {
	logrus.Error(err)
	c.AbortWithStatusJSON(response.Code, response)
}
