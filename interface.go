package i18n

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
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
	
	// HasLang checks if the given language is supported by the i18n bundle.
	// It returns true if the language is supported, false otherwise.
	HasLang(language string) bool

	// GetDefaultLanguage returns the default language tag.
	// It returns the default language tag.
	GetDefaultLanguage() language.Tag

	// GetCurrentLanguage returns the current language tag from the context.
	// It returns the current language tag.
	GetCurrentLanguage(context *gin.Context) language.Tag
}
