package model

import (
	"context"
	"fmt"

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
	courseExamSolutionRegister := struct {
		CourseExamSolutionRegister CourseExamSolution `json:"course_exam_solution_register"`
	}{
		CourseExamSolutionRegister: CourseExamSolution{},
	}

	query := ` mutation {
		course_exam_solution_register(
			course_exam_id:"%s",
			course_exam_answer_id:"%s"
		)
		{
			id,
			course_exam_id,
			course_exam_answer_id
		}
	}`

	resp, err := r.Send(fmt.Sprintf(query, c.CourseExamID, c.CourseExamAnswerID))
	if err != nil {
		return c.ID, err
	}

	err = resp.Body.Error()
	if err != nil {
		return c.ID, err
	}

	err = resp.Body.ConvertData(&courseExamSolutionRegister)
	if err != nil {
		return c.ID, err
	}

	c.ID = courseExamSolutionRegister.CourseExamSolutionRegister.ID

	return c.ID, nil
}

func (c *CourseExamSolution) One(ctx context.Context, r *util.PostData) (*CourseExamSolution, error) {
	courseExamSolutionDetail := struct {
		CourseExamSolutionDetail *CourseExamSolution `json:"course_exam_solution_detail"`
	}{
		CourseExamSolutionDetail: &CourseExamSolution{},
	}

	query := `query {
		course_exam_solution_detail(
			id: "%s"
		)
		{
			id,
			course_exam_id,
			course_exam_answer_id
		}
	}`

	resp, err := r.Send(fmt.Sprintf(query, c.ID))
	if err != nil {
		return courseExamSolutionDetail.CourseExamSolutionDetail, err
	}

	err = resp.Body.Error()
	if err != nil {
		return courseExamSolutionDetail.CourseExamSolutionDetail, err
	}

	err = resp.Body.ConvertData(&courseExamSolutionDetail)
	if err != nil {
		return courseExamSolutionDetail.CourseExamSolutionDetail, err
	}
	return courseExamSolutionDetail.CourseExamSolutionDetail, nil
}

func (c *CourseExamSolution) All(ctx context.Context, r *util.PostData, param AllCourseExamSolution) ([]*CourseExamSolution, error) {
	courseExamSolutionList := struct {
		CourseExamSolutionList []*CourseExamSolution `json:"course_exam_solution_list"`
	}{
		CourseExamSolutionList: []*CourseExamSolution{},
	}

	query := `query {
		course_exam_solution_list(
			course_exam_id : "%s",
			order_by:"%s",
			order_dir:"%s",
			offset:%d,
			limit:%d
		)
			id,
			course_exam_id,
			course_exam_answer_id
		}
	}`

	resp, err := r.Send(fmt.Sprintf(query,
		param.CourseExamID, param.OrderBy, param.OrderDir, param.Offset, param.Limit))
	if err != nil {
		return courseExamSolutionList.CourseExamSolutionList, err
	}

	err = resp.Body.Error()
	if err != nil {
		return courseExamSolutionList.CourseExamSolutionList, err
	}

	err = resp.Body.ConvertData(&courseExamSolutionList)
	if err != nil {
		return courseExamSolutionList.CourseExamSolutionList, err
	}
	return courseExamSolutionList.CourseExamSolutionList, nil
}

func (c *CourseExamSolution) Update(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}

func (c *CourseExamSolution) Delete(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}
