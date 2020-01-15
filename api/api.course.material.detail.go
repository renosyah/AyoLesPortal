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
	CourseMaterialDetailModule struct {
		r    *util.PostData
		Name string
	}

	AddCourseMaterialDetailParam struct {
		CourseMaterialID uuid.UUID `json:"course_material_id"`
		Position         int32     `json:"position"`
		Title            string    `json:"title"`
		TypeMaterial     int32     `json:"type_material"`
		Content          string    `json:"content"`
		ImageURL         string    `json:"image_url"`
	}

	AllCourseMaterialDetailParam struct {
		CourseMaterialID uuid.UUID `json:"course_material_id"`
		SearchBy         string    `json:"search_by"`
		SearchValue      string    `json:"search_value"`
		OrderBy          string    `json:"order_by"`
		OrderDir         string    `json:"order_dir"`
		Offset           int       `json:"offset"`
		Limit            int       `json:"limit"`
	}
)

func NewCourseMaterialDetailModule(r *util.PostData) *CourseMaterialDetailModule {
	return &CourseMaterialDetailModule{
		r:    r,
		Name: "module/course_material_detail",
	}
}

func (m CourseMaterialDetailModule) All(ctx context.Context, param AllCourseMaterialDetailParam) ([]model.CourseMaterialDetailResponse, *Error) {
	var allResp []model.CourseMaterialDetailResponse

	data, err := (&model.CourseMaterialDetail{}).All(ctx, m.r, model.AllCourseMaterialDetail{
		CourseMaterialID: param.CourseMaterialID,
		SearchBy:         param.SearchBy,
		SearchValue:      param.SearchValue,
		OrderBy:          param.OrderBy,
		OrderDir:         param.OrderDir,
		Offset:           param.Offset,
		Limit:            param.Limit,
	})
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on query all course material detail"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no Course material detail found"
		}

		return []model.CourseMaterialDetailResponse{}, NewErrorWrap(err, m.Name, "all/course_material_detail",
			message, status)
	}

	for _, each := range data {
		allResp = append(allResp, each.Response())
	}

	return allResp, nil

}

func (m CourseMaterialDetailModule) Add(ctx context.Context, param AddCourseMaterialDetailParam) (model.CourseMaterialDetailResponse, *Error) {
	courseMaterialDetail := &model.CourseMaterialDetail{
		CourseMaterialID: param.CourseMaterialID,
		Position:         param.Position,
		Title:            param.Title,
		TypeMaterial:     param.TypeMaterial,
		Content:          param.Content,
		ImageURL:         param.ImageURL,
	}

	id, err := courseMaterialDetail.Add(ctx, m.r)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on add course material detail"

		return model.CourseMaterialDetailResponse{}, NewErrorWrap(err, m.Name, "add/course_material_detail",
			message, status)
	}

	courseMaterialDetail.ID = id

	return courseMaterialDetail.Response(), nil
}
