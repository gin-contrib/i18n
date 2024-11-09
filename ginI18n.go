// Package i18n ginI18nImpl is an implementation of the GinI18n interface, providing
// localization support for Gin applications. It uses the go-i18n library
// to manage and retrieve localized messages.
//
// Fields:
// - bundle: The i18n.Bundle containing the localization messages.
// - localizerByLng: A map of language tags to their corresponding localizers.
// - defaultLanguage: The default language tag to use for localization.
// - getLngHandler: A handler function to retrieve the language tag from the Gin context.
//
// Methods:
// - GetMessage: Retrieves a localized message based on the provided context and parameter.
// - MustGetMessage: Retrieves a localized message and returns an empty string if retrieval fails.
// - HasLang: Retrieves a localized message and returns an empty string if retrieval fails.
// - SetBundle: Sets the i18n.Bundle configuration.
// - SetGetLngHandler: Sets the handler function to retrieve the language tag from the Gin context.
// - loadMessageFiles: Loads all localization files into the bundle.
// - loadMessageFile: Loads a single localization file into the bundle.
// - setLocalizerByLng: Sets the localizers for each accepted language.
// - newLocalizer: Creates a new localizer for a given language.
// - getLocalizerByLng: Retrieves the localizer for a given language.
// - getLocalizeConfig: Converts the parameter into an i18n.LocalizeConfig.
package i18n

