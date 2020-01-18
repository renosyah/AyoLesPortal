package router

import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"

	"github.com/renosyah/AyoLesPortal/api"
	"github.com/renosyah/AyoLesPortal/util"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	cok := &util.CookieData{
		Name: util.DefaultName,
	}
	cookieConfig.Get(r, cok)

	id, err := uuid.FromString(cok.Value)
	if err != nil {
		http.Redirect(w, r, fmt.Sprintf("/error?message=%s", err.Error()), http.StatusSeeOther)
		return
	}

	dteacher, errAp := teacherModule.One(r.Context(), api.OneTeacherParam{
		ID: id,
	})
	if errAp != nil {
		http.Redirect(w, r, fmt.Sprintf("/error?message=%s", errAp.Message), http.StatusSeeOther)
		return
	}

	err = temp.ExecuteTemplate(w, "dashboard.html", map[string]interface{}{
		"teacher": dteacher,
	})
	if err != nil {
		http.Redirect(w, r, fmt.Sprintf("/error?message=%s", err.Error()), http.StatusSeeOther)
		return
	}
}
