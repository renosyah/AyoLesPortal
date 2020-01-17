package model

import (
	"context"
	"fmt"

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
	courseExamAnswerRegister := struct {
		CourseExamAnswerRegister CourseExamAnswer `json:"course_exam_answer_register"`
	}{
		CourseExamAnswerRegister: CourseExamAnswer{},
	}

	query := `mutation {
		course_exam_answer_register(
			course_exam_id : "%s",
			type_answer : %d,
			label : "%s",
			text : "%s",
			image_url : "%s"
		)
		{
			id,
			course_exam_id,
			type_answer,
			label,
			text,
			image_url
		}
	}`

	resp, err := r.Send(fmt.Sprintf(query,
		c.CourseExamID, c.TypeAnswer, c.Label, c.Text, c.ImageURL))
	if err != nil {
		return c.ID, err
	}

	err = resp.Body.Error()
	if err != nil {
		return c.ID, err
	}

	err = resp.Body.ConvertData(&courseExamAnswerRegister)
	if err != nil {
		return c.ID, err
	}

	c.ID = courseExamAnswerRegister.CourseExamAnswerRegister.ID

	return c.ID, nil
}

func (c *CourseExamAnswer) One(ctx context.Context, r *util.PostData) (*CourseExamAnswer, error) {
	courseExamAnswerDetail := struct {
		CourseExamAnswerDetail *CourseExamAnswer `json:"course_exam_answer_detail"`
	}{
		CourseExamAnswerDetail: &CourseExamAnswer{},
	}

	query := `query {
		course_exam_answer_detail(
			id: "%s"
		)
		{
			id,
			course_exam_id,
			type_answer,
			label,
			text,
			image_url
		}
	}`

	resp, err := r.Send(fmt.Sprintf(query, c.ID))
	if err != nil {
		return courseExamAnswerDetail.CourseExamAnswerDetail, err
	}

	err = resp.Body.Error()
	if err != nil {
		return courseExamAnswerDetail.CourseExamAnswerDetail, err
	}

	err = resp.Body.ConvertData(&courseExamAnswerDetail)
	if err != nil {
		return courseExamAnswerDetail.CourseExamAnswerDetail, err
	}
	return courseExamAnswerDetail.CourseExamAnswerDetail, nil
}

func (c *CourseExamAnswer) All(ctx context.Context, r *util.PostData, param AllCourseExamAnswer) ([]*CourseExamAnswer, error) {
	courseExamAnswerList := struct {
		CourseExamAnswerList []*CourseExamAnswer `json:"course_exam_answer_list"`
	}{
		CourseExamAnswerList: []*CourseExamAnswer{},
	}

	query := `query {
		course_exam_answer_list(
			course_exam_id:"%s",
			search_by:"%s",
			search_value:"%s",
			order_by:"%s",
			order_dir:"%s",
			offset:%d,
			limit:%d
		)
			id,
			course_exam_id,
			type_answer,
			label,
			text,
			image_url
		}
	}`

	resp, err := r.Send(fmt.Sprintf(query,
		param.CourseExamID, param.SearchBy, param.SearchValue, param.OrderBy, param.OrderDir, param.Offset, param.Limit))
	if err != nil {
		return courseExamAnswerList.CourseExamAnswerList, err
	}

	err = resp.Body.Error()
	if err != nil {
		return courseExamAnswerList.CourseExamAnswerList, err
	}

	err = resp.Body.ConvertData(&courseExamAnswerList)
	if err != nil {
		return courseExamAnswerList.CourseExamAnswerList, err
	}
	return courseExamAnswerList.CourseExamAnswerList, nil
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