import (
	"errors"
	"fmt"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// GinI18n is an interface that defines methods for internationalization (i18n) in a Gin web framework context.
// It provides methods to get localized messages and configure the i18n bundle and language handler.
var _ GinI18n = (*ginI18nImpl)(nil)

type ginI18nImpl struct {
	bundle          *i18n.Bundle
	localizerByLng  map[string]*i18n.Localizer
	defaultLanguage language.Tag
	getLngHandler   GetLngHandler
}

// getDefaultLanguage get default language
func (i *ginI18nImpl) getDefaultLanguage() language.Tag {
	return i.defaultLanguage
}

// getCurrentLanguage get default language
func (i *ginI18nImpl) getCurrentLanguage(context *gin.Context) language.Tag {
	return language.Make(i.getLngHandler(context, i.defaultLanguage.String()))
}

// HasLang check language is exist
func (i *ginI18nImpl) hasLang(language string) bool {
	if _, exist := i.localizerByLng[language]; exist {
		return true
	}

	return false
}

// GetMessage retrieves a localized message based on the provided context and parameter.
// If the message cannot be retrieved, it returns an empty string.
//
// Parameters:
//   - ctx: The Gin context from which to retrieve the message.
//   - param: The parameter used to fetch the localized message.
//
// Returns:
//   - string: The localized message or an empty string if retrieval fails.
//   - error: An error if the message retrieval fails.
func (i *ginI18nImpl) GetMessage(ctx *gin.Context, param interface{}) (string, error) {
	lng := i.getLngHandler(ctx, i.defaultLanguage.String())
	localizer := i.getLocalizerByLng(lng)

	localizeConfig, err := i.getLocalizeConfig(param)
	if err != nil {
		return "", err
	}

	message, err := localizer.Localize(localizeConfig)
	if err != nil {
		return "", err
	}

	return message, nil
}

// MustGetMessage retrieves a localized message based on the provided context and parameter.
// If the message cannot be retrieved, it returns an empty string.
// This method panics if the message retrieval fails.
//
// Parameters:
//   - ctx: The Gin context from which to retrieve the message.
//   - param: The parameter used to fetch the localized message.
//
// Returns:
//   - string: The localized message or an empty string if retrieval fails.
func (i *ginI18nImpl) MustGetMessage(ctx *gin.Context, param interface{}) string {
	message, _ := i.GetMessage(ctx, param)
	return message
}

// SetBundle initializes the i18n bundle with the provided configuration.
// It sets the default language, registers the unmarshal function for the bundle files,
// loads the message files, and sets the localizer based on the accepted languages.
//
// Parameters:
//   - cfg: A pointer to a BundleCfg struct that contains the configuration for the bundle.
func (i *ginI18nImpl) SetBundle(cfg *BundleCfg) {
	bundle := i18n.NewBundle(cfg.DefaultLanguage)
	bundle.RegisterUnmarshalFunc(cfg.FormatBundleFile, cfg.UnmarshalFunc)

	i.bundle = bundle
	i.defaultLanguage = cfg.DefaultLanguage

	i.loadMessageFiles(cfg)
	i.setLocalizerByLng(cfg.AcceptLanguage)
}

// SetGetLngHandler sets the handler function that will be used to get the language.
// The handler should be a function that implements the GetLngHandler interface.
//
// Parameters:
//
//	handler - a function that implements the GetLngHandler interface
func (i *ginI18nImpl) SetGetLngHandler(handler GetLngHandler) {
	i.getLngHandler = handler
}

// loadMessageFiles load all file localize to bundle
func (i *ginI18nImpl) loadMessageFiles(config *BundleCfg) {
	for _, lng := range config.AcceptLanguage {
		src := path.Join(config.RootPath, lng.String()) + "." + config.FormatBundleFile
		if err := i.loadMessageFile(config, src); err != nil {
			panic(err)
		}
	}
}

func (i *ginI18nImpl) loadMessageFile(config *BundleCfg, src string) error {
	buf, err := config.Loader.LoadMessage(src)
	if err != nil {
		return err
	}

	if _, err = i.bundle.ParseMessageFileBytes(buf, src); err != nil {
		return err
	}
	return nil
}

// setLocalizerByLng set localizer by language
func (i *ginI18nImpl) setLocalizerByLng(acceptLanguage []language.Tag) {
	i.localizerByLng = map[string]*i18n.Localizer{}
	for _, lng := range acceptLanguage {
		lngStr := lng.String()
		i.localizerByLng[lngStr] = i.newLocalizer(lngStr)
	}

	// set defaultLanguage if it isn't exist
	defaultLng := i.defaultLanguage.String()
	if _, hasDefaultLng := i.localizerByLng[defaultLng]; !hasDefaultLng {
		i.localizerByLng[defaultLng] = i.newLocalizer(defaultLng)
	}
}

// newLocalizer create a localizer by language
func (i *ginI18nImpl) newLocalizer(lng string) *i18n.Localizer {
	lngDefault := i.defaultLanguage.String()
	lngs := []string{
		lng,
	}

	if lng != lngDefault {
		lngs = append(lngs, lngDefault)
	}

	localizer := i18n.NewLocalizer(
		i.bundle,
		lngs...,
	)
	return localizer
}

// getLocalizerByLng get localizer by language
func (i *ginI18nImpl) getLocalizerByLng(lng string) *i18n.Localizer {
	localizer, hasValue := i.localizerByLng[lng]
	if hasValue {
		return localizer
	}

	return i.localizerByLng[i.defaultLanguage.String()]
}

func (i *ginI18nImpl) getLocalizeConfig(param interface{}) (*i18n.LocalizeConfig, error) {
	switch paramValue := param.(type) {
	case string:
		localizeConfig := &i18n.LocalizeConfig{
			MessageID: paramValue,
		}
		return localizeConfig, nil
	case *i18n.Message:
		localizeConfig := &i18n.LocalizeConfig{
			DefaultMessage: paramValue,
		}
		return localizeConfig, nil
	case i18n.Message:
		localizeConfig := &i18n.LocalizeConfig{
			DefaultMessage: &paramValue,
		}
		return localizeConfig, nil
	case *i18n.LocalizeConfig:
		return paramValue, nil
	case i18n.LocalizeConfig:
		return &paramValue, nil
	}

	msg := fmt.Sprintf("un supported localize param: %v", param)
	return nil, errors.New(msg)
}
