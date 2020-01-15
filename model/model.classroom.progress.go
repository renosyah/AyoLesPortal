package model

import (
	"context"

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
	return c.ID, nil
}

func (c *ClassRoomProgress) One(ctx context.Context, r *util.PostData) (*ClassRoomProgress, error) {
	one := &ClassRoomProgress{}

	return one, nil
}

func (c *ClassRoomProgress) All(ctx context.Context, r *util.PostData, param AllClassRoomProgress) ([]*ClassRoomProgress, error) {
	all := []*ClassRoomProgress{}
	return all, nil
}

func (c *ClassRoomProgress) Update(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}

func (c *ClassRoomProgress) Delete(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}
