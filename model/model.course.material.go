package model

import (
	"context"

	"github.com/renosyah/AyoLesPortal/util"
	uuid "github.com/satori/go.uuid"
)

type (
	CourseMaterial struct {
		ID            uuid.UUID `json:"id"`
		CourseID      uuid.UUID `json:"course_id"`
		MaterialIndex int32     `json:"material_index"`
		Title         string    `json:"title"`
	}

	CourseMaterialResponse struct {
		ID            uuid.UUID `json:"id"`
		CourseID      uuid.UUID `json:"course_id"`
		MaterialIndex int32     `json:"material_index"`
		Title         string    `json:"title"`
	}

	AllCourseMaterial struct {
		CourseID    uuid.UUID `json:"course_id"`
		SearchBy    string    `json:"search_by"`
		SearchValue string    `json:"search_value"`
		OrderBy     string    `json:"order_by"`
		OrderDir    string    `json:"order_dir"`
		Offset      int       `json:"offset"`
		Limit       int       `json:"limit"`
	}
)

func (c *CourseMaterial) Response() CourseMaterialResponse {
	return CourseMaterialResponse{
		ID:            c.ID,
		CourseID:      c.CourseID,
		MaterialIndex: c.MaterialIndex,
		Title:         c.Title,
	}
}

func (c *CourseMaterial) Add(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	return c.ID, nil
}

func (c *CourseMaterial) One(ctx context.Context, r *util.PostData) (*CourseMaterial, error) {
	one := &CourseMaterial{}
	return one, nil
}

func (c *CourseMaterial) All(ctx context.Context, r *util.PostData, param AllCourseMaterial) ([]*CourseMaterial, error) {
	all := []*CourseMaterial{}
	return all, nil
}

func (c *CourseMaterial) Update(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}

func (c *CourseMaterial) Delete(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}
