package adapters

import (
	"context"
	"fmt"
	"time"

	"github.com/pipusana/goapi/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoAdapter interface {
	InsertOne(doc any) (string, error)
	FindAll() ([]entities.Nisit, error)
	FindOne(nisit_id string) (entities.Nisit, error)
	UpdateOne(nisit_id string, nisit_update entities.NisitUpdate) error
	DeleteOne(nisit_id string) error
	CloseMongoAdapter()
}

type mongoAdapter struct {
	connection *mongo.Client
	collection *mongo.Collection
}

func NewMongoAdapter(mongo_uri, database, collection string) MongoAdapter {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(mongo_uri)
	client, err := mongo.Connect(ctx, clientOpts)

	if err != nil {
		panic(fmt.Errorf("fatal error mongo connection: %s", err.Error()))
	}

	return &mongoAdapter{
		client,
		client.Database(database).Collection(collection),
	}
}

func (m *mongoAdapter) InsertOne(doc interface{}) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := m.collection.InsertOne(ctx, doc)
	if err != nil {
		return "0", err
	}
	return fmt.Sprintf("%v", result.InsertedID), err
}

func (m *mongoAdapter) FindOne(nisit_id string) (entities.Nisit, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var nisit entities.Nisit

	err := m.collection.FindOne(ctx, bson.D{{"id", nisit_id}}).Decode(&nisit)
	if err != nil {
		return nisit, err
	}

	return nisit, nil
}

func (m *mongoAdapter) UpdateOne(nisit_id string, nisit_update entities.NisitUpdate) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{"id", nisit_id}}
	update := bson.D{primitive.E{Key: "$set", Value: nisit_update}}
	_, err := m.collection.UpdateOne(
		ctx,
		filter,
		update,
	)
	if err != nil {
		return err
	}

	return nil
}

func (m *mongoAdapter) DeleteOne(nisit_id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := m.collection.DeleteOne(ctx, bson.D{{"id", nisit_id}})
	if err != nil {
		return err
	}
	return nil
}

func (m *mongoAdapter) FindAll() ([]entities.Nisit, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var nisits []entities.Nisit
	result, err := m.collection.Find(ctx, bson.M{})
	result.All(ctx, &nisits)
	if err != nil {
		return nil, err
	}
	return nisits, nil
}

func (m *mongoAdapter) CloseMongoAdapter() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	m.connection.Disconnect(ctx)
}
