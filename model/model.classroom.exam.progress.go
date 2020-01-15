package model

import (
	"context"

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
	return c.ID, nil
}

func (c *ClassRoomExamProgress) One(ctx context.Context, r *util.PostData) (*ClassRoomExamProgress, error) {
	one := &ClassRoomExamProgress{}
	return one, nil
}

func (c *ClassRoomExamProgress) All(ctx context.Context, r *util.PostData, param AllClassRoomExamProgress) ([]*ClassRoomExamProgress, error) {
	all := []*ClassRoomExamProgress{}
	return all, nil
}

func (c *ClassRoomExamProgress) Update(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}

func (c *ClassRoomExamProgress) Delete(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}
