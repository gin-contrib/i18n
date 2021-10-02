package i18n

import (
	"gin-i18n/pkg/logger"
	"testing"
)

func init() {
	NewI18nImpl("./example/localize")
}

func Test_testI18n(t *testing.T) {
	logger.AtLog.Info("tada")
	message := AutoI18n.GetMessage("welcome")
	logger.AtLog.Info("Message: ", message)

	message = AutoI18n.GetMessage( "welcome")
	logger.AtLog.Info("Message: ", message)

	message = AutoI18n.GetMessage("welcome")
	logger.AtLog.Info("Message: ", message)
}
