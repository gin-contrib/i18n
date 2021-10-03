package i18n

import (
	"github.com/gin-gonic/gin"
)

func NewError(message KeyMessage) error {
	result := &errorImpl{
		keyMessage: message,
	}

	return result
}

type errorImpl struct {
	keyMessage KeyMessage
}

func (a *errorImpl) Error() string {
	return a.keyMessage.GetMessage()
}

type ErrFunctionHandler func(ctx *gin.Context) error
func errHandler(message KeyMessage) ErrFunctionHandler {
	return func(ctx *gin.Context) error {
		return NewError(message)
	}
}
