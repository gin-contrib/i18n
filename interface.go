package i18n

import (
	"github.com/gin-gonic/gin"
)

// GinI18n is an interface that defines methods for internationalization (i18n) in a Gin web framework context.
// It provides methods to get localized messages and configure the i18n bundle and language handler.
type GinI18n interface {
	// GetMessage retrieves a localized message based on the provided context and parameter.
	// It returns the localized message as a string and an error if the message could not be retrieved.
	GetMessage(context *gin.Context, param interface{}) (string, error)

	// MustGetMessage retrieves a localized message based on the provided context and parameter.
	// It returns the localized message as a string and panics if the message could not be retrieved.
	MustGetMessage(context *gin.Context, param interface{}) string

	// SetBundle sets the i18n bundle configuration.
	SetBundle(cfg *BundleCfg)

	// SetGetLngHandler sets the handler function to determine the language from the context.
	SetGetLngHandler(handler GetLngHandler)
}
