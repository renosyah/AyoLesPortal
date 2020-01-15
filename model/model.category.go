package model

import (
	"context"

	"github.com/renosyah/AyoLesPortal/util"
	uuid "github.com/satori/go.uuid"
)

type (
	Category struct {
		ID       uuid.UUID `json:"id"`
		Name     string    `json:"name"`
		ImageURL string    `json:"image_url"`
	}

	CategoryResponse struct {
		ID       uuid.UUID `json:"id"`
		Name     string    `json:"name"`
		ImageURL string    `json:"image_url"`
	}

	AllCategory struct {
		SearchBy    string `json:"search_by"`
		SearchValue string `json:"search_value"`
		OrderBy     string `json:"order_by"`
		OrderDir    string `json:"order_dir"`
		Offset      int    `json:"offset"`
		Limit       int    `json:"limit"`
	}
)

func (c *Category) Response() CategoryResponse {
	return CategoryResponse{
		ID:       c.ID,
		Name:     c.Name,
		ImageURL: c.ImageURL,
	}
}

func (c *Category) Add(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	return c.ID, nil
}

func (c *Category) One(ctx context.Context, r *util.PostData) (*Category, error) {
	one := &Category{}
	return one, nil
}

func (c *Category) All(ctx context.Context, r *util.PostData, param AllCategory) ([]*Category, error) {
	all := []*Category{}
	return all, nil
}

func (c *Category) Update(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}

func (c *Category) Delete(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}
