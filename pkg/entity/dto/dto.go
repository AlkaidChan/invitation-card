package dto

import "github.com/alkaidchen/invitation-card/pkg/model"

type SubmitForm struct {
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Attendance int    `json:"attendance"`
}

func (s *SubmitForm) Transfer2Model() *model.Item {
	return &model.Item{
		Name:       s.Name,
		Phone:      s.Phone,
		Attendance: s.Attendance,
	}
}
