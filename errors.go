package i18n

import (
	"github.com/gin-gonic/gin"
)

func NewError(ctx *gin.Context, message KeyMessage) error {
	result := &errorImpl{
		context: ctx,
		keyMessage: message,
	}

	return result
}

type errorImpl struct {
	keyMessage KeyMessage
	context *gin.Context
}

func (a *errorImpl) Error() string {
	return a.keyMessage.GetMessageFromGinContext(a.context)
}

type ErrFunctionHandler func(ctx *gin.Context) error
func errHandler(message KeyMessage) ErrFunctionHandler {
	return func(ctx *gin.Context) error {
		return NewError(ctx, message)
	}
}
