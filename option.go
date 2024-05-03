package i18n

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// BundleCfg represents the configuration options for an i18n bundle.
type BundleCfg struct {
	DefaultLanguage  language.Tag       // DefaultLanguage specifies the default language for the bundle.
	FormatBundleFile string             // FormatBundleFile specifies the file format for the bundle.
	AcceptLanguage   []language.Tag     // AcceptLanguage specifies the accepted languages for the bundle.
	RootPath         string             // RootPath specifies the root path for the bundle.
	UnmarshalFunc    i18n.UnmarshalFunc // UnmarshalFunc specifies the function used for unmarshaling bundle files.
	Loader           Loader             // Loader specifies the loader for loading bundle files.
}

type Loader interface {
	LoadMessage(path string) ([]byte, error)
}

type LoaderFunc func(path string) ([]byte, error)

func (f LoaderFunc) LoadMessage(path string) ([]byte, error) { return f(path) }

// WithBundle returns an Option that sets the bundle configuration for GinI18n.
// If the loader is not provided in the BundleCfg, the defaultLoader will be used.
func WithBundle(config *BundleCfg) Option {
	return func(g GinI18n) {
		if config.Loader == nil {
			config.Loader = defaultLoader
		}
		g.setBundle(config)
	}
}

// WithGetLngHandle sets the handler function for retrieving the current language.
// The provided handler function should accept a GinI18n instance and return the current language as a string.
// This option allows you to customize how the current language is determined.
func WithGetLngHandle(handler GetLngHandler) Option {
	return func(g GinI18n) {
		g.setGetLngHandler(handler)
	}
}
