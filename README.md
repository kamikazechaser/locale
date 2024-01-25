# locale

> Simple locale template rendering using KV maps only

[![Go Reference](https://pkg.go.dev/badge/github.com/kamikazechaser/locale.svg)](https://pkg.go.dev/github.com/kamikazechaser/locale)

Provides a _simple_ and _flexible_ way to render your translations. You are not locked into any i18n/i10n standard.

## Install:

```
go get github.com/kamikazechaser/locale

```

## Example

```go
package main

import (
	"log"

	"github.com/kamikazechaser/locale"
)

func main() {
	lMap := locale.LangMap{
		"eng": locale.Map{
			"hello": "Hello {{ .Name }}",
		},
		"swa": locale.Map{
			"hello": "Habari {{ .Name }}",
		},
	}

	l, err := locale.NewLocale(lMap, "eng")
	if err != nil {
		log.Fatal(err)
	}

	payload := struct {
		Name string
	}{
		"Mohamed Sohail",
	}

	// Will use the default language here.
	eR, err := l.Render("hello", locale.WithPayload(payload))
	if err != nil {
		log.Fatal(err)
	}

	// Using all available Render optional function options.
	sR, err := l.Render("hello", locale.WithLangCode("swa"), locale.WithPayload(payload))
	if err != nil {
		log.Fatal(err)
	}

	log.Println(eR)
	log.Println(sR)
}

```

For more info, See the Go package reference.

## License

The Unlicense.

## Credits

- Some of the API inspired by https://github.com/kataras/i18n
