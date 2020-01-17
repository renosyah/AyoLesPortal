package model

import (
	"context"
	"fmt"

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
	courseExamRegister := struct {
		CourseExamRegister CourseExam `json:"course_exam_register"`
	}{
		CourseExamRegister: CourseExam{},
	}

	query := `mutation {
		course_exam_register(
			course_id : "%s",
			type_exam : "%s",
			exam_index : "%s",
			text : "%s",
			image_url : "%s"
		)
		{
			id,
			course_id,
			type_exam,
			exam_index,
			text,
			image_url,
			answers {
				id,
				course_exam_id,
				type_answer,
				label,
				text,
				image_url
			}
		}
	}`

	resp, err := r.Send(fmt.Sprintf(query, c.CourseID, c.TypeExam, c.ExamIndex, c.Text, c.ImageURL))
	if err != nil {
		return c.ID, err
	}

	err = resp.Body.Error()
	if err != nil {
		return c.ID, err
	}

	err = resp.Body.ConvertData(&courseExamRegister)
	if err != nil {
		return c.ID, err
	}

	c.ID = courseExamRegister.CourseExamRegister.ID

	return c.ID, nil
}

func (c *CourseExam) One(ctx context.Context, r *util.PostData, LimitAnswer int) (*CourseExam, error) {
	courseExamDetail := struct {
		CourseExamDetail *CourseExam `json:"course_exam_detail"`
	}{
		CourseExamDetail: &CourseExam{},
	}

	query := `query {
		course_exam_detail(
			id: "%s"
		)
		{
			id,
			course_id,
			type_exam,
			exam_index,
			text,
			image_url,
			answers {
				id,
				course_exam_id,
				type_answer,
				label,
				text,
				image_url
			}
		}
	}`

	resp, err := r.Send(fmt.Sprintf(query, c.ID))
	if err != nil {
		return courseExamDetail.CourseExamDetail, err
	}

	err = resp.Body.Error()
	if err != nil {
		return courseExamDetail.CourseExamDetail, err
	}

	err = resp.Body.ConvertData(&courseExamDetail)
	if err != nil {
		return courseExamDetail.CourseExamDetail, err
	}

	return courseExamDetail.CourseExamDetail, nil
}

func (c *CourseExam) All(ctx context.Context, r *util.PostData, param AllCourseExam) ([]*CourseExam, error) {
	courseExamList := struct {
		CourseExamList []*CourseExam `json:"course_exam_list"`
	}{
		CourseExamList: []*CourseExam{},
	}

	query := `query {
		course_exam_list(
			course_id:"%s",
			search_by:"%s",
			search_value:"%s",
			order_by:"%s",
			order_dir:"%s",
			offset:%d,
			limit:%d
			limit_answer:%d
		)
		{
			id,
			course_id,
			type_exam,
			exam_index,
			text,
			image_url,
			answers {
				id,
				course_exam_id,
				type_answer,
				label,
				text,
				image_url
			}
		}
	}`

	resp, err := r.Send(fmt.Sprintf(query,
		param.CourseID, param.SearchBy, param.SearchValue, param.OrderBy, param.OrderDir, param.Offset, param.Limit, param.LimitAnswer))
	if err != nil {
		return courseExamList.CourseExamList, err
	}

	err = resp.Body.Error()
	if err != nil {
		return courseExamList.CourseExamList, err
	}

	err = resp.Body.ConvertData(&courseExamList)
	if err != nil {
		return courseExamList.CourseExamList, err
	}

	return courseExamList.CourseExamList, nil
}

func (c *CourseExam) Update(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}

func (c *CourseExam) Delete(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}
