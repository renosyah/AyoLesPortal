package model

import (
	"context"
	"fmt"

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
	categoryRegister := struct {
		CategoryRegister Category `json:"category_register"`
	}{
		CategoryRegister: Category{},
	}

	query := `mutation {
		category_register(
				name : "%s",
				image_url : "%s"
			) {
				id,
				name,
				image_url
			}
	}`

	resp, err := r.Send(fmt.Sprintf(query, c.Name, c.ImageURL))
	if err != nil {
		return c.ID, err
	}

	err = resp.Body.Error()
	if err != nil {
		return c.ID, err
	}

	err = resp.Body.ConvertData(&categoryRegister)
	if err != nil {
		return c.ID, err
	}

	c.ID = categoryRegister.CategoryRegister.ID

	return c.ID, nil
}

func (c *Category) One(ctx context.Context, r *util.PostData) (*Category, error) {

	categoryDetail := struct {
		CategoryDetail *Category `json:"category_detail"`
	}{
		CategoryDetail: &Category{},
	}

	query := `query {
		category_detail(
				id:"%s"
			) {
				id,
				name,
				image_url
			}
	}`

	resp, err := r.Send(fmt.Sprintf(query, c.ID))
	if err != nil {
		return categoryDetail.CategoryDetail, err
	}

	err = resp.Body.Error()
	if err != nil {
		return categoryDetail.CategoryDetail, err
	}

	err = resp.Body.ConvertData(&categoryDetail)
	if err != nil {
		return categoryDetail.CategoryDetail, err
	}

	return categoryDetail.CategoryDetail, nil
}

func (c *Category) All(ctx context.Context, r *util.PostData, param AllCategory) ([]*Category, error) {
	categoryList := struct {
		CategoryList []*Category `json:"category_list"`
	}{
		CategoryList: []*Category{},
	}

	query := `query {
		category_list(
			search_by:"%s",
			search_value:"%s",
			order_by:"%s",
			order_dir:"%s",
			offset:%d,
			limit:%d
		)
		{
			id,
			name,
			image_url
		 }
	}`

	resp, err := r.Send(fmt.Sprintf(query,
		param.SearchBy, param.SearchValue, param.OrderBy, param.OrderDir, param.Offset, param.Limit))
	if err != nil {
		return categoryList.CategoryList, err
	}

	err = resp.Body.Error()
	if err != nil {
		return categoryList.CategoryList, err
	}

	err = resp.Body.ConvertData(&categoryList)
	if err != nil {
		return categoryList.CategoryList, err
	}

	return categoryList.CategoryList, nil
}

func (c *Category) Update(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}

func (c *Category) Delete(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}
