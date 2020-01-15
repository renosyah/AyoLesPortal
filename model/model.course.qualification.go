package model

import (
	"context"

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
	return c.ID, nil
}

func (c *CourseQualification) One(ctx context.Context, r *util.PostData) (*CourseQualification, error) {
	one := &CourseQualification{}
	return one, nil
}

func (c *CourseQualification) Update(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}

func (c *CourseQualification) Delete(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}
