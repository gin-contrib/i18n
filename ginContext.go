package i18n

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func getLngFromGinContext(context *gin.Context) string {
	defaultLng := DefaultLanguage.String()
	if context == nil || context.Request == nil {
		return defaultLng
	}

	lng := context.GetHeader("Accept-Language")
	if lng == "" {
		lng = context.Query("lng")
		if lng == "" {
			return defaultLng
		}
	}

	// lng format may be like this en-US show just get en
	lngSplit := strings.Split(lng, "-")
	if len(lngSplit) > 0 {
		lng = lngSplit[0]
	}

	return lng
}

func Localize() gin.HandlerFunc {
	return func(context *gin.Context) {
		GinI18n.setCurrentContext(context)
	}
}
