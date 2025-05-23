package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoService struct {
	Collection *mongo.Collection
}

func NewMongoService(coll *mongo.Collection) *MongoService {
	return &MongoService{
		Collection: coll,
	}
}

func (m *MongoService) withTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

func (m *MongoService) FindOne(filter interface{}, projection interface{}, result interface{}) error {
	ctx, cancel := m.withTimeout()
	defer cancel()

	opts := options.FindOne()
	if projection != nil {
		opts.Projection = projection
	}

	res := m.Collection.FindOne(ctx, filter, opts)
	return res.Decode(result)
}

func (m *MongoService) FindMany(filter interface{}, projection interface{}, sort interface{}, results interface{}) error {
	ctx, cancel := m.withTimeout()
	defer cancel()

	opts := options.Find()
	if projection != nil {
		opts.Projection = projection
	}
	if sort != nil {
		opts.Sort = sort
	}

	cursor, err := m.Collection.Find(ctx, filter, opts)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	return cursor.All(ctx, results)
}

func (m *MongoService) InsertOne(document interface{}) (*mongo.InsertOneResult, error) {
	ctx, cancel := m.withTimeout()
	defer cancel()

	return m.Collection.InsertOne(ctx, document)
}

func (m *MongoService) InsertMany(documents []interface{}) (*mongo.InsertManyResult, error) {
	ctx, cancel := m.withTimeout()
	defer cancel()

	return m.Collection.InsertMany(ctx, documents)
}

func (m *MongoService) UpdateOne(filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	ctx, cancel := m.withTimeout()
	defer cancel()

	return m.Collection.UpdateOne(ctx, filter, update)
}

func (m *MongoService) UpdateMany(filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	ctx, cancel := m.withTimeout()
	defer cancel()

	return m.Collection.UpdateMany(ctx, filter, update)
}

func (m *MongoService) UpsertOne(filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	ctx, cancel := m.withTimeout()
	defer cancel()

	opts := options.Update().SetUpsert(true)
	return m.Collection.UpdateOne(ctx, filter, update, opts)
}

func (m *MongoService) DeleteOne(filter interface{}) (*mongo.DeleteResult, error) {
	ctx, cancel := m.withTimeout()
	defer cancel()

	return m.Collection.DeleteOne(ctx, filter)
}

func (m *MongoService) CountAll(filter interface{}) (int64, error) {
	ctx, cancel := m.withTimeout()
	defer cancel()

	return m.Collection.CountDocuments(ctx, filter)
}

func (m *MongoService) Aggregate(pipeline interface{}, results interface{}) error {
	ctx, cancel := m.withTimeout()
	defer cancel()

	cursor, err := m.Collection.Aggregate(ctx, pipeline)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	return cursor.All(ctx, results)
}
