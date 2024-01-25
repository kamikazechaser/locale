package locale

import "errors"

var (
	ErrDefaultLangNotSet           = errors.New("default lang cannot be empty")
	ErrDefaultLangNonExistentOnMap = errors.New("default lang does not exist on lang map")
)

// validate is an internal function that validates the lang map based on the default lang.
func validate(langMap LangMap, defaultLang string) error {
	if defaultLang == "" {
		return ErrDefaultLangNotSet
	}

	if _, exists := langMap[defaultLang]; !exists {
		return ErrDefaultLangNonExistentOnMap
	}

	return nil
}
