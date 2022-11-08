package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheckHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "URL Shortener...")
}
