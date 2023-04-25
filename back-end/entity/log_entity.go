package entity

import "back-end/common/entity"

type LogEntity struct {
	*entity.BaseMongEntity `bson:",inline"`
	// Name string `json:"name,omitempty" validate:"required" bson:"name,omitempty"`
}
