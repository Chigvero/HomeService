package transport

import (
	"Avito/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		service: services,
	}
}

// +-+FlatUpdate исправить так как у квартир нет уникального id-+-
// +-+Изменить саму структуру модели скл-+-
// +-+Надо изменить  GetFlatById и Update в постгрес-+-
// ++Надо написать ручки dummyLogin++

// Надо написать ручки и subscribe
// Добавить рефреш токен
// Надо разработать модульные и интеграционные тесты

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.POST("/register", h.register)
	router.POST("/login", h.login)
	router.GET("/dummyLogin", h.dummyLogin)
	authorized := router.Group("/")
	authorized.Use(h.authMiddleware)
	{
		authorized.GET("/house/:id", h.getHouseFlatsList)
		authorized.POST("/house/create", h.createHouse)
		authorized.POST("/house/:id/subscribe", h.subscribeHouse)
		authorized.POST("/flat/create", h.flatCreate)
		authorized.POST("/flat/update", h.flatUpdate)
	}
	return router
}
