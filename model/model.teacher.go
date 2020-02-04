package model

import (
	"context"
	"errors"
	"fmt"

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
	teacherRegister := struct {
		TeacherRegister Teacher `json:"teacher_register"`
	}{
		TeacherRegister: Teacher{},
	}

	query := `mutation {
		teacher_register(
			name:"%s",
			email:"%s",
			password:"%s"
		)
		{
			id,
			name,
			email
		}
	}`

	resp, err := r.Send(fmt.Sprintf(query, t.Name, t.Email, t.Password))
	if err != nil {
		return t.ID, err
	}

	err = resp.Body.Error()
	if err != nil {
		return t.ID, err
	}

	err = resp.Body.ConvertData(&teacherRegister)
	if err != nil {
		return t.ID, err
	}

	t.ID = teacherRegister.TeacherRegister.ID

	return t.ID, nil
}

func (t *Teacher) One(ctx context.Context, r *util.PostData) (*Teacher, error) {
	teacherDetail := struct {
		TeacherDetail *Teacher `json:"teacher_detail"`
	}{
		TeacherDetail: &Teacher{},
	}

	query := `query {
		teacher_detail(
			id: "%s"
		)
		{
			id,
			name,
			email
		}
	}`

	resp, err := r.Send(fmt.Sprintf(query, t.ID))
	if err != nil {
		return teacherDetail.TeacherDetail, err
	}

	err = resp.Body.Error()
	if err != nil {
		return teacherDetail.TeacherDetail, err
	}

	err = resp.Body.ConvertData(&teacherDetail)
	if err != nil {
		return teacherDetail.TeacherDetail, err
	}
	return teacherDetail.TeacherDetail, nil
}

func (t *Teacher) All(ctx context.Context, r *util.PostData, param AllTeacher) ([]*Teacher, error) {
	teacherList := struct {
		TeacherList []*Teacher `json:"teacher_list"`
	}{
		TeacherList: []*Teacher{},
	}

	query := `query {
		teacher_list(
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
			email
		}
	}`

	resp, err := r.Send(fmt.Sprintf(query,
		param.SearchBy, param.SearchValue, param.OrderBy, param.OrderDir, param.Offset, param.Limit))
	if err != nil {
		return teacherList.TeacherList, err
	}

	err = resp.Body.Error()
	if err != nil {
		return teacherList.TeacherList, err
	}

	err = resp.Body.ConvertData(&teacherList)
	if err != nil {
		return teacherList.TeacherList, err
	}
	return teacherList.TeacherList, nil
}

func (t *Teacher) Login(ctx context.Context, r *util.PostData) (*Teacher, error) {
	teacherLogin := struct {
		TeacherLogin *Teacher `json:"teacher_login"`
	}{
		TeacherLogin: &Teacher{},
	}

	query := `query {
		teacher_login(
			email:"%s",
			password:"%s"
		)
		{
			id,
			name,
			email
		}
	}`

	resp, err := r.Send(fmt.Sprintf(query, t.Email, t.Password))
	if err != nil {
		return teacherLogin.TeacherLogin, err
	}

	err = resp.Body.Error()
	if err != nil {
		return teacherLogin.TeacherLogin, err
	}

	err = resp.Body.ConvertData(&teacherLogin)
	if err != nil {
		return teacherLogin.TeacherLogin, err
	}
	return teacherLogin.TeacherLogin, nil
}

func (t *Teacher) Update(ctx context.Context, r *util.PostData) error {

	teacherUpdate := struct {
		TeacherUpdate Teacher `json:"teacher_update"`
	}{
		TeacherUpdate: Teacher{},
	}

	query := `mutation {
		teacher_update(
			id : "%s",
			name:"%s",
			email:"%s",
			password:"%s"
		)
		{
			id,
			name,
			email
		}
	}`

	resp, err := r.Send(fmt.Sprintf(query, t.ID, t.Name, t.Email, t.Password))
	if err != nil {
		return err
	}

	err = resp.Body.Error()
	if err != nil {
		return err
	}

	err = resp.Body.ConvertData(&teacherUpdate)
	if err != nil {
		return err
	}

	var emptyUUID uuid.UUID
	if teacherUpdate.TeacherUpdate.ID == emptyUUID {
		return errors.New("update failed,id is empty")
	}

	return nil
}

func (t *Teacher) Delete(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}
