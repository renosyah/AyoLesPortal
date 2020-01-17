package model

import (
	"context"
	"fmt"

	"github.com/renosyah/AyoLesPortal/util"
	uuid "github.com/satori/go.uuid"
)

const (
	STATUS_NO_PROGRESS   = 0
	STATUS_PASS_EXAM     = 1
	STATUS_NOT_PASS_EXAM = 2
)

type (
	ClassRoomQualification struct {
		ClassRoomID         uuid.UUID            `json:"classroom_id"`
		CourseQualification *CourseQualification `json:"course_qualification"`
		TotalScore          int32                `json:"total_score"`
		Status              int32                `json:"status"`
	}

	ClassRoomQualificationResponse struct {
		ClassRoomID         uuid.UUID                   `json:"classroom_id"`
		CourseQualification CourseQualificationResponse `json:"course_qualification"`
		TotalScore          int32                       `json:"total_score"`
		Status              int32                       `json:"status"`
	}
)

func (c *ClassRoomQualification) Response() ClassRoomQualificationResponse {
	return ClassRoomQualificationResponse{
		ClassRoomID:         c.ClassRoomID,
		CourseQualification: c.CourseQualification.Response(),
		TotalScore:          c.TotalScore,
		Status:              c.Status,
	}
}

func (c *ClassRoomQualification) One(ctx context.Context, r *util.PostData) (*ClassRoomQualification, error) {
	classQualificationDetail := struct {
		ClassQualificationDetail *ClassRoomQualification `json:"class_qualification_detail"`
	}{
		ClassQualificationDetail: &ClassRoomQualification{},
	}

	query := `query {
		class_qualification_detail(
			classroom_id: "%s"
		)
		{
			classroom_id,
			total_score,
			status,
			course_qualification {
				id,
				course_id,
				course_level,
				min_score,
				course_material_total,
				course_exam_total
			}
		}
	}`

	resp, err := r.Send(fmt.Sprintf(query, c.ClassRoomID))
	if err != nil {
		return classQualificationDetail.ClassQualificationDetail, err
	}

	err = resp.Body.Error()
	if err != nil {
		return classQualificationDetail.ClassQualificationDetail, err
	}

	err = resp.Body.ConvertData(&classQualificationDetail)
	if err != nil {
		return classQualificationDetail.ClassQualificationDetail, err
	}
	return classQualificationDetail.ClassQualificationDetail, nil
}

// ITS DOESNOT HAVE TABLE
// THIS MODEL VALUE RESULT FROM
// QUERY JOIN
// NO UPDATE
// NO DELETE
