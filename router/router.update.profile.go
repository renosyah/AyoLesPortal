package router

import (
	"fmt"
	"net/http"

	"github.com/renosyah/AyoLesPortal/api"
	uuid "github.com/satori/go.uuid"
)

func UpdateProfile(w http.ResponseWriter, r *http.Request) {

	// check if method is not post
	if r.Method != http.MethodPost {
		http.Redirect(w, r, fmt.Sprintf("/login?message=%s", "method must be post"), http.StatusSeeOther)
		return
	}

	id, _ := uuid.FromString(r.FormValue("id"))

	teacher := api.TeacherParam{
		ID:       id,
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	err := teacherModule.Update(r.Context(), teacher)
	if err != nil {
		http.Redirect(w, r, fmt.Sprintf("/error?message=%s", err.Error()), http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/dashboard?menu=editprofile", http.StatusSeeOther)
}
