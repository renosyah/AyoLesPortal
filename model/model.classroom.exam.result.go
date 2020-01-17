package model

import (
	"context"
	"fmt"

	"github.com/renosyah/AyoLesPortal/util"
	uuid "github.com/satori/go.uuid"
)

type (
	ClassRoomExamResult struct {
		CourseExamID    uuid.UUID           `json:"course_exam_id"`
		CourseID        uuid.UUID           `json:"course_id"`
		ClassRoomID     uuid.UUID           `json:"classroom_id"`
		StudentAnswerID uuid.UUID           `json:"student_answer_id"`
		ValidAnswerID   uuid.UUID           `json:"valid_answer_id"`
		TypeExam        int32               `json:"type_exam"`
		ExamIndex       int32               `json:"exam_index"`
		Text            string              `json:"text"`
		ImageURL        string              `json:"image_url"`
		Answers         []*CourseExamAnswer `json:"answers"`
	}

	ClassRoomExamResultResponse struct {
		CourseExamID    uuid.UUID                  `json:"course_exam_id"`
		CourseID        uuid.UUID                  `json:"course_id"`
		ClassRoomID     uuid.UUID                  `json:"classroom_id"`
		StudentAnswerID uuid.UUID                  `json:"student_answer_id"`
		ValidAnswerID   uuid.UUID                  `json:"valid_answer_id"`
		TypeExam        int32                      `json:"type_exam"`
		ExamIndex       int32                      `json:"exam_index"`
		Text            string                     `json:"text"`
		ImageURL        string                     `json:"image_url"`
		Answers         []CourseExamAnswerResponse `json:"answers"`
	}

	AllClassRoomExamResult struct {
		ClassRoomID uuid.UUID `json:"classroom_id"`
		SearchBy    string    `json:"search_by"`
		SearchValue string    `json:"search_value"`
		OrderBy     string    `json:"order_by"`
		OrderDir    string    `json:"order_dir"`
		Offset      int       `json:"offset"`
		Limit       int       `json:"limit"`
		LimitAnswer int       `json:"limit_answer"`
	}
)

func (c *ClassRoomExamResult) Response() ClassRoomExamResultResponse {
	answers := []CourseExamAnswerResponse{}
	for _, one := range c.Answers {
		answers = append(answers, one.Response())
	}
	return ClassRoomExamResultResponse{
		CourseExamID:    c.CourseExamID,
		CourseID:        c.CourseID,
		ClassRoomID:     c.ClassRoomID,
		StudentAnswerID: c.StudentAnswerID,
		ValidAnswerID:   c.ValidAnswerID,
		TypeExam:        c.TypeExam,
		ExamIndex:       c.ExamIndex,
		Text:            c.Text,
		ImageURL:        c.ImageURL,
		Answers:         answers,
	}
}

func (c *ClassRoomExamResult) One(ctx context.Context, r *util.PostData, LimitAnswer int) (*ClassRoomExamResult, error) {
	classroomExamResultDetail := struct {
		ClassroomExamResultDetail *ClassRoomExamResult `json:"classroom_exam_result_detail"`
	}{
		ClassroomExamResultDetail: &ClassRoomExamResult{},
	}

	query := `query {
		classroom_exam_result_detail(
			course_exam_id: "%s",
			course_id: "%s",
			limit_answer: %d
		)
		{
			course_exam_id,
			course_id,
			classroom_id,
			student_answer_id,
			right_answer_id,
			type_exam,
			exam_index,
			text,
			image_url
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

	resp, err := r.Send(fmt.Sprintf(query, c.CourseExamID, c.CourseID, LimitAnswer))
	if err != nil {
		return classroomExamResultDetail.ClassroomExamResultDetail, err
	}

	err = resp.Body.Error()
	if err != nil {
		return classroomExamResultDetail.ClassroomExamResultDetail, err
	}

	err = resp.Body.ConvertData(&classroomExamResultDetail)
	if err != nil {
		return classroomExamResultDetail.ClassroomExamResultDetail, err
	}
	return classroomExamResultDetail.ClassroomExamResultDetail, nil
}

func (c *ClassRoomExamResult) All(ctx context.Context, r *util.PostData, param AllClassRoomExamResult) ([]*ClassRoomExamResult, error) {
	classroomExamResultList := struct {
		ClassroomExamResultList []*ClassRoomExamResult `json:"classroom_exam_result_list"`
	}{
		ClassroomExamResultList: []*ClassRoomExamResult{},
	}

	query := `query {
		classroom_exam_result_list(
			classroom_id : "%s",
			search_by:"%s",
			search_value:"%s",
			order_by:"%s",
			order_dir:"%s",
			offset:%d,
			limit:%d,
			limit_answer:%d
		)
		{
			course_exam_id,
			course_id,
			classroom_id,
			student_answer_id,
			right_answer_id,
			type_exam,
			exam_index,
			text,
			image_url
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
		param.ClassRoomID, param.SearchBy, param.SearchValue, param.OrderBy, param.OrderDir, param.Offset, param.Limit, param.LimitAnswer))
	if err != nil {
		return classroomExamResultList.ClassroomExamResultList, err
	}

	err = resp.Body.Error()
	if err != nil {
		return classroomExamResultList.ClassroomExamResultList, err
	}

	err = resp.Body.ConvertData(&classroomExamResultList)
	if err != nil {
		return classroomExamResultList.ClassroomExamResultList, err
	}
	return classroomExamResultList.ClassroomExamResultList, nil
}

// ITS DOESNOT HAVE TABLE
// THIS MODEL VALUE RESULT FROM
// QUERY JOIN
// NO UPDATE
// NO DELETE
