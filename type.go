package i18n

type (
	KeyMessage struct {
		_              struct{}
		I18nKey        string
		DefaultMessage string
	}
	MapStringKeyMessage map[string]KeyMessage
	MapIntKeyMessage    map[int]KeyMessage
)

func (k *KeyMessage) GetMessage() string {
	if k.I18nKey == "" {
		return k.DefaultMessage
	}

	return GinI18n.MustGetMessage(k.I18nKey)
}
