package auth

import (
	"net/http"

	"github.com/renosyah/AyoLesPortal/util"
)

func SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cok := &util.CookieData{}
		cookieConfig.Get(r, cok)

		if cok.Value == "" && r.URL.Path != "/" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func EmptyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)
	})
}
