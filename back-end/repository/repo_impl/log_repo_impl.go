package repo_impl

import (
	"back-end/db"
	"log"

	"back-end/entity"
	"back-end/repository"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type LogRepoImpl struct {
	mongo *db.MongoDB
}

func NewLogRepo(mongo *db.MongoDB) repository.LogRepo {
	return &LogRepoImpl{
		mongo: mongo,
	}
}

func (l *LogRepoImpl) GetAllLog(context context.Context) ([]entity.LogEntity, error) {
	logColection := l.mongo.Db.Database("ShopClone").Collection("logs")

	var logs []entity.LogEntity

	results, err := logColection.Find(context, bson.M{})
	if err != nil {
		log.Fatal(err.Error())
		return logs, err
	}
	//reading from the db in an optimal way
	defer results.Close(context)
	for results.Next(context) {
		singleLog := entity.LogEntity{}
		if err = results.Decode(&singleLog); err != nil {
			log.Fatal(err.Error())
			return logs, err
		}
		logs = append(logs, singleLog)
	}

	return logs, nil
}
