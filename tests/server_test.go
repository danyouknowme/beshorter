package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/danyouknowme/beshorter/config"
	"github.com/danyouknowme/beshorter/database"
	"github.com/danyouknowme/beshorter/handler"
	"github.com/danyouknowme/beshorter/helper"
	"github.com/danyouknowme/beshorter/repository"
	"github.com/danyouknowme/beshorter/service"
	"github.com/gin-gonic/gin"
)

var shortUrlHandler handler.ShortUrlHandler

func init() {
	config, _ := config.LoadConfig("../config")

	db, _ := database.NewDatabaseConnection(database.DatabaseConfig{
		Driver:       config.DB_Driver,
		Hostname:     config.DB_Hostname,
		Port:         config.DB_Port,
		Username:     config.DB_Username,
		Password:     config.DB_Password,
		DatabaseName: config.DB_Name,
	})

	shortUrlRepository := repository.NewShortUrlRepository(db)
	shortUrlService := service.NewShortUrlService(shortUrlRepository)
	shortUrlHandler = handler.NewShortUrlHandler(shortUrlService)
}

func TestHealthCheck(t *testing.T) {
	gin.SetMode(gin.TestMode)
	healthcheckPath := "/healthcheck"
	expectedResponse := "\"URL Shortener...\""

	res := httptest.NewRecorder()
	c, r := gin.CreateTestContext(res)
	c.Request = httptest.NewRequest(http.MethodGet, healthcheckPath, nil)
	c.Request.Header.Set("Content-Type", "application/json")

	r.GET(healthcheckPath, handler.HealthCheckHandler)
	r.ServeHTTP(res, c.Request)
	status := res.Code
	body, _ := io.ReadAll(res.Body)
	if status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	require.Equal(t, http.StatusOK, status)
	require.NotEmpty(t, body)
	require.Equal(t, expectedResponse, string(body))
}

func TestCreateUrlShortener(t *testing.T) {
	gin.SetMode(gin.TestMode)
	request := handler.CreateShortenerUrlHandlerRequest{
		Url: "https://www.google.com/search?q=link65&tbm=isch&ved=2ahUKEwiBnJ2ukaP7AhWJyqACHcTdDlgQ2-cCegQIABAA&oq=link65&gs_lcp=CgNpbWcQAzIHCAAQgAQQGDIHCAAQgAQQGDIHCAAQgAQQGDIHCAAQgAQQGDIHCAAQgAQQGDIHCAAQgAQQGDIHCAAQgAQQGDIHCAAQgAQQGDIHCAAQgAQQGDIHCAAQgAQQGDoECCMQJzoICAAQgAQQsQM6BQgAEIAEOgUIABCxAzoICAAQsQMQgwFQ2AVY3w5ghxBoAHAAeACAAfkBiAHSBpIBBTUuMC4ymAEAoAEBqgELZ3dzLXdpei1pbWfAAQE&sclient=img&ei=Pq5sY4GFN4mVg8UPxLu7wAU&bih=789&biw=1440&rlz=1C5CHFA_enTH991TH991#imgrc=g3ar9oBbSpfjTM",
	}
	jsonStr, _ := json.Marshal(request)

	res := httptest.NewRecorder()
	c, r := gin.CreateTestContext(res)
	c.Request = httptest.NewRequest(http.MethodPost, "/url/shorter", bytes.NewBuffer(jsonStr))
	c.Request.Header.Set("Content-Type", "application/json")

	r.POST("/url/shorter", shortUrlHandler.CreateShortenerUrlHandler)
	r.ServeHTTP(res, c.Request)
	status := res.Code
	body, _ := io.ReadAll(res.Body)
	if status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	require.Equal(t, http.StatusOK, status)
	require.NotEmpty(t, body)
}

func TestGetUrlShortenerNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	randomUrl, _ := helper.GenerateShorterUrl()
	getUrlShortenerPath := "/" + randomUrl

	res := httptest.NewRecorder()
	c, r := gin.CreateTestContext(res)
	c.Request = httptest.NewRequest(http.MethodGet, getUrlShortenerPath, nil)
	c.Request.Header.Set("Content-Type", "application/json")

	r.GET(getUrlShortenerPath, shortUrlHandler.GetShortenerUrlHandler)
	r.ServeHTTP(res, c.Request)
	status := res.Code
	body, _ := io.ReadAll(res.Body)
	if status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
	require.Equal(t, http.StatusNotFound, status)
	require.NotEmpty(t, body)
}
