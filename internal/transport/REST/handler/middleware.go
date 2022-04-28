package restHandler

import (
	"github.com/gin-gonic/gin"
)

const (
	access      = "Access-Control-Allow-Origin"
	contentType = "Content-Type"
)

func (h *Handler) setHeaders(c *gin.Context) {
	c.Header(access, "*")
	c.Header(contentType, "application/json; charset=utf-8")
	c.Next()
}
