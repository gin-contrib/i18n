package i18n

func (k *KeyMessage) GetMessage() string {
	if k.I18nKey == "" {
		return k.DefaultMessage
	}

	return GinI18n.MustGetMessage(k.I18nKey)
}
