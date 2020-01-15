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
	CourseExamModule struct {
		r    *util.PostData
		Name string
	}

	AddCourseExamParam struct {
		CourseID  uuid.UUID `json:"course_id"`
		TypeExam  int32     `json:"type_exam"`
		ExamIndex int32     `json:"exam_index"`
		Text      string    `json:"text"`
		ImageURL  string    `json:"image_url"`
	}

	OneCourseExamParam struct {
		ID          uuid.UUID `json:"id"`
		LimitAnswer int       `json:"limit_answer"`
	}

	AllCourseExamParam struct {
		CourseID    uuid.UUID `json:"course_id"`
		SearchBy    string    `json:"search_by"`
		SearchValue string    `json:"search_value"`
		OrderBy     string    `json:"order_by"`
		OrderDir    string    `json:"order_dir"`
		Offset      int       `json:"offset"`
		Limit       int       `json:"limit"`
		LimitAnswer int       `json:"limit_answer"`
	}
)

func NewCourseExamModule(r *util.PostData) *CourseExamModule {
	return &CourseExamModule{
		r:    r,
		Name: "module/course_exam",
	}
}

func (m CourseExamModule) All(ctx context.Context, param AllCourseExamParam) ([]model.CourseExamResponse, *Error) {
	var allResp []model.CourseExamResponse

	data, err := (&model.CourseExam{}).All(ctx, m.r, model.AllCourseExam{
		CourseID:    param.CourseID,
		SearchBy:    param.SearchBy,
		SearchValue: param.SearchValue,
		OrderBy:     param.OrderBy,
		OrderDir:    param.OrderDir,
		Offset:      param.Offset,
		Limit:       param.Limit,
		LimitAnswer: param.LimitAnswer,
	})
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on query all course exam"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no Course exam found"
		}

		return []model.CourseExamResponse{}, NewErrorWrap(err, m.Name, "all/course_examp",
			message, status)
	}

	for _, each := range data {
		allResp = append(allResp, each.Response())
	}

	return allResp, nil

}
func (m CourseExamModule) Add(ctx context.Context, param AddCourseExamParam) (model.CourseExamResponse, *Error) {
	courseExam := &model.CourseExam{
		CourseID:  param.CourseID,
		TypeExam:  param.TypeExam,
		ExamIndex: param.ExamIndex,
		Text:      param.Text,
		ImageURL:  param.ImageURL,
	}

	id, err := courseExam.Add(ctx, m.r)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on add course exam"

		return model.CourseExamResponse{}, NewErrorWrap(err, m.Name, "add/course_exam",
			message, status)
	}

	courseExam.ID = id

	return courseExam.Response(), nil
}

func (m CourseExamModule) One(ctx context.Context, param OneCourseExamParam) (model.CourseExamResponse, *Error) {
	courseExam := &model.CourseExam{
		ID: param.ID,
	}

	data, err := courseExam.One(ctx, m.r, param.LimitAnswer)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on get one course exam"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no course exam found"
		}

		return model.CourseExamResponse{}, NewErrorWrap(err, m.Name, "one/course_exam",
			message, status)
	}

	return data.Response(), nil
}
