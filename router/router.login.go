package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/renosyah/AyoLesPortal/api"
	"github.com/renosyah/AyoLesPortal/util"
)

func Login(w http.ResponseWriter, r *http.Request) {

	// get message from form
	message := r.FormValue("message")

	// render html
	// and add message to map
	err := temp.ExecuteTemplate(w, "login.html", map[string]interface{}{
		"message": message,
	})
	if err != nil {
		http.Redirect(w, r, fmt.Sprintf("/error?message=%s", err.Error()), http.StatusSeeOther)
		return
	}
}

func SubmitLogin(w http.ResponseWriter, r *http.Request) {

	// check if method is not post
	if r.Method != http.MethodPost {
		http.Redirect(w, r, fmt.Sprintf("/login?message=%s", "method must be post"), http.StatusSeeOther)
		return
	}

	// get value form
	// and add to object
	teacher := api.TeacherLoginParam{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	// check if value of string is not empty
	if strings.TrimSpace(teacher.Email) == "" || strings.TrimSpace(teacher.Password) == "" {
		http.Redirect(w, r, fmt.Sprintf("/login?message=%s", "email or password must not be empty"), http.StatusSeeOther)
		return
	}

	// send form login
	log, errs := teacherModule.Login(r.Context(), teacher)
	if errs != nil {
		http.Redirect(w, r, fmt.Sprintf("/login?message=%s", errs.Message), http.StatusSeeOther)
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
