package i18n

import "github.com/gin-gonic/gin"

type I18n interface {
	GetMessage(param interface{}) (string, error)
	MustGetMessage(param interface{}) string
	SetCurrentGinContext(ctx *gin.Context)
}
