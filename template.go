package locale

import (
	"fmt"
	"text/template"
)

const (
	baseTmplIdentifier = "base_"
)

func parseAndLoadTemplates(langMap LangMap) (*template.Template, error) {
	baseTmpl := template.New(baseTmplIdentifier)

	for langCode, tmplMap := range langMap {
		for tmplKey, tmplValue := range tmplMap {
			if _, err := baseTmpl.New(fmt.Sprintf("%s_%s", tmplKey, langCode)).Parse(tmplValue); err != nil {
				return nil, fmt.Errorf("locale (internal) failed to parse and load templates: %v", err)
			}
		}
	}

	return baseTmpl, nil
}
