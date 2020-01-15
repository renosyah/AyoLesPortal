package model

import (
	"context"

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
	one := &ClassRoomExamResult{}
	return one, nil
}

func (c *ClassRoomExamResult) All(ctx context.Context, r *util.PostData, param AllClassRoomExamResult) ([]*ClassRoomExamResult, error) {
	all := []*ClassRoomExamResult{}
	return all, nil
}

// ITS DOESNOT HAVE TABLE
// THIS MODEL VALUE RESULT FROM
// QUERY JOIN
// NO UPDATE
// NO DELETE
