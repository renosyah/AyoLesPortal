package model

import (
	"context"

	"github.com/renosyah/AyoLesPortal/util"
	uuid "github.com/satori/go.uuid"
)

type (
	CourseMaterialDetail struct {
		ID               uuid.UUID `json:"id"`
		CourseMaterialID uuid.UUID `json:"course_material_id"`
		Position         int32     `json:"position"`
		Title            string    `json:"title"`
		TypeMaterial     int32     `json:"type_material"`
		Content          string    `json:"content"`
		ImageURL         string    `json:"image_url"`
	}

	CourseMaterialDetailResponse struct {
		ID               uuid.UUID `json:"id"`
		CourseMaterialID uuid.UUID `json:"course_material_id"`
		Position         int32     `json:"position"`
		Title            string    `json:"title"`
		TypeMaterial     int32     `json:"type_material"`
		Content          string    `json:"content"`
		ImageURL         string    `json:"image_url"`
	}

	AllCourseMaterialDetail struct {
		CourseMaterialID uuid.UUID `json:"course_material_id"`
		SearchBy         string    `json:"search_by"`
		SearchValue      string    `json:"search_value"`
		OrderBy          string    `json:"order_by"`
		OrderDir         string    `json:"order_dir"`
		Offset           int       `json:"offset"`
		Limit            int       `json:"limit"`
	}
)

func (c *CourseMaterialDetail) Response() CourseMaterialDetailResponse {
	return CourseMaterialDetailResponse{
		ID:               c.ID,
		CourseMaterialID: c.CourseMaterialID,
		Position:         c.Position,
		Title:            c.Title,
		TypeMaterial:     c.TypeMaterial,
		Content:          c.Content,
		ImageURL:         c.ImageURL,
	}
}

func (c *CourseMaterialDetail) Add(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	return c.ID, nil
}

func (c *CourseMaterialDetail) All(ctx context.Context, r *util.PostData, param AllCourseMaterialDetail) ([]*CourseMaterialDetail, error) {
	all := []*CourseMaterialDetail{}
	return all, nil
}

func (c *CourseMaterialDetail) Update(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}

func (c *CourseMaterialDetail) Delete(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}
