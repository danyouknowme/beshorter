package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/danyouknowme/beshorter/config"
	"github.com/danyouknowme/beshorter/httpserver"
	"github.com/danyouknowme/beshorter/service"
	"github.com/gin-gonic/gin"
)

type ShortUrlHandler struct {
	shortUrlService service.ShortUrlService
}

func NewShortUrlHandler(shortUrlService service.ShortUrlService) ShortUrlHandler {
	return ShortUrlHandler{
		shortUrlService: shortUrlService,
	}
}

type CreateShortenerUrlHandlerRequest struct {
	Url string `json:"url"`
}

func (h *ShortUrlHandler) CreateShortenerUrlHandler(c *gin.Context) {
	var req CreateShortenerUrlHandlerRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, httpserver.ErrorResponse(err))
		return
	}

	url, err := h.shortUrlService.CreateShortenerUrl(req.Url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpserver.ErrorResponse(err))
		return
	}

	urlResponse := fmt.Sprintf("%s/%s", config.ServerUrl, url)
	c.JSON(http.StatusOK, gin.H{"url": urlResponse})
}

func (h *ShortUrlHandler) GetShortenerUrlHandler(c *gin.Context) {
	shortUrl := c.Param("url")
	fullUrl, err := h.shortUrlService.GetShortenerUrl(shortUrl)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, httpserver.ErrorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, httpserver.ErrorResponse(err))
		return
	}

	c.Redirect(http.StatusMovedPermanently, fullUrl)
}
