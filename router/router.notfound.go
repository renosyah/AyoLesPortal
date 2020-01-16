package router

import (
	"fmt"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	err := temp.ExecuteTemplate(w, "notfound.html", nil)
	if err != nil {
		fmt.Fprintln(w, http.StatusInternalServerError)
	}
}
