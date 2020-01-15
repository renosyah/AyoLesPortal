package model

import (
	"context"

	"github.com/renosyah/AyoLesPortal/util"
	uuid "github.com/satori/go.uuid"
)

type (
	Student struct {
		ID       uuid.UUID `json:"id"`
		Name     string    `json:"name"`
		Email    string    `json:"email"`
		Password string    `json:"password"`
	}

	StudentResponse struct {
		ID       uuid.UUID `json:"id"`
		Name     string    `json:"name"`
		Email    string    `json:"email"`
		Password string    `json:"-"`
	}

	AllStudent struct {
		SearchBy    string `json:"search_by"`
		SearchValue string `json:"search_value"`
		OrderBy     string `json:"order_by"`
		OrderDir    string `json:"order_dir"`
		Offset      int    `json:"offset"`
		Limit       int    `json:"limit"`
	}
)

func (s *Student) Response() StudentResponse {
	return StudentResponse{
		ID:       s.ID,
		Name:     s.Name,
		Email:    s.Email,
		Password: s.Password,
	}
}

func (s *Student) Add(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	return s.ID, nil
}

func (s *Student) One(ctx context.Context, r *util.PostData) (*Student, error) {
	one := &Student{}
	return one, nil
}

func (s *Student) All(ctx context.Context, r *util.PostData, param AllStudent) ([]*Student, error) {
	all := []*Student{}
	return all, nil
}

func (s *Student) OneByEmail(ctx context.Context, r *util.PostData) (*Student, error) {
	one := &Student{}
	return one, nil
}

func (s *Student) Update(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	return s.ID, nil
}

func (s *Student) Delete(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}
