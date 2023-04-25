package repository

import (
	"back-end/entity"
	"context"
)

type LogRepo interface {
	GetAllLog(context context.Context) ([]entity.LogEntity, error)
}
