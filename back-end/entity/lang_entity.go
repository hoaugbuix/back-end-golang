package entity

import "back-end/common/entity"

type LanguagesEnity struct {
	*entity.BaseEntity
	En        string `json:"en,omitempty" db:"en,omitempty"`
	Vn        string `json:"vn,omitempty" db:"vn,omitempty"`
	Localtion string `json:"localtion,omitempty" db:"localtion,omitempty"`
}
