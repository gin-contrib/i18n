package gini18n

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type BundleCfg struct {
	DefaultLanguage  language.Tag
	FormatBundleFile string
	AcceptLanguage   []language.Tag
	RootPath         string
	UnmarshalFunc    i18n.UnmarshalFunc
}

func WithBundle(config *BundleCfg) Option {
	return func(g GinI18n) {
		g.setBundle(config)
	}
}

func WithGetLngHandle(handler GetLngHandler) Option{
	return func(g GinI18n) {
		g.setGetLngHandler(handler)
	}
}
