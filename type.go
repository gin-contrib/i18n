package gini18n

import "github.com/gin-gonic/gin"

type (
	// GetLngHandler ...
	GetLngHandler = func(context *gin.Context, defaultLng string) string

	// Option ...
	Option func(GinI18n)
)
