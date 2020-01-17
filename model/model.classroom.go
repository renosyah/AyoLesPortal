package model

import (
	"context"

	"github.com/renosyah/AyoLesPortal/util"
	uuid "github.com/satori/go.uuid"
)

type (
	ClassRoom struct {
		ID        uuid.UUID `json:"id"`
		Course    *Course   `json:"course"`
		StudentID uuid.UUID `json:"student_id"`
	}

	ClassRoomResponse struct {
		ID        uuid.UUID      `json:"id"`
		Course    CourseResponse `json:"course"`
		StudentID uuid.UUID      `json:"student_id"`
	}

	AllClassRoom struct {
		StudentID   uuid.UUID `json:"student_id"`
		SearchBy    string    `json:"search_by"`
		SearchValue string    `json:"search_value"`
		OrderBy     string    `json:"order_by"`
		OrderDir    string    `json:"order_dir"`
		Offset      int       `json:"offset"`
		Limit       int       `json:"limit"`
	}
)

func (c *ClassRoom) Response() ClassRoomResponse {
	return ClassRoomResponse{
		ID:        c.ID,
		Course:    c.Course.Response(),
		StudentID: c.StudentID,
	}
}

func (c *ClassRoom) Add(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	return c.ID, nil
}

func (c *ClassRoom) One(ctx context.Context, r *util.PostData) (*ClassRoom, error) {
	one := &ClassRoom{
		Course: &Course{},
	}
	return one, nil
}

func (c *ClassRoom) OneByStudentIdAndCourseId(ctx context.Context, r *util.PostData) (*ClassRoom, error) {
	one := &ClassRoom{
		Course: &Course{},
	}
	return one, nil
}

func (c *ClassRoom) All(ctx context.Context, r *util.PostData, param AllClassRoom) ([]*ClassRoom, error) {
	all := []*ClassRoom{}
	return all, nil
}

func (c *ClassRoom) Update(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}

func (c *ClassRoom) Delete(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}

// this model
// doesnot have any api
// to call from core
