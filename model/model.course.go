package model

import (
	"context"

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
	return c.ID, nil
}

func (c *Course) One(ctx context.Context, r *util.PostData) (*Course, error) {
	one := &Course{
		Teacher:  &Teacher{},
		Category: &Category{},
	}
	return one, nil
}

func (c *Course) All(ctx context.Context, r *util.PostData, param AllCourse) ([]*Course, error) {
	all := []*Course{}
	return all, nil
}

func (c *Course) Update(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}

func (c *Course) Delete(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}
