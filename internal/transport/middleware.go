package transport

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
)

func (h *Handler) authMiddleware(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		err := errors.New("empty auth header")
		newErrorResponse(c, err, errorResponse{err.Error(), "request_id", 401})
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		err := errors.New("invalid auth header")
		newErrorResponse(c, err, errorResponse{err.Error(), "request_id", 401})
		return
	}
	userLogin, err := h.service.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, err, errorResponse{err.Error(), "request_id", 401})
		return
	}
	c.Set("user_type", userLogin.UserType)
	c.Set("user_id", userLogin.Id)
}
