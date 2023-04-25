package entity

import "time"

type BaseEntity struct {
	Id          string    `json:"id,omitempty" db:"id,omitempty"`
	ActiveFlag  int       `json:"activeFlag,omitempty" db:"active_flag,omitempty"`
	CreatedDate time.Time `json:"createdDate" db:"created_date,omitempty"`
	UpdatedDate time.Time `json:"updatedDate" db:"updated_date,omitempty"`
}
