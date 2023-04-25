package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseMongEntity struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	ActiveFlag  int                `json:"activeFlag,omitempty" bson:"active_flag,omitempty"`
	CreatedDate time.Time          `json:"createdDate" bson:"created_date,omitempty"`
	UpdatedDate time.Time          `json:"updatedDate" bson:"updated_date,omitempty"`
}
