package model

import (
	"context"
	"time"

	"github.com/renosyah/AyoLesPortal/util"
	uuid "github.com/satori/go.uuid"
)

type (
	ClassRoomCertificate struct {
		ID          uuid.UUID `json:"id"`
		ClassroomID uuid.UUID `json:"classroom_id"`
		HashID      string    `json:"hash_id"`
		CreateAt    time.Time `json:"-"`
	}

	ClassRoomCertificateResponse struct {
		ID          uuid.UUID `json:"id"`
		ClassroomID uuid.UUID `json:"classroom_id"`
		HashID      string    `json:"hash_id"`
		CreateAt    time.Time `json:"-"`
	}

	AllClassRoomCertificate struct {
		StudentID uuid.UUID `json:"student_id"`
		OrderBy   string    `json:"order_by"`
		OrderDir  string    `json:"order_dir"`
		Offset    int       `json:"offset"`
		Limit     int       `json:"limit"`
	}
)

func (c *ClassRoomCertificate) Response() ClassRoomCertificateResponse {
	return ClassRoomCertificateResponse{
		ID:          c.ID,
		ClassroomID: c.ClassroomID,
		HashID:      c.HashID,
		CreateAt:    c.CreateAt,
	}
}

func (c *ClassRoomCertificate) Add(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	return c.ID, nil
}

func (c *ClassRoomCertificate) One(ctx context.Context, r *util.PostData) (*ClassRoomCertificate, error) {
	one := &ClassRoomCertificate{}
	return one, nil
}

func (c *ClassRoomCertificate) All(ctx context.Context, r *util.PostData, param AllClassRoomCertificate) ([]*ClassRoomCertificate, error) {
	all := []*ClassRoomCertificate{}
	return all, nil
}

func (c *ClassRoomCertificate) Update(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}

func (c *ClassRoomCertificate) Delete(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}
