package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message   string `json:"message"`
	RequestId string `json:"request_id"`
	Code      int    `json:"code"`
}

func newErrorResponse(c *gin.Context, err error, response errorResponse) {
	logrus.Error(err)
	response.RequestId = generateRequestID()
	c.AbortWithStatusJSON(response.Code, response)
}
func generateRequestID() string {
	id, err := uuid.NewRandom()
	if err != nil {
		return "unknown"
	}
	return id.String()
}
