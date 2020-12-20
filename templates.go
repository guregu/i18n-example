package main

import (
	"context"
	"html/template"
	"path/filepath"
)

var templates *template.Template

func getTemplate(ctx context.Context, name string) *template.Template {
	t := templates.Lookup(name + ".gohtml")
	if t == nil {
		panic("no template: " + name)
	}

	t, err := t.Clone()
	if err != nil {
		panic(err)
	}
	t.Funcs(templateFuncs(ctx))
	return t
}

func templateFuncs(ctx context.Context) template.FuncMap {
	localizer, ok := localizerFrom(ctx)
	if !ok {
		localizer = defaultLocalizer
	}

	return template.FuncMap{
		"t":  translateFunc(localizer),
		"tc": translateCountFunc(localizer),
	}
}

func loadTemplates() error {
	glob := filepath.Join("assets", "templates", "*.gohtml")

	t := template.New("root").Funcs(templateFuncs(context.Background()))
	var err error
	t, err = t.ParseGlob(glob)
	if err != nil {
		return err
	}

	templates = t
	return nil
}
