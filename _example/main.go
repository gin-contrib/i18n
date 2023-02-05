package main

import (
	"log"
	"net/http"

	ginI18n "github.com/bizflycloud/i18n"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func main() {
	// new gin engine
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// apply i18n middleware
	router.Use(ginI18n.Localize())

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, ginI18n.MustGetMessage(context, "welcome"))
	})

	router.GET("/:name", func(context *gin.Context) {
		context.String(http.StatusOK, ginI18n.MustGetMessage(context, &i18n.LocalizeConfig{
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
