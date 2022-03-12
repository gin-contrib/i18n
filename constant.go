package i18n

import (
	"io/ioutil"

	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

const (
	defaultFormatBundleFile = "yaml"
	defaultRootPath         = "./_example/localize"
)

var (
	defaultLanguage       = language.English
	defaultUnmarshalFunc  = yaml.Unmarshal
	defaultAcceptLanguage = []language.Tag{
		defaultLanguage,
		language.German,
		language.French,
	}

	defaultLoader = LoaderFunc(ioutil.ReadFile)

	defaultBundleConfig = &BundleCfg{
		RootPath:         defaultRootPath,
		AcceptLanguage:   defaultAcceptLanguage,
		FormatBundleFile: defaultFormatBundleFile,
		DefaultLanguage:  defaultLanguage,
		UnmarshalFunc:    defaultUnmarshalFunc,
		Loader:           defaultLoader,
	}
)
