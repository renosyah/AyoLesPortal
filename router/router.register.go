package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/renosyah/AyoLesPortal/api"
	"github.com/renosyah/AyoLesPortal/util"
)

func Register(w http.ResponseWriter, r *http.Request) {

	// check any message from form
	message := r.FormValue("message")

	// render html template
	// with message add to map
	err := temp.ExecuteTemplate(w, "register.html", map[string]interface{}{
		"message": message,
	})
	if err != nil {
		http.Redirect(w, r, fmt.Sprintf("/error?message=%s", err.Error()), http.StatusSeeOther)
		return
	}
}

func SubmitRegister(w http.ResponseWriter, r *http.Request) {

	// check if submit method not post
	if r.Method != http.MethodPost {
		http.Redirect(w, r, fmt.Sprintf("/login?message=%s", "method must be post"), http.StatusSeeOther)
		return
	}

	// get from form
	// add to object teacher
	teacher := api.AddTeacherParam{
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	// check if string is realy empty
	if strings.TrimSpace(teacher.Name) == "" || strings.TrimSpace(teacher.Email) == "" || strings.TrimSpace(teacher.Password) == "" {
		http.Redirect(w, r, fmt.Sprintf("/register?message=%s", "form must not be empty"), http.StatusSeeOther)
		return
	}

	// add new teacher
	log, errs := teacherModule.Add(r.Context(), teacher)
	if errs != nil {
		http.Redirect(w, r, fmt.Sprintf("%s?status=%s", r.URL.Path, errs.Message), http.StatusSeeOther)
		return
	}

	// set cookie
	err := cookieConfig.Set(w, r, util.CookieData{
		Name:  util.DefaultName,
		Value: log.ID.String(),
	})
	if err != nil {
		http.Redirect(w, r, fmt.Sprintf("/error?message=%s", err.Error()), http.StatusSeeOther)
		return
	}

	// redirect to dashboard
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}
