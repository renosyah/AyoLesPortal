package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/renosyah/AyoLesPortal/api"
	"github.com/renosyah/AyoLesPortal/util"
)

func Register(w http.ResponseWriter, r *http.Request) {
	message := r.FormValue("message")
	err := temp.ExecuteTemplate(w, "register.html", map[string]interface{}{
		"message": message,
	})
	if err != nil {
		http.Redirect(w, r, fmt.Sprintf("/error?message=%s", err.Error()), http.StatusSeeOther)
		return
	}
}

func SubmitRegister(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	teacher := api.AddTeacherParam{
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	if strings.TrimSpace(teacher.Name) == "" || strings.TrimSpace(teacher.Email) == "" || strings.TrimSpace(teacher.Password) == "" {
		http.Redirect(w, r, fmt.Sprintf("/register?message=%s", "form must not be empty"), http.StatusSeeOther)
		return
	}

	log, errs := teacherModule.Add(r.Context(), teacher)
	if errs != nil {
		http.Redirect(w, r, fmt.Sprintf("%s?status=%s", r.URL.Path, errs.Message), http.StatusSeeOther)
		return
	}

	err := cookieConfig.Set(w, r, util.CookieData{
		Name:  util.DefaultName,
		Value: log.ID.String(),
	})
	if err != nil {
		http.Redirect(w, r, fmt.Sprintf("/error?message=%s", err.Error()), http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}
