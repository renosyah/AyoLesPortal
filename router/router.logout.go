package router

import (
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	cookieConfig.Delete(w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
