package i18n

import "github.com/nicksnyder/go-i18n/v2/i18n"

type (
	KeyMessage struct {
		_              struct{}
		I18nKey        string
		DefaultMessage string
	}
	MapStringKeyMessage map[string]KeyMessage
	MapIntKeyMessage    map[int]KeyMessage
	LocalizeConfig      i18n.LocalizeConfig
)

func (k *KeyMessage) GetMessage() string {
	if k.I18nKey == "" {
		return k.DefaultMessage
	}

	return GinI18n.MustGetMessage(k.I18nKey)
}
