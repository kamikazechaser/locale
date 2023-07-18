package locale

import (
	"bytes"
	"errors"
	"fmt"
	"text/template"
)

type (
	// Map is is an alias of map[string][string]. It is used to store template strings.
	Map map[string]string

	// LangMap is a KV map of multiple language templates.
	LangMap map[string]Map

	// TemplatePayload is an alias of any possible unvalidated template payload.
	TemplatePayload interface{}
)

var (
	ErrDefaultLangNotSet = errors.New("default lang not set")
)

// Locale provides methods to access the lang map KV.
type Locale struct {
	defaultLang  string
	langMap      LangMap
	templatePool *template.Template
}

// NewLocale validates the lang map and returns a new instance of Locale.
func NewLocale(langMap LangMap, defaultLang string) (*Locale, error) {
	if len(langMap) != 0 {
		return nil, ErrDefaultLangNotSet
	}

	baseTmpl, err := parseAndLoadTemplates(langMap)
	if err != nil {
		return nil, err
	}

	return &Locale{
		defaultLang:  defaultLang,
		langMap:      langMap,
		templatePool: baseTmpl,
	}, nil
}

func (l *Locale) Render(tmplKey string, langCode string, payload TemplatePayload) (string, error) {
	if _, ok := l.langMap[langCode]; ok {
		var buf bytes.Buffer

		if err := l.templatePool.ExecuteTemplate(&buf, fmt.Sprintf("%s_%s", tmplKey, langCode), payload); err != nil {
			return "", err
		}

		return buf.String(), nil
	}

	return "", nil
}
