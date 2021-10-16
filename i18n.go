package gini18n

import (
	"github.com/gin-gonic/gin"
)

var atI18n GinI18n

func newI18n(opts ...Option) {
	// init default value
	ins := &ginI18nImpl{
		getLngHandler: defaultGetLngHandler,
	}
	ins.setBundle(defaultBundleConfig)

	// overwrite default value by options
	for _, opt := range opts {
		opt(ins)
	}

	atI18n = ins
}

func Localize(opts ...Option) gin.HandlerFunc {
	newI18n(opts...)
	return func(context *gin.Context) {
		atI18n.setCurrentContext(context)
	}
}

/*GetMessage get the i18n message
 param is one of these type: messageID, *i18n.LocalizeConfig
 Example:
	GetMessage("hello") // messageID is hello
	GetMessage(&i18n.LocalizeConfig{
			MessageID: "welcomeWithName",
			TemplateData: map[string]string{
				"name": context.Param("name"),
			},
	})
 */
func GetMessage(param interface{}) (string, error) {
	return atI18n.getMessage(param)
}

/*MustGetMessage get the i18n message without error handling
  param is one of these type: messageID, *i18n.LocalizeConfig
  Example:
	MustGetMessage("hello") // messageID is hello
	MustGetMessage(&i18n.LocalizeConfig{
			MessageID: "welcomeWithName",
			TemplateData: map[string]string{
				"name": context.Param("name"),
			},
	})
 */
func MustGetMessage(param interface{}) string {
	return atI18n.mustGetMessage(param)
}
