package i18n

import (
	"github.com/gin-gonic/gin"
)

// GinI18n ...
type GinI18n interface {
	getMessage(context *gin.Context, param interface{}) (string, error)
	mustGetMessage(context *gin.Context, param interface{}) string
	setBundle(cfg *BundleCfg)
	setGetLngHandler(handler GetLngHandler)
	HasLang(language string) bool
}
