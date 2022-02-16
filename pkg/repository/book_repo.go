package repository

import (
	"context"
	"elibrary/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const BooksCollection = "books"

type BooksRepository interface {
	CreateBooks(ctx context.Context, books model.Books) (string, error)
	SetLabelToBooks(ctx context.Context, labelsName string, booksName string) error
	FindOne(ctx context.Context, filter interface{}, options *options.FindOneOptions) model.Books
}

type booksRepository struct {
	collection *mongo.Collection
}

func NewBookRepository(db *mongo.Database) BooksRepository {
	return &booksRepository{collection: db.Collection(BooksCollection)}
}

func (b booksRepository) FindOne(ctx context.Context, filter interface{}, options *options.FindOneOptions) model.Books {
	var books model.Books
	err := b.collection.FindOne(ctx, filter, options).Decode(&books)
	if err != mongo.ErrNoDocuments && err != nil {
		log.Fatal(err)
	}
	return books
}
func (b booksRepository) SetLabelToBooks(ctx context.Context, labelsName string, booksName string) error {
	filter := bson.M{"name": booksName}
	update := bson.M{"$set": bson.M{"labels": labelsName}}
	err := b.collection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetUpsert(true)).Err()
	if err != mongo.ErrNoDocuments && err != nil {
		return err
	}
	return nil
}

func (b *booksRepository) CreateBooks(ctx context.Context, books model.Books) (string, error) {
	doc := bson.M{
		"name":        books.Name,
		"description": books.Description,
	}
	_, err := b.collection.InsertOne(ctx, doc)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return "", err
}
