package i18n

import "github.com/gin-gonic/gin"

// defaultGetLngHandler ...
func defaultGetLngHandler(context *gin.Context, defaultLng string) string {
	if context == nil || context.Request == nil {
		return defaultLng
	}

	lng := context.GetHeader("Accept-Language")
	if lng != "" {
		return lng
	}

	lng = context.Query("lng")
	if lng == "" {
		return defaultLng
	}

	return lng
}
