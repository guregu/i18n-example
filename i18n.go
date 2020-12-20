package main

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var (
	defaultLang = language.Japanese
)

var translations *i18n.Bundle
var defaultLocalizer *i18n.Localizer

func loadTranslations() error {
	dir := filepath.Join("assets", "text")
	bundle := i18n.NewBundle(defaultLang)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || filepath.Ext(path) != ".toml" {
			return nil
		}
		_, err = bundle.LoadMessageFile(path)
		return err
	})
	if err != nil {
		return err
	}

	translations = bundle
	defaultLocalizer = i18n.NewLocalizer(bundle, defaultLang.String())
	return nil
}

func translateFunc(localizer *i18n.Localizer) interface{} {
	return func(id string, args ...interface{}) string {
		var data map[string]interface{}
		if len(args) > 0 {
			data = make(map[string]interface{}, len(args))
			for n, iface := range args {
				data["v"+strconv.Itoa(n)] = iface
			}
		}
		str, _, err := localizer.LocalizeWithTag(&i18n.LocalizeConfig{
			MessageID:    id,
			TemplateData: data,
		})
		if str == "" && err != nil {
			return "[TL err: " + err.Error() + "]"
		}
		return str
	}
}

func translateCountFunc(localizer *i18n.Localizer) interface{} {
	return func(id string, ct int, args ...interface{}) string {
		data := make(map[string]interface{}, len(args)+1)
		if len(args) > 0 {
			for n, iface := range args {
				data["v"+strconv.Itoa(n)] = iface
			}
		}
		data["ct"] = ct
		str, _, err := localizer.LocalizeWithTag(&i18n.LocalizeConfig{
			MessageID:    id,
			TemplateData: data,
			PluralCount:  ct,
		})
		if str == "" && err != nil {
			return "[TL err: " + err.Error() + "]"
		}
		return str
	}
}
