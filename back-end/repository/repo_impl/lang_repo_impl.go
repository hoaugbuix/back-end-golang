package repo_impl

import (
	"back-end/db"
	"back-end/entity"
	"back-end/repository"
	"context"
	"log"
)

type LangRepoImpl struct {
	sql *db.Sql
}

func NewLangRepo(sql *db.Sql) repository.LangRepo {
	return &LangRepoImpl{
		sql: sql,
	}
}

func (l *LangRepoImpl) GetAllLangs(context context.Context, limit string) ([]entity.LanguagesEnity, error) {
	langs := []entity.LanguagesEnity{}
	err := l.sql.Db.SelectContext(context, &langs, "SELECT * FROM languages ORDER BY created_at ASC limit $1", limit)
	if err != nil {
		log.Fatal(err.Error())
		return langs, err
	}

	return langs, nil
}
