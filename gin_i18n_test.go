package i18n

import (
	"log"
	"testing"
)

func init() {
	NewI18nImpl("./example/localize")
}

func Test_testI18n(t *testing.T) {
	message, _ := GinI18n.GetMessage("welcome")
	log.Println("Message: ", message)

	message, _ = GinI18n.GetMessage( "welcome")
	log.Println("Message: ", message)

	message, _ = GinI18n.GetMessage("welcome")
	log.Println("Message: ", message)
}
