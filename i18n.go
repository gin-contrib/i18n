package i18n

import (
	"context"
)

type I18n interface {
	GetMessage(param interface{}) (string, error)
	MustGetMessage(param interface{}) string

	setCurrentContext(ctx context.Context)
}
