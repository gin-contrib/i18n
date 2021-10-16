package main

import (
	gini18n "gin-i18n"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	// new gin engine
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// apply i18n middleware
	router.Use(gini18n.Localize())

	// Router Index
	index := router.Group("/")
	{
		index.GET("/", func(context *gin.Context) {
			context.String(http.StatusOK, gini18n.MustGetMessage("welcome"))
		})
	}

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
