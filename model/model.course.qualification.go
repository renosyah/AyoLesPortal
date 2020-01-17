package model

import (
	"context"
	"fmt"

	"github.com/renosyah/AyoLesPortal/util"
	uuid "github.com/satori/go.uuid"
)

type (
	CourseQualification struct {
		ID                  uuid.UUID `json:"id"`
		CourseID            uuid.UUID `json:"course_id"`
		CourseLevel         string    `json:"course_level"`
		MinScore            int32     `json:"min_score"`
		CourseMaterialTotal int32     `json:"course_material_total"`
		CourseExamTotal     int32     `json:"course_exam_total"`
	}
	CourseQualificationResponse struct {
		ID                  uuid.UUID `json:"id"`
		CourseID            uuid.UUID `json:"course_id"`
		CourseLevel         string    `json:"course_level"`
		MinScore            int32     `json:"min_score"`
		CourseMaterialTotal int32     `json:"course_material_total"`
		CourseExamTotal     int32     `json:"course_exam_total"`
	}
)

func (c *CourseQualification) Response() CourseQualificationResponse {
	return CourseQualificationResponse{
		ID:                  c.ID,
		CourseID:            c.CourseID,
		CourseLevel:         c.CourseLevel,
		MinScore:            c.MinScore,
		CourseMaterialTotal: c.CourseMaterialTotal,
		CourseExamTotal:     c.CourseExamTotal,
	}
}

func (c *CourseQualification) Add(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	courseQualificationRegister := struct {
		CoursequalificationRegister CourseQualification `json:"course_qualification_register"`
	}{
		CoursequalificationRegister: CourseQualification{},
	}

	query := `mutation {
		course_qualification_register(
			course_id:"%s",
			course_level:"%s",
			min_score:%d,
			course_material_total:%d,
			course_exam_total:%d
		)
		{
			id,
			course_id,
			course_level,
			min_score,
			course_material_total,
			course_exam_total
		}
	}`

	resp, err := r.Send(fmt.Sprintf(query, c.CourseID, c.CourseLevel, c.MinScore, c.CourseMaterialTotal, c.CourseExamTotal))
	if err != nil {
		return c.ID, err
	}

	err = resp.Body.Error()
	if err != nil {
		return c.ID, err
	}

	err = resp.Body.ConvertData(&courseQualificationRegister)
	if err != nil {
		return c.ID, err
	}

	c.ID = courseQualificationRegister.CoursequalificationRegister.ID

	return c.ID, nil
}

func (c *CourseQualification) One(ctx context.Context, r *util.PostData) (*CourseQualification, error) {
	courseQualificationDetail := struct {
		CourseQualificationDetail *CourseQualification `json:"course_qualification_detail"`
	}{
		CourseQualificationDetail: &CourseQualification{},
	}

	query := `query {
		course_qualification_detail(
			id: "%s",
			course_id: "%s",
		)
		{
			id,
			course_id,
			course_level,
			min_score,
			course_material_total,
			course_exam_total
		}
	}`

	resp, err := r.Send(fmt.Sprintf(query, c.ID, c.CourseID))
	if err != nil {
		return courseQualificationDetail.CourseQualificationDetail, err
	}

	err = resp.Body.Error()
	if err != nil {
		return courseQualificationDetail.CourseQualificationDetail, err
	}

	err = resp.Body.ConvertData(&courseQualificationDetail)
	if err != nil {
		return courseQualificationDetail.CourseQualificationDetail, err
	}
	return courseQualificationDetail.CourseQualificationDetail, nil
}

func (c *CourseQualification) Update(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}

func (c *CourseQualification) Delete(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}
