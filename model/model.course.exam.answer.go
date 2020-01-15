package model

import (
	"context"

	"github.com/renosyah/AyoLesPortal/util"
	uuid "github.com/satori/go.uuid"
)

type (
	CourseExamAnswer struct {
		ID           uuid.UUID `json:"id"`
		CourseExamID uuid.UUID `json:"course_exam_id"`
		TypeAnswer   int32     `json:"type_answer"`
		Label        string    `json:"label"`
		Text         string    `json:"text"`
		ImageURL     string    `json:"image_url"`
	}

	CourseExamAnswerResponse struct {
		ID           uuid.UUID `json:"id"`
		CourseExamID uuid.UUID `json:"course_exam_id"`
		TypeAnswer   int32     `json:"type_answer"`
		Label        string    `json:"label"`
		Text         string    `json:"text"`
		ImageURL     string    `json:"image_url"`
	}

	AllCourseExamAnswer struct {
		CourseExamID uuid.UUID `json:"course_exam_id"`
		SearchBy     string    `json:"search_by"`
		SearchValue  string    `json:"search_value"`
		OrderBy      string    `json:"order_by"`
		OrderDir     string    `json:"order_dir"`
		Offset       int       `json:"offset"`
		Limit        int       `json:"limit"`
	}
)

func (c *CourseExamAnswer) Response() CourseExamAnswerResponse {
	return CourseExamAnswerResponse{
		ID:           c.ID,
		CourseExamID: c.CourseExamID,
		TypeAnswer:   c.TypeAnswer,
		Label:        c.Label,
		Text:         c.Text,
		ImageURL:     c.ImageURL,
	}
}

func (c *CourseExamAnswer) Add(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	return c.ID, nil
}

func (c *CourseExamAnswer) One(ctx context.Context, r *util.PostData) (*CourseExamAnswer, error) {
	one := &CourseExamAnswer{}
	return one, nil
}

func (c *CourseExamAnswer) All(ctx context.Context, r *util.PostData, param AllCourseExamAnswer) ([]*CourseExamAnswer, error) {
	all := []*CourseExamAnswer{}
	return all, nil
}

func (c *CourseExamAnswer) AllById(ctx context.Context, r *util.PostData, LimitAnswer int) ([]*CourseExamAnswer, error) {
	all := []*CourseExamAnswer{}
	return all, nil
}

func (c *CourseExamAnswer) Update(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}

func (c *CourseExamAnswer) Delete(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}
