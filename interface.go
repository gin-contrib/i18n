package i18n

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

// GinI18n ...
type GinI18n interface {
	GetMessage(context *gin.Context, param interface{}) (string, error)
	MustGetMessage(context *gin.Context, param interface{}) string
	SetBundle(cfg *BundleCfg)
	SetGetLngHandler(handler GetLngHandler)
	hasLang(language string) bool
	getDefaultLanguage() language.Tag
	getCurrentLanguage(context *gin.Context) language.Tag
}
