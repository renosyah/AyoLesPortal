package model

import (
	"context"
	"fmt"

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
	studentRegister := struct {
		StudentRegister Student `json:"student_register"`
	}{
		StudentRegister: Student{},
	}

	query := `mutation {
		student_register(
			name:"%s",
			email:"%s",
			password:"%s"
		)
		{
			id,
			name,
			email
		}
	} `

	resp, err := r.Send(fmt.Sprintf(query, s.Name, s.Email, s.Password))
	if err != nil {
		return s.ID, err
	}

	err = resp.Body.Error()
	if err != nil {
		return s.ID, err
	}

	err = resp.Body.ConvertData(&studentRegister)
	if err != nil {
		return s.ID, err
	}

	s.ID = studentRegister.StudentRegister.ID

	return s.ID, nil
}

func (s *Student) One(ctx context.Context, r *util.PostData) (*Student, error) {
	studentDetail := struct {
		StudentDetail *Student `json:"student_detail"`
	}{
		StudentDetail: &Student{},
	}

	query := `query {
		student_detail(
			id: "%s"
		)
		{
			id,
			name,
			email
		}
	} `

	resp, err := r.Send(fmt.Sprintf(query, s.ID))
	if err != nil {
		return studentDetail.StudentDetail, err
	}

	err = resp.Body.Error()
	if err != nil {
		return studentDetail.StudentDetail, err
	}

	err = resp.Body.ConvertData(&studentDetail)
	if err != nil {
		return studentDetail.StudentDetail, err
	}
	return studentDetail.StudentDetail, nil
}

func (s *Student) All(ctx context.Context, r *util.PostData, param AllStudent) ([]*Student, error) {
	studentList := struct {
		StudentList []*Student `json:"student_list"`
	}{
		StudentList: []*Student{},
	}

	query := `query {
		student_list(
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
	} `

	resp, err := r.Send(fmt.Sprintf(query,
		param.SearchBy, param.SearchValue, param.OrderBy, param.OrderDir, param.Offset, param.Limit))
	if err != nil {
		return studentList.StudentList, err
	}

	err = resp.Body.Error()
	if err != nil {
		return studentList.StudentList, err
	}

	err = resp.Body.ConvertData(&studentList)
	if err != nil {
		return studentList.StudentList, err
	}
	return studentList.StudentList, nil
}

func (s *Student) Login(ctx context.Context, r *util.PostData) (*Student, error) {
	studentLogin := struct {
		StudentLogin *Student `json:"student_login"`
	}{
		StudentLogin: &Student{},
	}

	query := `query {
		student_login(
			email:"%s",
			password:"%s"
		)
		{
			id,
			name,
			email
		}
	}  `

	resp, err := r.Send(fmt.Sprintf(query, s.Email, s.Password))
	if err != nil {
		return studentLogin.StudentLogin, err
	}

	err = resp.Body.Error()
	if err != nil {
		return studentLogin.StudentLogin, err
	}

	err = resp.Body.ConvertData(&studentLogin)
	if err != nil {
		return studentLogin.StudentLogin, err
	}
	return studentLogin.StudentLogin, nil
}

func (s *Student) Update(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	studentUpdate := struct {
		StudentUpdate Student `json:"student_update"`
	}{
		StudentUpdate: Student{},
	}

	query := `mutation {
		student_update(
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

	resp, err := r.Send(fmt.Sprintf(query, s.ID, s.Name, s.Email, s.Password))
	if err != nil {
		return s.ID, err
	}

	err = resp.Body.Error()
	if err != nil {
		return s.ID, err
	}

	err = resp.Body.ConvertData(&studentUpdate)
	if err != nil {
		return s.ID, err
	}

	s.ID = studentUpdate.StudentUpdate.ID

	return s.ID, nil
}

func (s *Student) Delete(ctx context.Context, r *util.PostData) (uuid.UUID, error) {
	var id uuid.UUID
	return id, nil
}
