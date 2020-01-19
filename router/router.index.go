package router

import (
	"fmt"
	"net/http"

	"github.com/renosyah/AyoLesPortal/util"
)

func Index(w http.ResponseWriter, r *http.Request) {

	// get cookie
	cok := &util.CookieData{
		Name: util.DefaultName,
	}
	cookieConfig.Get(r, cok)

	// check if value not empty
	if cok.Value != "" {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	// render html template
	err := temp.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Redirect(w, r, fmt.Sprintf("/error?message=%s", err.Error()), http.StatusSeeOther)
		return
	}
}
