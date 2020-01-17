package model

import (
	"context"
	"fmt"

	"github.com/renosyah/AyoLesPortal/util"
	uuid "github.com/satori/go.uuid"
)

type (
	ClassRoomProgress struct {
		ID               uuid.UUID `json:"id"`
		ClassRoomID      uuid.UUID `json:"classroom_id"`
		CourseMaterialID uuid.UUID `json:"course_material_id"`
	}
	ClassRoomProgressResponse struct {
		ID               uuid.UUID `json:"id"`
		ClassRoomID      uuid.UUID `json:"classroom_id"`
		CourseMaterialID uuid.UUID `json:"course_material_id"`
	}

	AllClassRoomProgress struct {
		ClassRoomID uuid.UUID `json:"classroom_id"`
		Offset      int       `json:"offset"`
		Limit       int       `json:"limit"`
	}
)

func (c *ClassRoomProgress) Response() ClassRoomProgressResponse {
	return ClassRoomProgressResponse{
		ID:               c.ID,
		ClassRoomID:      c.ClassRoomID,
		CourseMaterialID: c.CourseMaterialID,
	}
}

func (c *ClassRoomProgress) Add(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	classroomProgressRegister := struct {
		ClassroomProgressRegister ClassRoomProgress `json:"classroom_progress_register"`
	}{
		ClassroomProgressRegister: ClassRoomProgress{},
	}

	query := `mutation {
		classroom_progress_register(
			classroom_id :"%s",
			course_material_id :"%s"
		)
		{
			id,
			classroom_id,
			course_material_id
		 }
	}`

	resp, err := r.Send(fmt.Sprintf(query, c.ClassRoomID, c.CourseMaterialID))
	if err != nil {
		return c.ID, err
	}

	err = resp.Body.Error()
	if err != nil {
		return c.ID, err
	}

	err = resp.Body.ConvertData(&classroomProgressRegister)
	if err != nil {
		return c.ID, err
	}

	c.ID = classroomProgressRegister.ClassroomProgressRegister.ID

	return c.ID, nil
}

func (c *ClassRoomProgress) One(ctx context.Context, r *util.PostData) (*ClassRoomProgress, error) {
	classroomProgressDetail := struct {
		ClassroomProgressDetail *ClassRoomProgress `json:"classroom_progress_detail"`
	}{
		ClassroomProgressDetail: &ClassRoomProgress{},
	}

	query := `query {
		classroom_progress_detail(
			id: "%s"
		)
		{
			id,
			classroom_id,
			course_material_id
		 }
	}`

	resp, err := r.Send(fmt.Sprintf(query, c.ID))
	if err != nil {
		return classroomProgressDetail.ClassroomProgressDetail, err
	}

	err = resp.Body.Error()
	if err != nil {
		return classroomProgressDetail.ClassroomProgressDetail, err
	}

	err = resp.Body.ConvertData(&classroomProgressDetail)
	if err != nil {
		return classroomProgressDetail.ClassroomProgressDetail, err
	}
	return classroomProgressDetail.ClassroomProgressDetail, nil
}

func (c *ClassRoomProgress) All(ctx context.Context, r *util.PostData, param AllClassRoomProgress) ([]*ClassRoomProgress, error) {
	classroomProgressList := struct {
		ClassroomProgressList []*ClassRoomProgress `json:"classroom_progress_list"`
	}{
		ClassroomProgressList: []*ClassRoomProgress{},
	}

	query := `query {
		classroom_progress_list(
			classroom_id:"%s",
			offset:%d,
			limit:%d
		)
		{
			id,
			classroom_id,
			course_material_id
		 }
	}`

	resp, err := r.Send(fmt.Sprintf(query, param.ClassRoomID, param.Offset, param.Limit))
	if err != nil {
		return classroomProgressList.ClassroomProgressList, err
	}

	err = resp.Body.Error()
	if err != nil {
		return classroomProgressList.ClassroomProgressList, err
	}

	err = resp.Body.ConvertData(&classroomProgressList)
	if err != nil {
		return classroomProgressList.ClassroomProgressList, err
	}
	return classroomProgressList.ClassroomProgressList, nil
}

func (c *ClassRoomProgress) Update(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}

func (c *ClassRoomProgress) Delete(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}
