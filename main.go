package main

import (
	"flag"
	"log"
	"net/http"
)

var bindFlag = flag.String("bind", ":8000", "address listen on")

func main() {
	if err := loadTranslations(); err != nil {
		log.Fatal(err)
	}

	if err := loadTemplates(); err != nil {
		log.Fatal(err)
	}

	http.Handle("/", personalize(http.HandlerFunc(index)))
	http.Handle("/favicon.ico", http.NotFoundHandler())

	log.Println("Starting server", *bindFlag)
	if err := http.ListenAndServe(*bindFlag, nil); err != nil {
		log.Fatal(err)
	}
}
