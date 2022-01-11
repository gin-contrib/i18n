package i18n

import (
	"context"
)

// GinI18n ...
type GinI18n interface {
	getMessage(param interface{}) (string, error)
	mustGetMessage(param interface{}) string
	setCurrentContext(ctx context.Context)
	setBundle(cfg *BundleCfg)
	setCustomerBundle(cfg *BundleCfg)
	setGetLngHandler(handler GetLngHandler)
}
