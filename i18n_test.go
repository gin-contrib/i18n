package gini18n

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	testMessageID = "welcome"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func newServer() *gin.Engine {
	router := gin.New()
	router.Use(Localize())
	router.GET("/", func(c *gin.Context) {
		c.String(200, MustGetMessage(testMessageID))
	})
	return router
}

func TestI18nEn(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("Accept-Language", language.English.String())

	// Perform the request
	w := httptest.NewRecorder()
	r := newServer()
	r.ServeHTTP(w, req)


	log.Println(w.Body.String())
	assert.Equal(t, w.Code, 200)
	assert.Equal(t, w.Body.String(), "hello")
}

func TestI18nDE(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("Accept-Language", language.German.String())

	// Perform the request
	w := httptest.NewRecorder()
	r := newServer()
	r.ServeHTTP(w, req)


	log.Println(w.Body.String())
	assert.Equal(t, w.Code, 200)
	assert.Equal(t, w.Body.String(), "hello german")
}