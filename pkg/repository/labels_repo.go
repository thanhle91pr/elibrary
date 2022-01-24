package repository

import (
	"context"
	"elibrary/pkg/model"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const LabelCollection = "labels"

type LabelsRepository interface {
	InsertMany(ctx context.Context, docs []interface{}) (string, error)
	Find(ctx context.Context, query interface{}, opts ...*options.FindOptions) (labels []model.Label, err error)
}

type labelsRepository struct {
	col *mongo.Collection
}

func NewLabelsRepository(db *mongo.Database) LabelsRepository {
	return &labelsRepository{
		col: db.Collection(LabelCollection),
	}
}
func (l labelsRepository) Find(ctx context.Context, query interface{}, opts ...*options.FindOptions) (labels []model.Label, err error) {
	res, err := l.col.Find(ctx, query, opts...)
	if err = res.All(ctx, &labels); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return labels, err
}

func (l *labelsRepository) InsertMany(ctx context.Context, docs []interface{}) (string, error) {
	opts := options.InsertMany().SetOrdered(false)
	res, err := l.col.InsertMany(ctx, docs, opts)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	fmt.Printf("inserted documents with IDs %v\n", res.InsertedIDs)
	return "", err
}
