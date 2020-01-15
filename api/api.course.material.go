package api

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/renosyah/AyoLesPortal/model"
	"github.com/renosyah/AyoLesPortal/util"
	uuid "github.com/satori/go.uuid"
)

type (
	CourseMaterialModule struct {
		r    *util.PostData
		Name string
	}

	AddCourseMaterialParam struct {
		CourseID      uuid.UUID `json:"course_id"`
		MaterialIndex int32     `json:"material_index"`
		Title         string    `json:"title"`
	}

	OneCourseMaterialParam struct {
		ID uuid.UUID `json:"id"`
	}

	AllCourseMaterialParam struct {
		CourseID    uuid.UUID `json:"course_id"`
		SearchBy    string    `json:"search_by"`
		SearchValue string    `json:"search_value"`
		OrderBy     string    `json:"order_by"`
		OrderDir    string    `json:"order_dir"`
		Offset      int       `json:"offset"`
		Limit       int       `json:"limit"`
	}
)

func NewCourseMaterialModule(r *util.PostData) *CourseMaterialModule {
	return &CourseMaterialModule{
		r:    r,
		Name: "module/course_material",
	}
}

func (m CourseMaterialModule) All(ctx context.Context, param AllCourseMaterialParam) ([]model.CourseMaterialResponse, *Error) {
	var allResp []model.CourseMaterialResponse

	data, err := (&model.CourseMaterial{}).All(ctx, m.r, model.AllCourseMaterial{
		CourseID:    param.CourseID,
		SearchBy:    param.SearchBy,
		SearchValue: param.SearchValue,
		OrderBy:     param.OrderBy,
		OrderDir:    param.OrderDir,
		Offset:      param.Offset,
		Limit:       param.Limit,
	})
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on query all course material"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no Course material found"
		}

		return []model.CourseMaterialResponse{}, NewErrorWrap(err, m.Name, "all/course_material",
			message, status)
	}

	for _, each := range data {
		allResp = append(allResp, each.Response())
	}

	return allResp, nil

}
func (m CourseMaterialModule) Add(ctx context.Context, param AddCourseMaterialParam) (model.CourseMaterialResponse, *Error) {
	courseMaterial := &model.CourseMaterial{
		CourseID:      param.CourseID,
		MaterialIndex: param.MaterialIndex,
		Title:         param.Title,
	}

	id, err := courseMaterial.Add(ctx, m.r)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on add course material"

		return model.CourseMaterialResponse{}, NewErrorWrap(err, m.Name, "add/course_material",
			message, status)
	}

	courseMaterial.ID = id

	return courseMaterial.Response(), nil
}

func (m CourseMaterialModule) One(ctx context.Context, param OneCourseMaterialParam) (model.CourseMaterialResponse, *Error) {
	courseMaterial := &model.CourseMaterial{
		ID: param.ID,
	}

	data, err := courseMaterial.One(ctx, m.r)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on get one course material"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no course material found"
		}

		return model.CourseMaterialResponse{}, NewErrorWrap(err, m.Name, "one/course_material",
			message, status)
	}

	return data.Response(), nil
}
