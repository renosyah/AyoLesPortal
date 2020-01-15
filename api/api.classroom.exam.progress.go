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
	ClassRoomExamProgressModule struct {
		r    *util.PostData
		Name string
	}
	AddClassRoomExamParam struct {
		ClassroomID        uuid.UUID `json:"classroom_id"`
		CourseExamID       uuid.UUID `json:"course_exam_id"`
		CourseExamAnswerID uuid.UUID `json:"course_exam_answer_id"`
	}

	OneClassRoomExamProgressParam struct {
		ID uuid.UUID `json:"id"`
	}

	DeleteClassRoomExamProgressParam struct {
		ClassroomID uuid.UUID `json:"classroom_id"`
	}

	AllClassRoomExamProgressParam struct {
		ClassroomID uuid.UUID `json:"classroom_id"`
		OrderBy     string    `json:"order_by"`
		OrderDir    string    `json:"order_dir"`
		Offset      int       `json:"offset"`
		Limit       int       `json:"limit"`
	}
)

func NewClassRoomExamProgressModule(r *util.PostData) *ClassRoomExamProgressModule {
	return &ClassRoomExamProgressModule{
		r:    r,
		Name: "module/classroom_exam_progress",
	}
}

func (m ClassRoomExamProgressModule) All(ctx context.Context, param AllClassRoomExamProgressParam) ([]model.ClassRoomExamProgressResponse, *Error) {
	var allResp []model.ClassRoomExamProgressResponse

	data, err := (&model.ClassRoomExamProgress{}).All(ctx, m.r, model.AllClassRoomExamProgress{
		ClassroomID: param.ClassroomID,
		OrderBy:     param.OrderBy,
		OrderDir:    param.OrderDir,
		Offset:      param.Offset,
		Limit:       param.Limit,
	})
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on query all classRoom exam progress"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no classroom exam found"
		}

		return []model.ClassRoomExamProgressResponse{}, NewErrorWrap(err, m.Name, "all/classroom_exam_progress_module",
			message, status)
	}

	for _, each := range data {
		allResp = append(allResp, each.Response())
	}

	return allResp, nil

}

func (m ClassRoomExamProgressModule) Add(ctx context.Context, param AddClassRoomExamParam) (model.ClassRoomExamProgressResponse, *Error) {
	classroomExamProgress := &model.ClassRoomExamProgress{
		ClassroomID:        param.ClassroomID,
		CourseExamID:       param.CourseExamID,
		CourseExamAnswerID: param.CourseExamAnswerID,
	}

	id, err := classroomExamProgress.Add(ctx, m.r)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on add classRoom exam progress"

		return model.ClassRoomExamProgressResponse{}, NewErrorWrap(err, m.Name, "add/classroom_exam_progress_module",
			message, status)
	}

	classroomExamProgress.ID = id

	return classroomExamProgress.Response(), nil
}

func (m ClassRoomExamProgressModule) One(ctx context.Context, param OneClassRoomExamProgressParam) (model.ClassRoomExamProgressResponse, *Error) {
	courseExamProgress := &model.ClassRoomExamProgress{
		ID: param.ID,
	}

	data, err := courseExamProgress.One(ctx, m.r)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on get one all classRoom exam progress"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no all classRoom exam progress found"
		}

		return model.ClassRoomExamProgressResponse{}, NewErrorWrap(err, m.Name, "one/classroom_exam_progress_module",
			message, status)
	}

	return data.Response(), nil
}

func (m ClassRoomExamProgressModule) Delete(ctx context.Context, param DeleteClassRoomExamProgressParam) (model.ClassRoomExamProgressResponse, *Error) {
	courseExamProgress := &model.ClassRoomExamProgress{
		ClassroomID: param.ClassroomID,
	}

	_, err := courseExamProgress.Delete(ctx, m.r)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on delete all classRoom exam progress"

		return model.ClassRoomExamProgressResponse{}, NewErrorWrap(err, m.Name, "one/classroom_exam_progress_module",
			message, status)
	}

	return (&model.ClassRoomExamProgress{ClassroomID: param.ClassroomID}).Response(), nil
}
