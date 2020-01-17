package model

import (
	"context"
	"fmt"

	"github.com/renosyah/AyoLesPortal/util"
	uuid "github.com/satori/go.uuid"
)

type (
	Course struct {
		ID            uuid.UUID      `json:"id"`
		CourseName    string         `json:"course_name"`
		ImageURL      string         `json:"image_url"`
		Teacher       *Teacher       `json:"teacher"`
		Category      *Category      `json:"category"`
		CourseDetails []CourseDetail `json:"course_details"`
	}

	CourseResponse struct {
		ID            uuid.UUID              `json:"id"`
		CourseName    string                 `json:"course_name"`
		ImageURL      string                 `json:"image_url"`
		Teacher       TeacherResponse        `json:"teacher"`
		Category      CategoryResponse       `json:"category"`
		CourseDetails []CourseDetailResponse `json:"course_details"`
	}

	AllCourse struct {
		CategoryID  uuid.UUID `json:"category_id"`
		TeacherID   uuid.UUID `json:"teacher_id"`
		SearchBy    string    `json:"search_by"`
		SearchValue string    `json:"search_value"`
		OrderBy     string    `json:"order_by"`
		OrderDir    string    `json:"order_dir"`
		Offset      int       `json:"offset"`
		Limit       int       `json:"limit"`
	}
)

func (c *Course) Response() CourseResponse {
	details := []CourseDetailResponse{}
	for _, v := range c.CourseDetails {
		details = append(details, v.Response())
	}
	return CourseResponse{
		ID:            c.ID,
		CourseName:    c.CourseName,
		ImageURL:      c.ImageURL,
		Teacher:       c.Teacher.Response(),
		Category:      c.Category.Response(),
		CourseDetails: details,
	}
}

func (c *Course) Add(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	courseRegister := struct {
		CourseRegister Course `json:"course_register"`
	}{
		CourseRegister: Course{},
	}

	query := `mutation {
		course_register(
			course_name:"%s",
			teacher_id :"%s",
			category_id:"%s",
			image_url : "%s",
		)
		{
			id,
			course_name,
			image_url,
			teacher { id, name, email } ,
			category {id, name, image_url},
			course_details { id,course_id , overview_text, description_text,image_url }
		}
	}`

	resp, err := r.Send(fmt.Sprintf(query, c.CourseName, c.Teacher.ID, c.Category.ID, c.ImageURL))
	if err != nil {
		return c.ID, err
	}

	err = resp.Body.Error()
	if err != nil {
		return c.ID, err
	}

	err = resp.Body.ConvertData(&courseRegister)
	if err != nil {
		return c.ID, err
	}

	c.ID = courseRegister.CourseRegister.ID

	return c.ID, nil
}

func (c *Course) One(ctx context.Context, r *util.PostData) (*Course, error) {
	courseDetail := struct {
		CourseDetail *Course `json:"course_register"`
	}{
		CourseDetail: &Course{},
	}

	query := `query {
		course_detail(
			id: "%s"
		)
		{
			id,
			course_name,
			image_url,
			teacher { id, name, email } ,
			category {id, name, image_url},
			course_details { id,course_id , overview_text, description_text,image_url }
		}
	}`

	resp, err := r.Send(fmt.Sprintf(query, c.ID))
	if err != nil {
		return courseDetail.CourseDetail, err
	}

	err = resp.Body.Error()
	if err != nil {
		return courseDetail.CourseDetail, err
	}

	err = resp.Body.ConvertData(&courseDetail)
	if err != nil {
		return courseDetail.CourseDetail, err
	}
	return courseDetail.CourseDetail, nil
}

func (c *Course) All(ctx context.Context, r *util.PostData, param AllCourse) ([]*Course, error) {
	courseList := struct {
		CourseList []*Course `json:"course_list"`
	}{
		CourseList: []*Course{},
	}

	query := `query {
		course_list(
			category_id:"%s",
			teacher_id : "%s",
			search_by:"%s",
			search_value:"%s",
			order_by:"%s",
			order_dir:"%s",
			offset:%d,
			limit:%d
		)
		{
			id,
			course_name,
			image_url,
			teacher {id, name, email } ,
			category {id, name, image_url},
			course_details { id,course_id , overview_text, description_text,image_url }
		}
	}`

	resp, err := r.Send(fmt.Sprintf(query,
		param.CategoryID, param.TeacherID, param.SearchBy, param.SearchValue, param.OrderBy, param.OrderDir, param.Offset, param.Limit))
	if err != nil {
		return courseList.CourseList, err
	}

	err = resp.Body.Error()
	if err != nil {
		return courseList.CourseList, err
	}

	err = resp.Body.ConvertData(&courseList)
	if err != nil {
		return courseList.CourseList, err
	}
	return courseList.CourseList, nil
}

func (c *Course) Update(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}

func (c *Course) Delete(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}
