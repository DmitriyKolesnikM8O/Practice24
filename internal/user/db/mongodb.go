package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/user"
	"github.com/DmitriyKolesnikM8O/Practice24/pkg/logging"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Реализация mongodb
type db struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

// func (d db) Create(ctx context.Context, product user.User) (string, error) {
func (d db) Create(ctx context.Context, product user.Product) (string, error) {
	d.logger.Debugf("create product")
	result, err := d.collection.InsertOne(ctx, product)
	if err != nil {
		return "", fmt.Errorf("error inserting product: %w", err)
	}

	d.logger.Debugf("convert InsertedID to ObjectID")
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	d.logger.Trace(product)
	return "", fmt.Errorf("failed to convert obj ID to HEX: %w", err)
}

func (d db) FindOne(ctx context.Context, id string) (p user.Product, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return p, fmt.Errorf("failed to convert Hex to ObjectID: %w", err)
	}
	//запрос
	filter := bson.M{"_id": oid}
	result := d.collection.FindOne(ctx, filter)

	//если ошибка или нет в базе
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			//TODO Not found
			return p, fmt.Errorf("not found")
		}
		return p, fmt.Errorf("failed to find product by ID: %w", result.Err())
	}

	if err := result.Decode(&p); err != nil {
		return p, fmt.Errorf("failed to decode product by ID: %w", result.Err())
	}
	return p, nil
}

func (d db) Update(ctx context.Context, product user.Product) error {
	objectID, err := primitive.ObjectIDFromHex(product.ID)
	if err != nil {
		return fmt.Errorf("failed to convert Hex to ObjectID: %w", err)
	}

	filter := bson.M{"_id": objectID}

	productBytes, err := bson.Marshal(product)
	if err != nil {
		return fmt.Errorf("failed to marshal product: %w", err)
	}
	//{"id": ""}
	var updateProductObj bson.M
	err = bson.Unmarshal(productBytes, &updateProductObj)
	if err != nil {
		return fmt.Errorf("failed to unmarshal update: %w", err)
	}
	//удаляем id, чтобы не обновилось в базе
	delete(updateProductObj, "_id")

	update := bson.M{
		"$set": updateProductObj,
	}

	result, err := d.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update product: %w", err)
	}
	//если 0, то ничего не нашлось
	if result.MatchedCount == 0 {
		//TODO Error not found (404)
		return fmt.Errorf("failed to find updated product")
	}
	d.logger.Tracef("Matcheed: %d products and Modified %d products", result.MatchedCount, result.ModifiedCount)
	return nil
}

func (d db) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("failed to convert Hex to ObjectID: %w", err)
	}

	filter := bson.M{"_id": objectID}

	result, err := d.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}
	if result.DeletedCount == 0 {
		//TODO ErrEntityNotFound
		return fmt.Errorf("nof found")
	}
	d.logger.Tracef("Delete documents: %d", result.DeletedCount)
	return nil
}

func NewStorage(database *mongo.Database, collection string, logger *logging.Logger) user.Storage {
	return &db{
		collection: database.Collection(collection),
		logger:     logger,
	}
}
