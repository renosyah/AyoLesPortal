package model

import (
	"context"
	"fmt"

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
	courseMaterialRegister := struct {
		CourseMaterialRegister CourseMaterial `json:"course_material_register"`
	}{
		CourseMaterialRegister: CourseMaterial{},
	}

	query := `mutation {
		course_material_register(
			course_id : "%s",
			material_index : %d,
			title : "%s"
		)
		{
			id,
			course_id,
			material_index,
			title
		}
	}`

	resp, err := r.Send(fmt.Sprintf(query, c.CourseID, c.MaterialIndex, c.Title))
	if err != nil {
		return c.ID, err
	}

	err = resp.Body.Error()
	if err != nil {
		return c.ID, err
	}

	err = resp.Body.ConvertData(&courseMaterialRegister)
	if err != nil {
		return c.ID, err
	}

	c.ID = courseMaterialRegister.CourseMaterialRegister.ID

	return c.ID, nil
}

func (c *CourseMaterial) One(ctx context.Context, r *util.PostData) (*CourseMaterial, error) {
	courseMaterialDetail := struct {
		CourseMaterialDetail *CourseMaterial `json:"course_material_detail"`
	}{
		CourseMaterialDetail: &CourseMaterial{},
	}

	query := `query {
		course_material_detail(
			id: "%s"
		)
		{
			id,
			course_id,
			material_index,
			title
		}
	}`

	resp, err := r.Send(fmt.Sprintf(query, c.ID))
	if err != nil {
		return courseMaterialDetail.CourseMaterialDetail, err
	}

	err = resp.Body.Error()
	if err != nil {
		return courseMaterialDetail.CourseMaterialDetail, err
	}

	err = resp.Body.ConvertData(&courseMaterialDetail)
	if err != nil {
		return courseMaterialDetail.CourseMaterialDetail, err
	}
	return courseMaterialDetail.CourseMaterialDetail, nil
}

func (c *CourseMaterial) All(ctx context.Context, r *util.PostData, param AllCourseMaterial) ([]*CourseMaterial, error) {
	courseMaterialList := struct {
		CourseMaterialList []*CourseMaterial `json:"course_material_list"`
	}{
		CourseMaterialList: []*CourseMaterial{},
	}

	query := `query {
		course_material_list(
			course_id:"%s",
			search_by:"%s",
			search_value:"%s",
			order_by:"%s",
			order_dir:"%s",
			offset:%d,
			limit:%d
		)
		{
			id,
			course_id,
			material_index,
			title
		}
	}`

	resp, err := r.Send(fmt.Sprintf(query,
		param.CourseID, param.SearchBy, param.SearchValue, param.OrderBy, param.OrderDir, param.Offset, param.Limit))
	if err != nil {
		return courseMaterialList.CourseMaterialList, err
	}

	err = resp.Body.Error()
	if err != nil {
		return courseMaterialList.CourseMaterialList, err
	}

	err = resp.Body.ConvertData(&courseMaterialList)
	if err != nil {
		return courseMaterialList.CourseMaterialList, err
	}
	return courseMaterialList.CourseMaterialList, nil
}

func (c *CourseMaterial) Update(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}

func (c *CourseMaterial) Delete(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}
