package model

import (
	"context"
	"fmt"

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

	bannerRegister := struct {
		BannerRegister Banner `json:"banner_register"`
	}{
		BannerRegister: Banner{},
	}

	query := `mutation {
		banner_register(
				title : "%s",
				content : "%s",
				image_url: "%s"
			) {
				id
			}
	}`

	resp, err := r.Send(fmt.Sprintf(query, b.Title, b.Content, b.ImageURL))
	if err != nil {
		return b.ID, err
	}

	err = resp.Body.Error()
	if err != nil {
		return b.ID, err
	}

	err = resp.Body.ConvertData(&bannerRegister)
	if err != nil {
		return b.ID, err
	}

	b.ID = bannerRegister.BannerRegister.ID

	return b.ID, nil
}

func (b *Banner) One(ctx context.Context, r *util.PostData) (*Banner, error) {
	bannerDetail := struct {
		BannerDetail *Banner `json:"banner_detail"`
	}{
		BannerDetail: &Banner{},
	}

	query := `query {
		banner_detail(
				id:"%s"
			) {
				id,
				title,
				content,
				image_url
			}
	}`

	resp, err := r.Send(fmt.Sprintf(query, b.ID))
	if err != nil {
		return bannerDetail.BannerDetail, err
	}

	err = resp.Body.Error()
	if err != nil {
		return bannerDetail.BannerDetail, err
	}

	err = resp.Body.ConvertData(&bannerDetail)
	if err != nil {
		return bannerDetail.BannerDetail, err
	}

	return bannerDetail.BannerDetail, nil
}

func (b *Banner) All(ctx context.Context, r *util.PostData, param AllBanner) ([]*Banner, error) {
	bannerList := struct {
		BannerList []*Banner `json:"banner_list"`
	}{
		BannerList: []*Banner{},
	}

	query := `query {
		banner_list(
			search_by:"%s",
			search_value:"%s",
			order_by:"%s",
			order_dir:"%s",
			offset:%d,
			limit:%d
		)
		{
			id,
			title,
			content,
			image_url
		 }
	}`

	resp, err := r.Send(fmt.Sprintf(query,
		param.SearchBy, param.SearchValue, param.OrderBy, param.OrderDir, param.Offset, param.Limit))
	if err != nil {
		return bannerList.BannerList, err
	}

	err = resp.Body.Error()
	if err != nil {
		return bannerList.BannerList, err
	}

	err = resp.Body.ConvertData(&bannerList)
	if err != nil {
		return bannerList.BannerList, err
	}

	return bannerList.BannerList, nil
}

func (b *Banner) Update(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}

func (b *Banner) Delete(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}
