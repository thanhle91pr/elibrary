package repository

import (
	"context"
	"elibrary/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const SongsCollection = "songs"

type SongsRepository interface {
	CreateSongs(ctx context.Context, songs model.Songs) (string, error)
	SetLabelToSongs(ctx context.Context, labelsName string, songsName string) error
}

type songsRepository struct {
	db  *mongo.Database
	col *mongo.Collection
}

func NewSongsRepository(db *mongo.Database) SongsRepository {
	return &songsRepository{
		col: db.Collection(SongsCollection),
	}
}

func (s songsRepository) SetLabelToSongs(ctx context.Context, labelsName string, songsName string) error {
	filter := bson.M{"name": songsName}
	update := bson.M{"$set": bson.M{"labels": labelsName}}
	err := s.col.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetUpsert(true)).Err()
	if err != mongo.ErrNoDocuments && err != nil {
		return err
	}
	return nil
}

func (s *songsRepository) CreateSongs(ctx context.Context, songs model.Songs) (string, error) {
	doc := bson.M{
		"name":        songs.Name,
		"description": songs.Description,
	}
	_, err := s.col.InsertOne(ctx, doc)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return "", err
}
