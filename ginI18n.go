package i18n

import (
	"context"
	"errors"
	"fmt"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var _ GinI18n = (*ginI18nImpl)(nil)

type ginI18nImpl struct {
	bundle          *i18n.Bundle
	currentContext  *gin.Context
	localizerByLng  map[string]*i18n.Localizer
	defaultLanguage language.Tag
	getLngHandler   GetLngHandler
}

// getMessage get localize message by lng and messageID
func (i *ginI18nImpl) getMessage(param interface{}) (string, error) {
	lng := i.getLngHandler(i.currentContext, i.defaultLanguage.String())
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

// mustGetMessage ...
func (i *ginI18nImpl) mustGetMessage(param interface{}) string {
	message, _ := i.getMessage(param)
	return message
}

func (i *ginI18nImpl) setCurrentContext(ctx context.Context) {
	i.currentContext = ctx.(*gin.Context)
}

func (i *ginI18nImpl) setBundle(cfg *BundleCfg) {
	bundle := i18n.NewBundle(cfg.DefaultLanguage)
	bundle.RegisterUnmarshalFunc(cfg.FormatBundleFile, cfg.UnmarshalFunc)

	i.bundle = bundle
	i.defaultLanguage = cfg.DefaultLanguage

	i.loadMessageFiles(cfg)
	i.setLocalizerByLng(cfg.AcceptLanguage)
}

func (i *ginI18nImpl) setGetLngHandler(handler GetLngHandler) {
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
	case *i18n.LocalizeConfig:
		return paramValue, nil
	}

	msg := fmt.Sprintf("un supported localize param: %v", param)
	return nil, errors.New(msg)
}
