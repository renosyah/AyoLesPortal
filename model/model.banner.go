package model

import (
	"context"

	"github.com/renosyah/AyoLesPortal/util"
	uuid "github.com/satori/go.uuid"
)

type (
	Banner struct {
		ID       uuid.UUID `json:"id"`
		Title    string    `json:"title"`
		Content  string    `json:"content"`
		ImageURL string    `json:"image_url"`
	}

	BannerResponse struct {
		ID       uuid.UUID `json:"id"`
		Title    string    `json:"title"`
		Content  string    `json:"content"`
		ImageURL string    `json:"image_url"`
	}

	AllBanner struct {
		SearchBy    string `json:"search_by"`
		SearchValue string `json:"search_value"`
		OrderBy     string `json:"order_by"`
		OrderDir    string `json:"order_dir"`
		Offset      int    `json:"offset"`
		Limit       int    `json:"limit"`
	}
)

func (c *Banner) Response() BannerResponse {
	return BannerResponse{
		ID:       c.ID,
		Title:    c.Title,
		Content:  c.Content,
		ImageURL: c.ImageURL,
	}
}

func (b *Banner) Add(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	return b.ID, nil
}

func (b *Banner) One(ctx context.Context, r *util.PostData) (*Banner, error) {
	one := &Banner{}
	return one, nil
}

func (b *Banner) All(ctx context.Context, r *util.PostData, param AllBanner) ([]*Banner, error) {
	all := []*Banner{}
	return all, nil
}

func (b *Banner) Update(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}

func (b *Banner) Delete(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}
