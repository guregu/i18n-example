package main

import (
	"context"
	"net/http"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type localizerKey struct{}

func personalize(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		acceptLang := r.Header.Get("Accept-Language")
		localizer := i18n.NewLocalizer(translations, acceptLang)
		ctx = withLocalizer(ctx, localizer)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func withLocalizer(ctx context.Context, loc *i18n.Localizer) context.Context {
	return context.WithValue(ctx, localizerKey{}, loc)
}

func localizerFrom(ctx context.Context) (*i18n.Localizer, bool) {
	loc, ok := ctx.Value(localizerKey{}).(*i18n.Localizer)
	return loc, ok
}
