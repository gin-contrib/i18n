package i18n

import (
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

const (
	defaultFormatFile = "yaml"
)

var (
	defaultLanguage      = language.English
	defaultUnmarshalFunc = yaml.Unmarshal

	acceptLanguage = map[language.Tag]bool{
		defaultLanguage: true,
		language.German: true,
		language.French: true,
	}
)
