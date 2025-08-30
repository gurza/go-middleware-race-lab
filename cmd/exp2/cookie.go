package main

import (
	"context"
	"net/http"
)

type cookieSetter interface {
	Set(*http.Cookie)
}

type cookieSetterKey struct{}

type cookieSetterImpl struct{ w http.ResponseWriter }

func (c cookieSetterImpl) Set(ck *http.Cookie) { http.SetCookie(c.w, ck) }

func withCookieSetter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), cookieSetterKey{}, cookieSetterImpl{w})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func cookieSetterFrom(ctx context.Context) (cookieSetter, bool) {
	cs, ok := ctx.Value(cookieSetterKey{}).(cookieSetter)
	return cs, ok
}
