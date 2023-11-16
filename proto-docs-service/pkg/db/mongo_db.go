package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"proto-docs-service/grpc/docs"
)

const collectionName = "grpcReports"
const db = "loadtest"

type ReportMongoRepository struct {
	Client mongo.Client
	ReportClientRepository
}

func NewMongoRepository(address string) (*ReportMongoRepository, error) {
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

	return &ReportMongoRepository{Client: *client}, nil
}

func (rm *ReportMongoRepository) Find(limit int, offset int) ([]docs.Document, error) {
	collection := rm.Client.Database(db).Collection(collectionName)
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

	var docsList []docs.Document
	for cursor.Next(context.Background()) {
		var doc docs.Document
		if err := cursor.Decode(&doc); err != nil {
			return nil, err
		}
		docsList = append(docsList, doc)
	}

	return docsList, nil
}

func (rm *ReportMongoRepository) Save(doc *docs.Document) error {
	collection := rm.Client.Database(db).Collection(collectionName)
	insertResult, err := collection.InsertOne(context.Background(), doc)
	if err != nil {
		return err
	}

	fmt.Println("Inserted. ID:", insertResult.InsertedID)
	return nil
}
