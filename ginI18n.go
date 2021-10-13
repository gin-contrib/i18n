package i18n

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var (
	_ I18n = (*ginI18n)(nil)
)

type ginI18n struct {
	bundle           *i18n.Bundle
	currentContext   *gin.Context
	localizerByLng   map[string]*i18n.Localizer
	defaultLanguage  language.Tag
}

type Config struct {
	DefaultLanguage  language.Tag
	FormatBundleFile string
	AcceptLanguage   []language.Tag
	RootPath         string
	UnmarshalFunc    i18n.UnmarshalFunc
}

func NewI18n(config *Config) {
	bundle := i18n.NewBundle(config.DefaultLanguage)
	bundle.RegisterUnmarshalFunc(config.FormatBundleFile, config.UnmarshalFunc)
	ins := &ginI18n{
		bundle:           bundle,
		defaultLanguage:  config.DefaultLanguage,
	}
	ins.loadMessageFiles(config)
	ins.setLocalizerByLng(config.AcceptLanguage)

	GinI18n = ins
}

var GinI18n I18n

// loadMessageFiles load all file localize to bundle
func (i *ginI18n) loadMessageFiles(config *Config) {
	for _, lng := range config.AcceptLanguage {
		path := fmt.Sprintf("%s/%s.%s", config.RootPath, lng.String(), config.FormatBundleFile)
		i.bundle.MustLoadMessageFile(path)
	}
}

// setLocalizerByLng set localizer by language
func (i *ginI18n) setLocalizerByLng(acceptLanguage []language.Tag) {
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
func (i *ginI18n) newLocalizer(lng string) *i18n.Localizer {
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
func (i *ginI18n) getLocalizerByLng(lng string) *i18n.Localizer {
	localizer, hasValue := i.localizerByLng[lng]
	if hasValue {
		return localizer
	}

	return i.localizerByLng[i.defaultLanguage.String()]
}

// GetMessage get localize message by lng and messageID
func (i *ginI18n) GetMessage(param interface{}) (string, error) {
	lng := getLngFromGinContext(i.currentContext)
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

func (i *ginI18n) getLocalizeConfig(param interface{}) (*i18n.LocalizeConfig, error) {
	switch paramValue := param.(type) {
	case string:
		localizeConfig := &i18n.LocalizeConfig{
			MessageID: paramValue,
		}
		return localizeConfig, nil
	case *LocalizeConfig:
		result := i18n.LocalizeConfig(*paramValue)
		return &result, nil
	}

	msg := fmt.Sprintf("un supported localize param: %v", param)
	return nil, errors.New(msg)
}

// MustGetMessage ...
func (i *ginI18n) MustGetMessage(param interface{}) string {
	message, _ := i.GetMessage(param)
	return message
}

func (i *ginI18n) setCurrentContext(ctx context.Context) {
	i.currentContext = ctx.(*gin.Context)
}
