package routes

import (
	"time"

	"github.com/gin-gonic/gin"
)

type healthCheck struct {
	Timestamptz time.Time `json:"timestamptz"`
}

// Healthz godoc
// @Summary health check endpoint
// @Description health check endpoint
// @Accept  json
// @Produce  json
// @Success 200 {object} routes.healthCheck
// @Router /healthz [get]
// @Tags health
func Healthz(c *gin.Context) {
	c.JSON(200, gin.H{
		"timestamptz": time.Now().Format(time.RFC3339),
	})
}
