package router

import (
	"net/http"

	"github.com/renosyah/AyoLesPortal/util"
)

func getCookie(r *http.Request) *util.CookieData {
	cok := &util.CookieData{
		Name: util.DefaultName,
	}
	cookieConfig.Get(r, cok)

	return cok
}
