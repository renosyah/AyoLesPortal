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
	ClassRoomQualificationModule struct {
		r    *util.PostData
		Name string
	}
	OneClassRoomQualificationParam struct {
		ClassRoomID uuid.UUID `json:"classroom_id"`
	}
)

func NewClassRoomQualificationModule(r *util.PostData) *ClassRoomQualificationModule {
	return &ClassRoomQualificationModule{
		r:    r,
		Name: "module/classroom_qualification",
	}
}

func (m ClassRoomQualificationModule) One(ctx context.Context, param OneClassRoomQualificationParam) (model.ClassRoomQualificationResponse, *Error) {
	classRoomQualification := &model.ClassRoomQualification{
		ClassRoomID: param.ClassRoomID,
	}

	data, err := classRoomQualification.One(ctx, m.r)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on get one classRoom qualification"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no one classRoom qualification found"
		}

		return model.ClassRoomQualificationResponse{}, NewErrorWrap(err, m.Name, "one/course_progress_qualification",
			message, status)
	}

	return data.Response(), nil
}
