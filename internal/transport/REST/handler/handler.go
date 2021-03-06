package restHandler

import (
	"github.com/gin-gonic/gin"
	"gitlab.digital-spirit.ru/study/artem_crud/internal/service"
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
	router.GET("/:uuid", h.getRecordByUuid)
	router.POST("/", h.createRecord)
	router.PUT("/:uuid", h.updateRecord)
	router.DELETE("/:uuid", h.deleteRecord)

	return router
}
