package locale

import (
	"bytes"
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

	// RenderOption describes the function signature of optional Render options
	RenderOption func(c *renderConfig)

	// Locale provides methods to access the lang map KV e.g. Render templates
	Locale struct {
		defaultLang  string
		langMap      LangMap
		templatePool *template.Template
	}

	// renderConfig is an internal struct to hold optional Render options
	renderConfig struct {
		langCode string
		payload  TemplatePayload
	}
)

// NewLocale validates the lang map and returns a new instance of Locale.
func NewLocale(langMap LangMap, defaultLang string) (*Locale, error) {
	if err := validate(langMap, defaultLang); err != nil {
		return nil, err
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

// WithLangCode is an optional functional option for Render to set the language code.
func WithLangCode(langCode string) RenderOption {
	return func(c *renderConfig) {
		c.langCode = langCode
	}
}

// WithPayload is an optional functional option for Render to set the template payload.
func WithPayload(payload TemplatePayload) RenderOption {
	return func(c *renderConfig) {
		c.payload = payload
	}
}

// Render accepts a temeplate key and other optional functional options to Render a specific template
func (l *Locale) Render(tmplKey string, opts ...RenderOption) (string, error) {
	c := &renderConfig{
		langCode: l.defaultLang,
	}

	for _, opt := range opts {
		opt(c)
	}

	if _, ok := l.langMap[c.langCode]; ok {
		var buf bytes.Buffer

		if err := l.templatePool.ExecuteTemplate(&buf, fmt.Sprintf("%s_%s", tmplKey, c.langCode), c.payload); err != nil {
			return "", err
		}

		return buf.String(), nil
	}

	return "", nil
}
