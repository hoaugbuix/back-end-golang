package repository

import (
	"back-end/entity"
	"context"
)

type LangRepo interface {
	GetAllLangs(context context.Context, limit string) ([]entity.LanguagesEnity, error)
}
