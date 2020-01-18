package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/renosyah/AyoLesPortal/api"
	"github.com/renosyah/AyoLesPortal/util"
)

func Login(w http.ResponseWriter, r *http.Request) {
	message := r.FormValue("message")
	err := temp.ExecuteTemplate(w, "login.html", map[string]interface{}{
		"message": message,
	})
	if err != nil {
		http.Redirect(w, r, fmt.Sprintf("/error?message=%s", err.Error()), http.StatusSeeOther)
		return
	}
}

func SubmitLogin(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	teacher := api.TeacherLoginParam{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	if strings.TrimSpace(teacher.Email) == "" || strings.TrimSpace(teacher.Password) == "" {
		http.Redirect(w, r, fmt.Sprintf("/login?message=%s", "email or password must not be empty"), http.StatusSeeOther)
		return
	}

	log, errs := teacherModule.Login(r.Context(), teacher)
	if errs != nil {
		http.Redirect(w, r, fmt.Sprintf("/login?message=%s", errs.Message), http.StatusSeeOther)
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
