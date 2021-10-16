# i18n
[![Run Tests](https://github.com/bos-hieu/i18n/actions/workflows/go.yml/badge.svg)](https://github.com/bos-hieu/i18n/actions/workflows/go.yml)
[![CodeQL](https://github.com/bos-hieu/i18n/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/bos-hieu/i18n/actions/workflows/codeql-analysis.yml)


## Usage

Canonical example:
```go
package main

import (
	gini18n "gin-i18n"
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
```

Customized Bundle
```go
package main

import (
	"encoding/json"
	gini18n "gin-i18n"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"log"
	"net/http"
)

func main() {
	// new gin engine
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// apply i18n middleware
	router.Use(gini18n.Localize(gini18n.WithBundle(&gini18n.BundleCfg{
		RootPath:         "./example/localizeJSON",
		AcceptLanguage:   []language.Tag{language.German, language.English},
		DefaultLanguage:  language.English,
		UnmarshalFunc:    json.Unmarshal,
		FormatBundleFile: "json",
	})))

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
```

Customized Get Language Handler
```go
package main

import (
	gini18n "gin-i18n"
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
	router.Use(gini18n.Localize(
		gini18n.WithGetLngHandle(
			func(context *gin.Context, defaultLng string) string {
				lng := context.Query("lng")
				if lng == "" {
					return defaultLng
				}
				return lng
			},
		),
	))

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

```