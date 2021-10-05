package i18n

import (
	"log"
	"testing"
)

func init() {
	NewI18n("./example/localize")
}

func Test_testI18n(t *testing.T) {
	message, _ := GinI18n.GetMessage("welcome")
	log.Println("Message: ", message)

	message, _ = GinI18n.GetMessage(&LocalizeConfig{
		MessageID: "welcomeWithName",
		TemplateData: map[string]string{
			"name": "aksJH",
		},
	})
	log.Println("Message: ", message)

	message, _ = GinI18n.GetMessage("welcome")
	log.Println("Message: ", message)
}
