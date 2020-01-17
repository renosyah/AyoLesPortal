package model

import (
	"context"
	"fmt"

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
	courseMaterialDetailRegister := struct {
		CourseMaterialDetailRegister CourseMaterialDetail `json:"course_material_detail_register"`
	}{
		CourseMaterialDetailRegister: CourseMaterialDetail{},
	}

	query := `mutation {
		course_material_detail_register(
			course_material_id : "%s",
			position : %d,
			title  : "%s",
			type_material : %d,
			content  : "%s",
			image_url  : "%s",
		)
		{
			id,
			course_material_id,
			position,
			title,
			type_material,
			content,
			image_url
		}
	}`

	resp, err := r.Send(fmt.Sprintf(query, c.CourseMaterialID, c.Position, c.Title, c.TypeMaterial, c.Content, c.ImageURL))
	if err != nil {
		return c.ID, err
	}

	err = resp.Body.Error()
	if err != nil {
		return c.ID, err
	}

	err = resp.Body.ConvertData(&courseMaterialDetailRegister)
	if err != nil {
		return c.ID, err
	}

	c.ID = courseMaterialDetailRegister.CourseMaterialDetailRegister.ID

	return c.ID, nil
}

func (c *CourseMaterialDetail) All(ctx context.Context, r *util.PostData, param AllCourseMaterialDetail) ([]*CourseMaterialDetail, error) {
	courseMaterialDetailList := struct {
		CourseMaterialDetailList []*CourseMaterialDetail `json:"course_material_detail_list"`
	}{
		CourseMaterialDetailList: []*CourseMaterialDetail{},
	}

	query := `query {
		course_material_detail_list(
			course_material_id:"%s",
			search_by:"%s",
			search_value:"%s",
			order_by:"%s",
			order_dir:"%s",
			offset:%d,
			limit:%d
		)
		{
			id,
			course_material_id,
			position,
			title,
			type_material,
			content,
			image_url
		}
	}`

	resp, err := r.Send(fmt.Sprintf(query,
		param.CourseMaterialID, param.SearchBy, param.SearchValue, param.OrderBy, param.OrderDir, param.Offset, param.Limit))
	if err != nil {
		return courseMaterialDetailList.CourseMaterialDetailList, err
	}

	err = resp.Body.Error()
	if err != nil {
		return courseMaterialDetailList.CourseMaterialDetailList, err
	}

	err = resp.Body.ConvertData(&courseMaterialDetailList)
	if err != nil {
		return courseMaterialDetailList.CourseMaterialDetailList, err
	}
	return courseMaterialDetailList.CourseMaterialDetailList, nil
}

func (c *CourseMaterialDetail) Update(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}

func (c *CourseMaterialDetail) Delete(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}
