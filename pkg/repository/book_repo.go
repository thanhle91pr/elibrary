package repository

import (
	"context"
	"elibrary/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

const BooksCollection = "books"

type BooksRepository interface {
	CreateBooks(ctx context.Context, books model.Books) (string, error)
}

type bookRepository struct {
	collection *mongo.Collection
}

func NewBookRepository(db *mongo.Database) BooksRepository {
	return &bookRepository{collection: db.Collection(BooksCollection)}
}

func (b *bookRepository) CreateBooks(ctx context.Context, books model.Books) (string, error) {
	books.ID = primitive.NewObjectID().Hex()
	doc := bson.M{
		"_id":         books.ID,
		"id":          books.ID,
		"name":        books.Name,
		"description": books.Description,
	}
	_, err := b.collection.InsertOne(ctx, doc)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return books.ID, err
}

