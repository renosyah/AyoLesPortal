package model

import (
	"context"

	"github.com/renosyah/AyoLesPortal/util"
	uuid "github.com/satori/go.uuid"
)

type (
	CourseExamSolution struct {
		ID                 uuid.UUID `json:"id"`
		CourseExamID       uuid.UUID `json:"course_exam_id"`
		CourseExamAnswerID uuid.UUID `json:"course_exam_answer_id"`
	}

	CourseExamSolutionResponse struct {
		ID                 uuid.UUID `json:"id"`
		CourseExamID       uuid.UUID `json:"course_exam_id"`
		CourseExamAnswerID uuid.UUID `json:"course_exam_answer_id"`
	}
	AllCourseExamSolution struct {
		CourseExamID uuid.UUID `json:"course_exam_id"`
		OrderBy      string    `json:"order_by"`
		OrderDir     string    `json:"order_dir"`
		Offset       int       `json:"offset"`
		Limit        int       `json:"limit"`
	}
)

func (c *CourseExamSolution) Response() CourseExamSolutionResponse {
	return CourseExamSolutionResponse{
		ID:                 c.ID,
		CourseExamID:       c.CourseExamID,
		CourseExamAnswerID: c.CourseExamAnswerID,
	}
}

func (c *CourseExamSolution) Add(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	return c.ID, nil
}

func (c *CourseExamSolution) One(ctx context.Context, r *util.PostData) (*CourseExamSolution, error) {
	one := &CourseExamSolution{}
	return one, nil
}

func (c *CourseExamSolution) All(ctx context.Context, r *util.PostData, param AllCourseExamSolution) ([]*CourseExamSolution, error) {
	all := []*CourseExamSolution{}
	return all, nil
}

func (c *CourseExamSolution) Update(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}

func (c *CourseExamSolution) Delete(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}
