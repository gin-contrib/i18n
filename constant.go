package i18n

import (
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

const (
	DefaultFormatBundleFile = "yaml"
)

var (
	DefaultLanguage      = language.English
	DefaultUnmarshalFunc = yaml.Unmarshal

	DefaultAcceptLanguage = []language.Tag{
		DefaultLanguage,
		language.German,
		language.French,
	}
)
