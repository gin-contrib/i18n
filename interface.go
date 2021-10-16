package gini18n

import (
	"context"
)

type GinI18n interface {
	getMessage(param interface{}) (string, error)
	mustGetMessage(param interface{}) string

	setCurrentContext(ctx context.Context)
	setBundle(cfg *BundleCfg)
	setGetLngHandler(handler GetLngHandler)
}
