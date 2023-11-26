package db

import (
	"context"
	"flat-docs-service/pkg/model"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const db = "loadtest"

type ReportMongoRepository struct {
	ReportClientRepository

	Client         mongo.Client
	CollectionName string
}

func NewMongoRepository(address, collectionName string) (*ReportMongoRepository, error) {
	clientOptions := options.Client().ApplyURI(address)
	clientOptions.SetAuth(options.Credential{ // not secure, I know, sorry :D
		Username: "root",
		Password: "root",
	})
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return &ReportMongoRepository{Client: *client, CollectionName: collectionName}, nil
}

func (rm *ReportMongoRepository) Find(limit int, offset int) ([]*model.Document, error) {
	collection := rm.Client.Database(db).Collection(rm.CollectionName)
	findOptions := options.Find().SetLimit(int64(limit)).SetSkip(int64(offset))

	cursor, err := collection.Find(context.Background(), bson.D{}, findOptions)
	if err != nil {
		log.Println("Error while processing cursor:", err)
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Println("Error closing cursor:", err)
		}
	}(cursor, context.Background())

	var docsList []*model.Document
	for cursor.Next(context.Background()) {
		var doc model.Document
		if err := cursor.Decode(&doc); err != nil {
			return nil, err
		}
		docsList = append(docsList, &doc)
	}

	return docsList, nil
}

func (rm *ReportMongoRepository) Save(doc model.Document) error {
	collection := rm.Client.Database(db).Collection(rm.CollectionName)
	insertResult, err := collection.InsertOne(context.Background(), doc)
	if err != nil {
		return err
	}

	fmt.Println("Inserted. ID:", insertResult.InsertedID)
	return nil
}
