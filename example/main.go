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
		"Sohail",
	}

	eR, err := l.Render("hello", "eng", payload)
	if err != nil {
		log.Fatal(err)
	}

	sR, err := l.Render("hello", "swa", payload)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(eR)
	log.Println(sR)
}
