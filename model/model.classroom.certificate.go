package model

import (
	"context"
	"fmt"
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
	classroomCertificateRegister := struct {
		ClassroomCertificateRegister ClassRoomCertificate `json:"classroom_certificate_register"`
	}{
		ClassroomCertificateRegister: ClassRoomCertificate{},
	}

	query := `mutation {
		classroom_certificate_register(
			classroom_id : "%s"
		)
		{
			id,
			classroom_id,
			hash_id
		 }
	}`

	resp, err := r.Send(fmt.Sprintf(query, c.ClassroomID))
	if err != nil {
		return c.ID, err
	}

	err = resp.Body.Error()
	if err != nil {
		return c.ID, err
	}

	err = resp.Body.ConvertData(&classroomCertificateRegister)
	if err != nil {
		return c.ID, err
	}

	c.ID = classroomCertificateRegister.ClassroomCertificateRegister.ID

	return c.ID, nil
}

func (c *ClassRoomCertificate) One(ctx context.Context, r *util.PostData) (*ClassRoomCertificate, error) {
	classroomCertificateDetail := struct {
		ClassroomCertificateDetail *ClassRoomCertificate `json:"classroom_certificate_detail"`
	}{
		ClassroomCertificateDetail: &ClassRoomCertificate{},
	}

	query := `query {
		classroom_certificate_detail(
			classroom_id: "%s"
		)
		{
			id,
			classroom_id,
			hash_id
		 }
	}`

	resp, err := r.Send(fmt.Sprintf(query, c.ClassroomID))
	if err != nil {
		return classroomCertificateDetail.ClassroomCertificateDetail, err
	}

	err = resp.Body.Error()
	if err != nil {
		return classroomCertificateDetail.ClassroomCertificateDetail, err
	}

	err = resp.Body.ConvertData(&classroomCertificateDetail)
	if err != nil {
		return classroomCertificateDetail.ClassroomCertificateDetail, err
	}

	return classroomCertificateDetail.ClassroomCertificateDetail, nil
}

func (c *ClassRoomCertificate) All(ctx context.Context, r *util.PostData, param AllClassRoomCertificate) ([]*ClassRoomCertificate, error) {
	classroomCertificateList := struct {
		ClassroomCertificateList []*ClassRoomCertificate `json:"classroom_certificate_list"`
	}{
		ClassroomCertificateList: []*ClassRoomCertificate{},
	}

	query := `query {
		classroom_certificate_list(
			student_id:"%s",
			order_by:"%s",
			order_dir:"%s",
			offset:%d,
			limit:%d
		)
		{
			id,
			classroom_id,
			hash_id
		 }
	}`

	resp, err := r.Send(fmt.Sprintf(query,
		param.StudentID, param.OrderBy, param.OrderDir, param.Offset, param.Limit))
	if err != nil {
		return classroomCertificateList.ClassroomCertificateList, err
	}

	err = resp.Body.Error()
	if err != nil {
		return classroomCertificateList.ClassroomCertificateList, err
	}

	err = resp.Body.ConvertData(&classroomCertificateList)
	if err != nil {
		return classroomCertificateList.ClassroomCertificateList, err
	}

	return classroomCertificateList.ClassroomCertificateList, nil
}

func (c *ClassRoomCertificate) Update(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}

func (c *ClassRoomCertificate) Delete(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}
