package gini18n

import "github.com/gin-gonic/gin"

type (
	GetLngHandler = func(context *gin.Context, defaultLng string) string
	Option        func(GinI18n)
)
