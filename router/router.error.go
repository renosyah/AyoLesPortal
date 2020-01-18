package router

import (
	"fmt"
	"net/http"
	"strings"
)

func ErrorPage(w http.ResponseWriter, r *http.Request) {
	msg := r.FormValue("message")
	err := temp.ExecuteTemplate(w, "error.html", map[string]string{
		"message": strings.ToUpper(msg),
	})
	if err != nil {
		fmt.Fprintln(w, http.StatusInternalServerError)
	}
}
