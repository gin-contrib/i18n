package i18n

import "golang.org/x/text/language"

func sliceLanguageToMap(lngs []language.Tag) map[language.Tag]bool {
	result := map[language.Tag]bool{}
	for _, lng := range lngs {
		result[lng] = true
	}

	return result
}
