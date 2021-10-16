package main

import (
	gini18n "github.com/bos-hieu/gin-i18n"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"log"
	"net/http"
)

func main() {
	// new gin engine
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// apply i18n middleware
	router.Use(gini18n.Localize())

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, gini18n.MustGetMessage("welcome"))
	})

	router.GET("/:name", func(context *gin.Context) {
		context.String(http.StatusOK, gini18n.MustGetMessage(&i18n.LocalizeConfig{
			MessageID: "welcomeWithName",
			TemplateData: map[string]string{
				"name": context.Param("name"),
			},
		}))
	})

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
