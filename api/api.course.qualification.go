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
	CourseQualificationModule struct {
		r    *util.PostData
		Name string
	}

	AddCourseQualificationParam struct {
		CourseID            uuid.UUID `json:"course_id"`
		CourseLevel         string    `json:"course_level"`
		MinScore            int32     `json:"min_score"`
		CourseMaterialTotal int32     `json:"course_material_total"`
		CourseExamTotal     int32     `json:"course_exam_total"`
	}

	OneCourseQualificationParam struct {
		ID       uuid.UUID `json:"id"`
		CourseID uuid.UUID `json:"course_id"`
	}
)

func NewCourseQualificationModule(r *util.PostData) *CourseQualificationModule {
	return &CourseQualificationModule{
		r:    r,
		Name: "module/course_qualification_module",
	}
}

func (m CourseQualificationModule) Add(ctx context.Context, param AddCourseQualificationParam) (model.CourseQualificationResponse, *Error) {
	courseQualification := &model.CourseQualification{
		CourseID:            param.CourseID,
		CourseLevel:         param.CourseLevel,
		MinScore:            param.MinScore,
		CourseMaterialTotal: param.CourseMaterialTotal,
		CourseExamTotal:     param.CourseExamTotal,
	}

	id, err := courseQualification.Add(ctx, m.r)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on add course qualification"

		return model.CourseQualificationResponse{}, NewErrorWrap(err, m.Name, "add/course_qualification",
			message, status)
	}

	courseQualification.ID = id

	return courseQualification.Response(), nil
}

func (m CourseQualificationModule) One(ctx context.Context, param OneCourseQualificationParam) (model.CourseQualificationResponse, *Error) {
	courseQualification := &model.CourseQualification{
		ID:       param.ID,
		CourseID: param.CourseID,
	}

	data, err := courseQualification.One(ctx, m.r)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on get one course qualification"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no course qualificationfound"
		}

		return model.CourseQualificationResponse{}, NewErrorWrap(err, m.Name, "one/course_qualification",
			message, status)
	}

	return data.Response(), nil
}
