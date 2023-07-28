package main

import (
	"embed"
	"encoding/json"
	"log"
	"net/http"

	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

//go:embed i18n/localizeJSON/*
var fs embed.FS

func main() {
	// new gin engine
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// apply i18n middleware
	router.Use(ginI18n.Localize(ginI18n.WithBundle(&ginI18n.BundleCfg{
		DefaultLanguage:  language.English,
		FormatBundleFile: "json",
		AcceptLanguage:   []language.Tag{language.English, language.German, language.Chinese},
		RootPath:         "./i18n/localizeJSON/",
		UnmarshalFunc:    json.Unmarshal,
		// After commenting this line, use defaultLoader
		// it will be loaded from the file
		Loader: &ginI18n.EmbedLoader{
			FS: fs,
		},
	})))

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, ginI18n.MustGetMessage(ctx, "welcome"))
	})

	router.GET("/:name", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, ginI18n.MustGetMessage(
			ctx,
			&i18n.LocalizeConfig{
				MessageID: "welcomeWithName",
				TemplateData: map[string]string{
					"name": ctx.Param("name"),
				},
			}))
	})

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
