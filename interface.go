package i18n

import (
	"github.com/gin-gonic/gin"
)

// GinI18n ...
type GinI18n interface {
	GetMessage(context *gin.Context, param interface{}) (string, error)
	MustGetMessage(context *gin.Context, param interface{}) string
	SetBundle(cfg *BundleCfg)
	SetGetLngHandler(handler GetLngHandler)
}
