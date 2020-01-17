package model

import (
	"context"
	"fmt"

	"github.com/renosyah/AyoLesPortal/util"
	uuid "github.com/satori/go.uuid"
)

type (
	CourseDetail struct {
		ID              uuid.UUID `json:"id"`
		CourseID        uuid.UUID `json:"course_id"`
		OverviewText    string    `json:"overview_text"`
		DescriptionText string    `json:"description_text"`
		ImageURL        string    `json:"image_url"`
	}

	CourseDetailResponse struct {
		ID              uuid.UUID `json:"id"`
		CourseID        uuid.UUID `json:"course_id"`
		OverviewText    string    `json:"overview_text"`
		DescriptionText string    `json:"description_text"`
		ImageURL        string    `json:"image_url"`
	}

	AllCourseDetail struct {
		CourseID    uuid.UUID `json:"course_id"`
		SearchBy    string    `json:"search_by"`
		SearchValue string    `json:"search_value"`
		OrderBy     string    `json:"order_by"`
		OrderDir    string    `json:"order_dir"`
		Offset      int       `json:"offset"`
		Limit       int       `json:"limit"`
	}
)

func (a AllCourseDetail) IsWithCourseID() string {
	var emptyID uuid.UUID
	if a.CourseID == emptyID {
		return ""
	}
	return fmt.Sprintf(`AND course_id = '%s'`, a.CourseID)
}

func (c *CourseDetail) Response() CourseDetailResponse {
	return CourseDetailResponse{
		ID:              c.ID,
		CourseID:        c.CourseID,
		OverviewText:    c.OverviewText,
		DescriptionText: c.DescriptionText,
		ImageURL:        c.ImageURL,
	}
}

func (c *CourseDetail) Add(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	courseDetailRegister := struct {
		CourseDetailRegister CourseDetail `json:"course_detail_register"`
	}{
		CourseDetailRegister: CourseDetail{},
	}

	query := `mutation {
		course_detail_register(
			course_id : "%s",
			overview_text:"%s",
			description_text : "%s",
			image_url : "%s"
		)
		{
			id,
			course_id,
			overview_text,
			description_text,
			image_url
		}
	}`

	resp, err := r.Send(fmt.Sprintf(query, c.CourseID, c.OverviewText, c.DescriptionText, c.ImageURL))
	if err != nil {
		return c.ID, err
	}

	err = resp.Body.Error()
	if err != nil {
		return c.ID, err
	}

	err = resp.Body.ConvertData(&courseDetailRegister)
	if err != nil {
		return c.ID, err
	}

	c.ID = courseDetailRegister.CourseDetailRegister.ID

	return c.ID, nil
}

func (c *CourseDetail) One(ctx context.Context, r *util.PostData) (*CourseDetail, error) {
	one := &CourseDetail{}
	return one, nil
}

func (c *CourseDetail) All(ctx context.Context, r *util.PostData, param AllCourseDetail) ([]*CourseDetail, error) {
	courseDetailList := struct {
		CourseDetailList []*CourseDetail `json:"course_detail_list"`
	}{
		CourseDetailList: []*CourseDetail{},
	}

	query := `query {
		course_detail_list(
			course_id : "%s",
			search_by:"%s",
			search_value:"%s",
			order_by:"%s",
			order_dir:"%s",
			offset:%d,
			limit:%d
		)
		{
			id,
			course_id,
			overview_text,
			description_text,
			image_url
		}
	} `

	resp, err := r.Send(fmt.Sprintf(query,
		param.CourseID, param.SearchBy, param.SearchValue, param.OrderBy, param.OrderDir, param.Offset, param.Limit))
	if err != nil {
		return courseDetailList.CourseDetailList, err
	}

	err = resp.Body.Error()
	if err != nil {
		return courseDetailList.CourseDetailList, err
	}

	err = resp.Body.ConvertData(&courseDetailList)
	if err != nil {
		return courseDetailList.CourseDetailList, err
	}

	return courseDetailList.CourseDetailList, nil

}

func (c *CourseDetail) AllByCourseID(ctx context.Context, r *util.PostData) ([]CourseDetail, error) {
	all := []CourseDetail{}
	return all, nil
}

func (c *CourseDetail) Update(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}

func (c *CourseDetail) Delete(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}
