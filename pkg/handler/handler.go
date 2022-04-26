package handler

import (
	"github.com/gin-gonic/gin"
	"gitlab.digital-spirit.ru/study/artem_crud/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(h.setHeaders)

	router.GET("/", h.getRecordsByFilter)
	router.GET("/:uid", h.getRecordById)
	router.POST("/", h.createRecord)
	router.PUT("/:uid", h.updateRecord)
	router.DELETE("/:uid", h.deleteRecord)

	return router
}
