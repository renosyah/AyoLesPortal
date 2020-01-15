package model

import (
	"context"

	"github.com/renosyah/AyoLesPortal/util"
	uuid "github.com/satori/go.uuid"
)

type (
	CourseExam struct {
		ID        uuid.UUID           `json:"id"`
		CourseID  uuid.UUID           `json:"course_id"`
		TypeExam  int32               `json:"type_exam"`
		ExamIndex int32               `json:"exam_index"`
		Text      string              `json:"text"`
		ImageURL  string              `json:"image_url"`
		Answers   []*CourseExamAnswer `json:"answers"`
	}

	CourseExamResponse struct {
		ID        uuid.UUID                  `json:"id"`
		CourseID  uuid.UUID                  `json:"course_id"`
		TypeExam  int32                      `json:"type_exam"`
		ExamIndex int32                      `json:"exam_index"`
		Text      string                     `json:"text"`
		ImageURL  string                     `json:"image_url"`
		Answers   []CourseExamAnswerResponse `json:"answers"`
	}

	AllCourseExam struct {
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

func (c *CourseExam) Response() CourseExamResponse {
	answers := []CourseExamAnswerResponse{}
	for _, one := range c.Answers {
		answers = append(answers, one.Response())
	}
	return CourseExamResponse{
		ID:        c.ID,
		CourseID:  c.CourseID,
		TypeExam:  c.TypeExam,
		ExamIndex: c.ExamIndex,
		Text:      c.Text,
		ImageURL:  c.ImageURL,
		Answers:   answers,
	}
}

func (c *CourseExam) Add(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	return c.ID, nil
}

func (c *CourseExam) One(ctx context.Context, r *util.PostData, LimitAnswer int) (*CourseExam, error) {
	one := &CourseExam{}
	return one, nil
}

func (c *CourseExam) All(ctx context.Context, r *util.PostData, param AllCourseExam) ([]*CourseExam, error) {
	all := []*CourseExam{}
	return all, nil
}

func (c *CourseExam) Update(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}

func (c *CourseExam) Delete(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}
