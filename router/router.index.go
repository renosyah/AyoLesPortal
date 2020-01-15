package router

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	err := temp.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		fmt.Fprintln(w, http.StatusInternalServerError)
	}
}
