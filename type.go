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

	GinI18nOption func(n I18n)
)
