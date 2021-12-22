package main

import (
	"encoding/json"
	"log"
	"net/http"

	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func main() {
	// new gin engine
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// apply i18n middleware
	r.Use(ginI18n.Localize(ginI18n.WithBundle(&ginI18n.BundleCfg{
		RootPath:         "./i18n",
		AcceptLanguage:   []language.Tag{language.Chinese, language.English},
		DefaultLanguage:  language.English,
		FormatBundleFile: "json",
		UnmarshalFunc:    json.Unmarshal,
	})))

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, ginI18n.MustGetMessage("welcome"))
	})

	router.GET("/:name", func(context *gin.Context) {
		context.String(http.StatusOK, ginI18n.MustGetMessage(&i18n.LocalizeConfig{
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
