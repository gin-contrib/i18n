package i18n

import (
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
	"io/ioutil"
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

	defaultLoader = LoaderFunc(func(path string) ([]byte, error) {
		return ioutil.ReadFile(path)
	})

	defaultBundleConfig = &BundleCfg{
		RootPath:         defaultRootPath,
		AcceptLanguage:   defaultAcceptLanguage,
		FormatBundleFile: defaultFormatBundleFile,
		DefaultLanguage:  defaultLanguage,
		UnmarshalFunc:    defaultUnmarshalFunc,
		Loader:           defaultLoader,
	}
)
