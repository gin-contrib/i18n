package i18n

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

// newI18n ...
func newI18n(opts ...Option) GinI18n {
	// init ins
	ins := &ginI18nImpl{}

	// set ins property from opts
	for _, opt := range opts {
		opt(ins)
	}

	// 	if bundle isn't constructed then assign it from default
	if ins.bundle == nil {
		ins.SetBundle(defaultBundleConfig)
	}

	// if getLngHandler isn't constructed then assign it from default
	if ins.getLngHandler == nil {
		ins.getLngHandler = defaultGetLngHandler
	}

	return ins
}

// Localize ...
func Localize(opts ...Option) gin.HandlerFunc {
	atI18n := newI18n(opts...)
	return func(context *gin.Context) {
		context.Set("i18n", atI18n)
	}
}

// GetMessage get the i18n message with error handling
// param is one of these type: messageID, *i18n.LocalizeConfig
// Example:
// GetMessage(context, "hello") // messageID is hello
//
//	GetMessage(context, &i18n.LocalizeConfig{
//	  MessageID: "welcomeWithName",
//	  TemplateData: map[string]string{
//	    "name": context.Param("name"),
//	  },
//	})
func GetMessage(context *gin.Context, param interface{}) (string, error) {
	atI18n := context.Value("i18n").(GinI18n)
	return atI18n.GetMessage(context, param)
}

// MustGetMessage get the i18n message without error handling
// param is one of these type: messageID, *i18n.LocalizeConfig
// Example:
// MustGetMessage(context, "hello") // messageID is hello
//
//	MustGetMessage(context, &i18n.LocalizeConfig{
//	  MessageID: "welcomeWithName",
//	  TemplateData: map[string]string{
//	    "name": context.Param("name"),
//	  },
//	})
func MustGetMessage(context *gin.Context, param interface{}) string {
	atI18n := context.MustGet("i18n").(GinI18n)
	return atI18n.MustGetMessage(context, param)
}

// HasLang check all i18n lang exists
// Example:
// HasLang(context, "ZH-cn") // return false or true
func HasLang(context *gin.Context, language string) bool {
	atI18n := context.MustGet("i18n").(GinI18n)
	return atI18n.hasLang(language)
}

// GetDefaultLanguage get the default language
// Example:
// GetDefaultLanguage(context)
func GetDefaultLanguage(context *gin.Context) language.Tag {
	atI18n := context.MustGet("i18n").(GinI18n)
	return atI18n.getDefaultLanguage()
}

// GetCurrentLanguage get the current language
// Example:
// GetCurrentLanguage(context)
func GetCurrentLanguage(context *gin.Context) language.Tag {
	atI18n := context.MustGet("i18n").(GinI18n)
	return atI18n.getCurrentLanguage(context)
}

// I18n get GinI18n from gin.Context
func I18n(context *gin.Context) GinI18n {
	return context.MustGet("i18n").(GinI18n)
}
