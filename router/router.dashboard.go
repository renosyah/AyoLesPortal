package router

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	uuid "github.com/satori/go.uuid"

	"github.com/renosyah/AyoLesPortal/api"
	"github.com/renosyah/AyoLesPortal/model"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	id, _ := uuid.FromString(getCookie(r).Value)
	dteacher, errAp := teacherModule.One(ctx, api.OneTeacherParam{ID: id})
	if errAp != nil {
		http.Redirect(w, r, fmt.Sprintf("/error?message=%s", errAp.Message), http.StatusSeeOther)
		return
	}

	listCourse, err := getCourses(ctx, w, r, dteacher)
	if err != nil {
		http.Redirect(w, r, fmt.Sprintf("/error?message=%s", err.Error()), http.StatusSeeOther)
		return
	}

	data := map[string]interface{}{"teacher": dteacher}

	switch r.FormValue("menu") {
	case "newcourse":
		data["newCourse"] = model.Course{}
		break

	case "listcourse":
		data["listCourse"] = listCourse
		break

	case "newcoursematerial":
		data["newCourseMaterial"] = model.CourseMaterial{}
		break

	case "listcoursematerial":
		data["listCourseMaterial"] = listCourse
		break

	case "newexam":
		data["newCourseExam"] = model.CourseExam{}
		break

	case "listexam":
		data["listCourseExam"] = listCourse
		break

	case "newcategory":
		data["newCategory"] = model.Category{}
		break

	case "listcategory":
		listCategory, err := getCategories(ctx, w, r)
		if err != nil {
			http.Redirect(w, r, fmt.Sprintf("/error?message=%s", err.Error()), http.StatusSeeOther)
			return
		}

		data["listCategory"] = listCategory
		break

	case "editprofile":
		data["teacherProfile"] = dteacher
		break

	default:
		break
	}

	errTemp := temp.ExecuteTemplate(w, "dashboard.html", data)
	if errTemp != nil {
		http.Redirect(w, r, fmt.Sprintf("/error?message=%s", errTemp.Error()), http.StatusSeeOther)
		return
	}
}

//-------------------//

func getCourses(ctx context.Context, w http.ResponseWriter, r *http.Request, t model.TeacherResponse) ([]model.CourseResponse, *api.Error) {

	offset, _ := strconv.Atoi(r.FormValue("offset"))
	limit, _ := strconv.Atoi(r.FormValue("limit"))

	listCourse, err := courseModule.All(ctx, api.AllCourseParam{
		TeacherID:   t.ID,
		SearchBy:    "course_name",
		SearchValue: "",
		OrderBy:     "course_name",
		OrderDir:    "asc",
		Offset:      offset,
		Limit:       limit,
	})
	if err != nil {
		return listCourse, err
	}

	return listCourse, nil
}

func getCategories(ctx context.Context, w http.ResponseWriter, r *http.Request) ([]model.CategoryResponse, *api.Error) {

	offset, _ := strconv.Atoi(r.FormValue("offset"))
	limit, _ := strconv.Atoi(r.FormValue("limit"))

	listCategory, err := categoryModule.All(ctx, api.AllCategoryParam{
		SearchBy:    "name",
		SearchValue: "",
		OrderBy:     "name",
		OrderDir:    "asc",
		Offset:      offset,
		Limit:       limit,
	})
	if err != nil {
		return listCategory, err
	}

	return listCategory, nil
}
