package db

import (
	"context"
	"fmt"
	"time"

	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Db       *mongo.Client
	Host     string
	Port     int
	Username string
	Password string
	DbName   string
	Url      string
}

func (m *MongoDB) ConnectMG() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dataSource := fmt.Sprintf(m.Url)
	if m.Url == "" {
		dataSource = fmt.Sprintf("mongodb://%s:%s@%s:%d", m.Username, m.Password, m.Host, m.Port)
	}

	if m.Url == "" && (m.Username == "" || m.Password == "") {
		dataSource = fmt.Sprintf("mongodb://%s:%d", m.Host, m.Port)
	}

	m.Db, _ = mongo.NewClient(options.Client().ApplyURI(dataSource))

	if err := m.Db.Connect(ctx); err != nil {
		log.Error(err.Error())
		return
	}
	//ping the database
	if err := m.Db.Ping(ctx, nil); err != nil {
		log.Error(err.Error())
		return
	}
	fmt.Println("Connected to MongoDB")
}

func (m *MongoDB) CloseMG() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := m.Db.Disconnect(ctx); err != nil {
		log.Error(err.Error())
		return
	}
}
