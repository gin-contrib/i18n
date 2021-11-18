package i18n

import (
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

const (
	defaultFormatBundleFile = "yaml"
	defaultRootPath         = "./example/localize"
)

var (
	defaultLanguage       = language.English
	defaultUnmarshalFunc  = yaml.Unmarshal
	defaultAcceptLanguage = []language.Tag{
		defaultLanguage,
		language.German,
		language.French,
	}

	defaultBundleConfig = &BundleCfg{
		RootPath:         defaultRootPath,
		AcceptLanguage:   defaultAcceptLanguage,
		FormatBundleFile: defaultFormatBundleFile,
		DefaultLanguage:  defaultLanguage,
		UnmarshalFunc:    defaultUnmarshalFunc,
	}
)
