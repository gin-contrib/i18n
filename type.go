package i18n

import "github.com/gin-gonic/gin"

type (
	// Option is a function type that takes a GinI18n instance and applies a configuration to it.
	// It is used to customize the behavior of the GinI18n middleware.
	GetLngHandler = func(context *gin.Context, defaultLng string) string

	// Option is a function type that takes a GinI18n instance and applies a configuration to it.
	Option func(GinI18n)
)
