package model

import (
	"context"
	"fmt"

	"github.com/renosyah/AyoLesPortal/util"
	uuid "github.com/satori/go.uuid"
)

type (
	ClassRoomExamProgress struct {
		ID                 uuid.UUID `json:"id"`
		ClassroomID        uuid.UUID `json:"classroom_id"`
		CourseExamID       uuid.UUID `json:"course_exam_id"`
		CourseExamAnswerID uuid.UUID `json:"course_exam_answer_id"`
	}
	ClassRoomExamProgressResponse struct {
		ID                 uuid.UUID `json:"id"`
		ClassroomID        uuid.UUID `json:"classroom_id"`
		CourseExamID       uuid.UUID `json:"course_exam_id"`
		CourseExamAnswerID uuid.UUID `json:"course_exam_answer_id"`
	}
	AllClassRoomExamProgress struct {
		ClassroomID uuid.UUID `json:"classroom_id"`
		OrderBy     string    `json:"order_by"`
		OrderDir    string    `json:"order_dir"`
		Offset      int       `json:"offset"`
		Limit       int       `json:"limit"`
	}
)

func (c *ClassRoomExamProgress) Response() ClassRoomExamProgressResponse {
	return ClassRoomExamProgressResponse{
		ID:                 c.ID,
		ClassroomID:        c.ClassroomID,
		CourseExamID:       c.CourseExamID,
		CourseExamAnswerID: c.CourseExamAnswerID,
	}
}

func (c *ClassRoomExamProgress) Add(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	classroomExamProgressRegister := struct {
		ClassroomExamProgressRegister ClassRoomExamProgress `json:"classroom_exam_progress_register"`
	}{
		ClassroomExamProgressRegister: ClassRoomExamProgress{},
	}

	query := `mutation {
		classroom_exam_progress_register(
			classroom_id : "%s",
			course_exam_id : "%s",
			course_exam_answer_id : "%s"
		)
		{
			id,
			classroom_id,
			course_exam_id,
			course_exam_answer_id
		 }
	}`

	resp, err := r.Send(fmt.Sprintf(query, c.ClassroomID, c.CourseExamID, c.CourseExamAnswerID))
	if err != nil {
		return c.ID, err
	}

	err = resp.Body.Error()
	if err != nil {
		return c.ID, err
	}

	err = resp.Body.ConvertData(&classroomExamProgressRegister)
	if err != nil {
		return c.ID, err
	}

	c.ID = classroomExamProgressRegister.ClassroomExamProgressRegister.ID

	return c.ID, nil
}

func (c *ClassRoomExamProgress) One(ctx context.Context, r *util.PostData) (*ClassRoomExamProgress, error) {
	classroomExamProgressDetail := struct {
		ClassroomExamProgressDetail *ClassRoomExamProgress `json:"classroom_exam_progress_detail"`
	}{
		ClassroomExamProgressDetail: &ClassRoomExamProgress{},
	}

	query := `query {
		classroom_exam_progress_detail(
			id: "%s"
		)
		{
			id,
			classroom_id,
			course_exam_id,
			course_exam_answer_id
		 }
	}`

	resp, err := r.Send(fmt.Sprintf(query, c.ID))
	if err != nil {
		return classroomExamProgressDetail.ClassroomExamProgressDetail, err
	}

	err = resp.Body.Error()
	if err != nil {
		return classroomExamProgressDetail.ClassroomExamProgressDetail, err
	}

	err = resp.Body.ConvertData(&classroomExamProgressDetail)
	if err != nil {
		return classroomExamProgressDetail.ClassroomExamProgressDetail, err
	}
	return classroomExamProgressDetail.ClassroomExamProgressDetail, nil
}

func (c *ClassRoomExamProgress) All(ctx context.Context, r *util.PostData, param AllClassRoomExamProgress) ([]*ClassRoomExamProgress, error) {
	classroomExamProgressList := struct {
		ClassroomExamProgressList []*ClassRoomExamProgress `json:"classroom_exam_progress_list"`
	}{
		ClassroomExamProgressList: []*ClassRoomExamProgress{},
	}

	query := `query {
		classroom_exam_progress_list(
			classroom_id:"%s",
			order_by:"%s",
			order_dir:"%s",
			offset:%d,
			limit:%d
		)
		{
			id,
			classroom_id,
			course_exam_id,
			course_exam_answer_id
		 }
	}`

	resp, err := r.Send(fmt.Sprintf(query,
		param.ClassroomID, param.OrderBy, param.OrderDir, param.Offset, param.Limit))
	if err != nil {
		return classroomExamProgressList.ClassroomExamProgressList, err
	}

	err = resp.Body.Error()
	if err != nil {
		return classroomExamProgressList.ClassroomExamProgressList, err
	}

	err = resp.Body.ConvertData(&classroomExamProgressList)
	if err != nil {
		return classroomExamProgressList.ClassroomExamProgressList, err
	}
	return classroomExamProgressList.ClassroomExamProgressList, nil
}

func (c *ClassRoomExamProgress) Update(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}

func (c *ClassRoomExamProgress) Delete(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}
