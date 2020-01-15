package model

import (
	"context"

	"github.com/renosyah/AyoLesPortal/util"
	uuid "github.com/satori/go.uuid"
)

type (
	Teacher struct {
		ID       uuid.UUID `json:"id"`
		Name     string    `json:"name"`
		Email    string    `json:"email"`
		Password string    `json:"password"`
	}

	TeacherResponse struct {
		ID       uuid.UUID `json:"id"`
		Name     string    `json:"name"`
		Email    string    `json:"email"`
		Password string    `json:"-"`
	}
	AllTeacher struct {
		SearchBy    string `json:"search_by"`
		SearchValue string `json:"search_value"`
		OrderBy     string `json:"order_by"`
		OrderDir    string `json:"order_dir"`
		Offset      int    `json:"offset"`
		Limit       int    `json:"limit"`
	}
)

func (t *Teacher) Response() TeacherResponse {
	return TeacherResponse{
		ID:       t.ID,
		Name:     t.Name,
		Email:    t.Email,
		Password: t.Password,
	}
}

func (t *Teacher) Add(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	return t.ID, nil
}

func (t *Teacher) One(ctx context.Context, r *util.PostData) (*Teacher, error) {
	one := &Teacher{}
	return one, nil
}

func (t *Teacher) All(ctx context.Context, r *util.PostData, param AllTeacher) ([]*Teacher, error) {
	all := []*Teacher{}
	return all, nil
}

func (t *Teacher) OneByEmail(ctx context.Context, r *util.PostData) (*Teacher, error) {
	one := &Teacher{}
	return one, nil
}

func (t *Teacher) Update(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	return t.ID, nil
}

func (t *Teacher) Delete(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}
