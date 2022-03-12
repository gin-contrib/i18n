package i18n

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// BundleCfg ...
type BundleCfg struct {
	DefaultLanguage  language.Tag
	FormatBundleFile string
	AcceptLanguage   []language.Tag
	RootPath         string
	UnmarshalFunc    i18n.UnmarshalFunc
	Loader           Loader
}

type Loader interface {
	LoadMessage(path string) ([]byte, error)
}

type LoaderFunc func(path string) ([]byte, error)

func (f LoaderFunc) LoadMessage(path string) ([]byte, error) { return f(path) }

// WithBundle ...
func WithBundle(config *BundleCfg) Option {
	return func(g GinI18n) {
		if config.Loader == nil {
			config.Loader = defaultLoader
		}
		g.setBundle(config)
	}
}

// WithGetLngHandle ...
func WithGetLngHandle(handler GetLngHandler) Option {
	return func(g GinI18n) {
		g.setGetLngHandler(handler)
	}
}
