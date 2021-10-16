package gini18n

import "github.com/gin-gonic/gin"

func defaultGetLngHandler(context *gin.Context) string {
	defaultLng := defaultLanguage.String()
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

