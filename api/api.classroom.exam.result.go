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
	ClassRoomExamResultModule struct {
		r    *util.PostData
		Name string
	}
	OneClassRoomExamResultParam struct {
		CourseExamID uuid.UUID `json:"course_exam_id"`
		CourseID     uuid.UUID `json:"course_id"`
		LimitAnswer  int       `json:"limit_answer"`
	}
	AllClassRoomExamResultParam struct {
		ClassRoomID uuid.UUID `json:"classroom_id"`
		SearchBy    string    `json:"search_by"`
		SearchValue string    `json:"search_value"`
		OrderBy     string    `json:"order_by"`
		OrderDir    string    `json:"order_dir"`
		Offset      int       `json:"offset"`
		Limit       int       `json:"limit"`
		LimitAnswer int       `json:"limit_answer"`
	}
)

func NewClassRoomExamResultModule(r *util.PostData) *ClassRoomExamResultModule {
	return &ClassRoomExamResultModule{
		r:    r,
		Name: "module/classroom_exam_result_module",
	}
}

func (m ClassRoomExamResultModule) All(ctx context.Context, param AllClassRoomExamResultParam) ([]model.ClassRoomExamResultResponse, *Error) {
	var allResp []model.ClassRoomExamResultResponse

	data, err := (&model.ClassRoomExamResult{}).All(ctx, m.r, model.AllClassRoomExamResult{
		ClassRoomID: param.ClassRoomID,
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
		message := "error on query all classRoom exam result"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no classroom exam result found"
		}

		return []model.ClassRoomExamResultResponse{}, NewErrorWrap(err, m.Name, "all/classroom_exam_result_module",
			message, status)
	}

	for _, each := range data {
		allResp = append(allResp, each.Response())
	}

	return allResp, nil

}

func (m ClassRoomExamResultModule) One(ctx context.Context, param OneClassRoomExamResultParam) (model.ClassRoomExamResultResponse, *Error) {
	courseExamResult := &model.ClassRoomExamResult{
		CourseExamID: param.CourseExamID,
		CourseID:     param.CourseID,
	}

	data, err := courseExamResult.One(ctx, m.r, param.LimitAnswer)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on get one all classRoom exam result"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no all classRoom exam result found"
		}

		return model.ClassRoomExamResultResponse{}, NewErrorWrap(err, m.Name, "one/classroom_exam_progress_module",
			message, status)
	}

	return data.Response(), nil
}
