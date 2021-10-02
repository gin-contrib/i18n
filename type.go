package i18n

import "github.com/gin-gonic/gin"

type (
	KeyMessage struct {
		_              struct{}
		I18nKey        string
		DefaultMessage string
	}
	MapStringKeyMessage map[string]KeyMessage
	MapIntKeyMessage    map[int]KeyMessage
)

func (k *KeyMessage) GetMessageFromGinContext(context *gin.Context) string {
	if k.I18nKey == "" {
		return k.DefaultMessage
	}

	return AutoI18n.GetMessage(k.I18nKey)
}
